package server

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/hashicorp/raft"
	"google.golang.org/grpc"

	raftboltdb "github.com/hashicorp/raft-boltdb"
	context "golang.org/x/net/context"

	"github.com/namreg/godown/internal/api"
	"github.com/namreg/godown/internal/clock"
	"github.com/namreg/godown/internal/command"
	"github.com/namreg/godown/internal/storage"
)

const (
	defaultListenAddr = "127.0.0.1:4000"
	defaultRaftAddr   = "127.0.0.1:4001"
	defaultDir        = "../../data"
	defaultGCInterval = 500 * time.Millisecond
)

const (
	raftApplyTimeout     = 500 * time.Millisecond
	raftLogCacheSize     = 512
	raftMaxPoolSize      = 3
	raftTransportTimeout = 10 * time.Second
	raftSnapshotsRetain  = 3
	raftDBFile           = "raft.db"
)

const leaderIPMetaKey = "leader_ip"

//commandParser parses command with args from the given string.
type commandParser interface {
	Parse(str string) (cmd command.Command, args []string, err error)
}

//metaStore is used for storing meta information.
type metaStore interface {
	//AllMeta returns all stored metadata.
	AllMeta() (map[storage.MetaKey]storage.MetaValue, error)
	//PutMeta puts a new value at the given key.
	PutMeta(storage.MetaKey, storage.MetaValue) error
	//GetMeta get a value at the given key.
	GetMeta(storage.MetaKey) (storage.MetaValue, error)
	//RestoreMeta replaces current metadata with the given one.
	RestoreMeta(map[storage.MetaKey]storage.MetaValue) error
}

//dataStore is used to access stored data.
//go:generate minimock -i github.com/namreg/godown/internal/server.dataStore -o ./
type dataStore interface {
	//AllWithTTL returns all stored values.
	All() (map[storage.Key]*storage.Value, error)
	//AllWithTTL returns all values that have TTL.
	AllWithTTL() (map[storage.Key]*storage.Value, error)
	//Del deletes the given key.
	Del(storage.Key) error
	//Restore restores current data with the given one.
	Restore(map[storage.Key]*storage.Value) error
}

//go:generate minimock -i github.com/namreg/godown/internal/server.serverClock -o ./
type serverClock interface {
	Now() time.Time
}

//Options is a server options.
type Options struct {
	ID         string
	ListenAddr string
	RaftAddr   string
	Dir        string
	Logger     *log.Logger
	Clock      serverClock
	GCInterval time.Duration
}

//DefaultOptions returns default options.
func DefaultOptions() Options {
	return Options{
		ListenAddr: defaultListenAddr,
		RaftAddr:   defaultRaftAddr,
		Dir:        defaultDir,
		Logger:     log.New(os.Stdout, "", log.LstdFlags),
		Clock:      clock.New(),
		GCInterval: defaultGCInterval,
	}
}

//Server represents a server that handles user requests and executes commands.
type Server struct {
	meta   metaStore
	data   dataStore
	parser commandParser

	opts Options

	gc            *gc
	srv           *grpc.Server
	leader        *grpc.ClientConn
	raft          *raft.Raft //consensus protocol
	raftTransport raft.Transport
}

//New creates a new server.
func New(meta metaStore, data dataStore, parser commandParser, opts Options) *Server {
	srv := &Server{
		data:   data,
		meta:   meta,
		parser: parser,
		opts:   opts,
	}

	return srv
}

//BootstrapCluster creates a new cluster and runs the server.
func (s *Server) BootstrapCluster() error {
	if err := s.setupRaft(); err != nil {
		return fmt.Errorf("could not to setup raft node: %v", err)
	}
	configuration := raft.Configuration{
		Servers: []raft.Server{
			{
				ID:      raft.ServerID(s.opts.ID),
				Address: s.raftTransport.LocalAddr(),
			},
		},
	}
	s.raft.BootstrapCluster(configuration)

	return s.start(s.opts.ListenAddr)
}

//JoinCluster joins to an existing cluster and runs the server.
func (s *Server) JoinCluster(joinAddr string) error {
	if err := s.setupRaft(); err != nil {
		return fmt.Errorf("could not to setup raft node: %v", err)
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- s.start(s.opts.ListenAddr)
	}()

	conn, err := grpc.Dial(joinAddr, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("could not dial %s: %v", joinAddr, err)
	}

	client := api.NewGodownClient(conn)
	req := &api.AddToClusterRequest{
		Id:   s.opts.ID,
		Addr: s.opts.RaftAddr,
	}
	if _, err = client.AddToCluster(context.Background(), req); err != nil {
		return fmt.Errorf("could not add a new node to the cluster: %v", err)
	}
	conn.Close()

	return <-errCh
}

func (s *Server) setupRaft() error {
	if s.opts.ID == "" {
		return errors.New("empty server ID")
	}
	absDir, err := filepath.Abs(filepath.Join(s.opts.Dir, fmt.Sprintf("node%s", s.opts.ID)))
	if err != nil {
		return err
	}
	s.opts.Dir = absDir

	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(s.opts.ID)
	config.Logger = s.opts.Logger

	addr, err := net.ResolveTCPAddr("tcp", s.opts.RaftAddr)
	if err != nil {
		return fmt.Errorf("could not resolve tcp address: %v", err)
	}

	s.raftTransport, err = raft.NewTCPTransport(s.opts.RaftAddr, addr, raftMaxPoolSize, raftTransportTimeout, os.Stderr)
	if err != nil {
		return fmt.Errorf("could create create TCP transport: %v", err)
	}

	snapshots, err := raft.NewFileSnapshotStore(s.opts.Dir, raftSnapshotsRetain, os.Stderr)
	if err != nil {
		return fmt.Errorf("could not create file snapshot store: %v", err)
	}

	store, err := raftboltdb.NewBoltStore(filepath.Join(s.opts.Dir, raftDBFile))
	if err != nil {
		return fmt.Errorf("could not create bolt store: %v", err)
	}

	logStore, err := raft.NewLogCache(raftLogCacheSize, store)
	if err != nil {
		return fmt.Errorf("could not create log cache: %v", err)
	}

	s.raft, err = raft.NewRaft(config, newFsm(s), logStore, store, snapshots, s.raftTransport)
	if err != nil {
		return fmt.Errorf("could not create raft node: %v", err)
	}

	go s.whenLeaderChanged(s.updateLeaderIP, s.controlGC, s.controlLeaderConn)

	return nil
}

//Start starts a server.
func (s *Server) start(hostPort string) error {
	l, err := net.Listen("tcp", hostPort)
	if err != nil {
		return fmt.Errorf("could not listen on %s: %v", hostPort, err)
	}

	s.srv = grpc.NewServer()
	api.RegisterGodownServer(s.srv, s)

	return s.srv.Serve(l)
}

//Stop stops a grpc server.
func (s *Server) Stop() error {
	s.srv.Stop()
	if s.leader != nil {
		return s.leader.Close()
	}
	return nil
}

//whenLeaderChanged executes given functions when a leader in the cluster changed.
func (s *Server) whenLeaderChanged(funcs ...func(isLeader bool) error) {
	for isLeader := range s.raft.LeaderCh() {
		wg := new(sync.WaitGroup)
		wg.Add(len(funcs))
		for _, f := range funcs {
			go func(f func(isLeader bool) error) {
				defer wg.Done()
				if err := f(isLeader); err != nil {
					s.opts.Logger.Printf("[WARN] server: error while executing function when leader changed: %v", err)
				}
			}(f)
		}
		wg.Wait()
	}
}

func (s *Server) updateLeaderIP(isLeader bool) error {
	if !isLeader {
		return nil
	}
	cmd, err := newApplyMetadataFSMCommand(leaderIPMetaKey, s.opts.ListenAddr)
	if err != nil {
		return fmt.Errorf("could not create set meta fsm command: %v", err)
	}

	b, err := proto.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("could not marshal fsm command: %v", err)
	}

	future := s.raft.Apply(b, raftApplyTimeout)

	if err = future.Error(); err != nil {
		return fmt.Errorf("could not apply set meta value request: %v", err)
	}

	if err, ok := future.Response().(error); ok {
		return fmt.Errorf("could not apply set meta value request: %v", err)
	}
	return nil
}

func (s *Server) controlGC(isLeader bool) error {
	if isLeader {
		if s.gc == nil {
			s.gc = newGc(s, s.data, s.opts.Logger, s.opts.Clock, s.opts.GCInterval)
		}
		s.gc.start()
	} else if s.gc != nil {
		s.gc.stop()
	}
	return nil
}

func (s *Server) controlLeaderConn(isLeader bool) error {
	var err error
	if isLeader {
		if s.leader != nil {
			err = s.leader.Close()
			s.leader = nil
		}
		return err
	}
	s.leader, err = s.leaderConn()
	if err != nil {
		return fmt.Errorf("could not connect to leader: %v", err)
	}
	return nil

}

func (s *Server) leaderConn() (*grpc.ClientConn, error) {
	if s.leader != nil {
		return s.leader, nil
	}
	leaderIP, err := s.meta.GetMeta(storage.MetaKey(leaderIPMetaKey))
	if err != nil {
		return nil, fmt.Errorf("could not get leader ip from meta store: %v", err)
	}
	conn, err := grpc.Dial(string(leaderIP), grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("could not dial %s: %v", leaderIP, err)
	}
	return conn, err
}

//ExecuteCommand executes a command that placed into the request.
func (s *Server) ExecuteCommand(ctx context.Context, req *api.ExecuteCommandRequest) (*api.ExecuteCommandResponse, error) {
	cmd, args, err := s.parser.Parse(req.Command)
	if err != nil {
		if err == command.ErrCommandNotFound {
			return &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  fmt.Sprintf("command %q not found", req.Command),
			}, nil
		}
		return nil, fmt.Errorf("could not parse command: %v", err)
	}

	if s.isCommandModifiesState(cmd) {
		if !s.isLeader() {
			conn, err := s.leaderConn()
			if err != nil {
				return nil, err
			}
			return api.NewGodownClient(conn).ExecuteCommand(ctx, req)
		}

		fsmCmd, err := newExecuteFSMCommand(req.Command)
		if err != nil {
			return nil, fmt.Errorf("could not create execute fsm command: %v", err)
		}

		b, err := proto.Marshal(fsmCmd)
		if err != nil {
			return nil, fmt.Errorf("could not marshal fsm command: %v", err)
		}

		future := s.raft.Apply(b, raftApplyTimeout)
		if err := future.Error(); err != nil {
			return nil, fmt.Errorf("could not apply raft log entry: %v", err)
		}

		if err, ok := future.Response().(error); ok {
			return nil, fmt.Errorf("could not apply raft log entry: %v", err)
		}

		applyResp := &api.ExecuteCommandResponse{}

		err = proto.Unmarshal(future.Response().([]byte), applyResp)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal applied response: %v", err)
		}

		return applyResp, nil
	}

	res := cmd.Execute(args...)
	if err != nil {
		return nil, fmt.Errorf("could not execute command: %v", err)
	}
	return s.createResponse(res)
}

//AddToCluster add a new node to the existing cluster.
func (s *Server) AddToCluster(ctx context.Context, req *api.AddToClusterRequest) (*api.AddToClusterResponse, error) {
	if !s.isLeader() {
		conn, err := s.leaderConn()
		if err != nil {
			return nil, err
		}
		return api.NewGodownClient(conn).AddToCluster(ctx, req)
	}
	configFuture := s.raft.GetConfiguration()
	if err := configFuture.Error(); err != nil {
		return nil, fmt.Errorf("failed to get raft configuration: %v", err)
	}

	for _, srv := range configFuture.Configuration().Servers {
		if srv.ID == raft.ServerID(req.Id) || srv.Address == raft.ServerAddress(req.Addr) {
			if srv.Address == raft.ServerAddress(req.Addr) && srv.ID == raft.ServerID(req.Id) {
				return &api.AddToClusterResponse{}, nil
			}

			future := s.raft.RemoveServer(srv.ID, 0, 0)
			if err := future.Error(); err != nil {
				return nil, fmt.Errorf("error removing existing node %s at %s: %s", req.Id, req.Addr, err)
			}
		}
	}

	f := s.raft.AddVoter(raft.ServerID(req.Id), raft.ServerAddress(req.Addr), 0, 0)
	if f.Error() != nil {
		return nil, f.Error()
	}
	return &api.AddToClusterResponse{}, nil
}

func (s *Server) createResponse(res command.Reply) (*api.ExecuteCommandResponse, error) {
	apiRes := new(api.ExecuteCommandResponse)

	switch t := res.(type) {
	case command.NilReply:
		apiRes.Reply = api.NilCommandReply
	case command.OkReply:
		apiRes.Reply = api.OkCommandReply
	case command.RawStringReply:
		apiRes.Reply = api.RawStringCommandReply
		apiRes.Item = t.Value
	case command.StringReply:
		apiRes.Reply = api.StringCommandReply
		apiRes.Item = t.Value
	case command.IntReply:
		apiRes.Reply = api.IntCommandReply
		apiRes.Item = strconv.FormatInt(t.Value, 10)
	case command.SliceReply:
		apiRes.Reply = api.SliceCommandReply
		apiRes.Items = t.Value
	case command.ErrReply:
		apiRes.Reply = api.ErrCommandReply
		apiRes.Item = t.Value.Error()
	default:
		return nil, fmt.Errorf("unsupported type %T", res)
	}

	return apiRes, nil
}

func (s *Server) isCommandModifiesState(cmd command.Command) bool {
	switch cmd.(type) {
	case *command.Set, *command.Del, *command.Expire, *command.Hset,
		*command.Lpop, *command.Lpush, *command.Lrem, *command.SetBit,
		*command.Rpop, *command.Rpush:
		return true
	}
	return false
}

func (s *Server) handleSetMetaValueRequest(req *api.UpdateMetadataRequest) (*api.UpdateMetadataResponse, error) {
	key := storage.MetaKey(req.Key)
	value := storage.MetaValue(req.Value)
	if err := s.meta.PutMeta(key, value); err != nil {
		return nil, fmt.Errorf("could not save meta %q: %v", key, err)
	}
	return &api.UpdateMetadataResponse{}, nil
}

func (s *Server) handleExecuteCommandRequest(req *api.ExecuteCommandRequest) (*api.ExecuteCommandResponse, error) {
	cmd, args, err := s.parser.Parse(req.Command)
	if err != nil {
		return nil, err
	}
	res := cmd.Execute(args...)
	resp, err := s.createResponse(res)
	if err != nil {
		return nil, fmt.Errorf("could not create response: %v", err)
	}
	return resp, nil
}

func (s *Server) isLeader() bool {
	return s.raft.State() == raft.Leader
}
