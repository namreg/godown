package server

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"

	"github.com/centrifugal/centrifugo/libcentrifugo/raw"
	"github.com/gogo/protobuf/proto"

	"github.com/hashicorp/raft"
	"github.com/namreg/godown-v2/internal/api"
	"github.com/namreg/godown-v2/internal/storage"
)

//fsm stands for a finite state machine.
type fsm struct {
	srv *Server
}

func newFsm(s *Server) *fsm {
	return &fsm{srv: s}
}

//Apply applies raft log entry.
func (f *fsm) Apply(entry *raft.Log) interface{} {
	fsmCommand := &api.FSMCommand{}
	err := proto.Unmarshal(entry.Data, fsmCommand)
	if err != nil {
		return fmt.Errorf("could not unmarshal log entry: %v", err)
	}

	var resp proto.Message

	switch {
	case fsmCommand.Type == api.FSMApplyMetadata:
		req := &api.UpdateMetadataRequest{}
		if err := proto.Unmarshal([]byte(fsmCommand.Command), req); err != nil {
			return fmt.Errorf("could not unmarshal set meta value request: %v", err)
		}
		resp, err = f.srv.handleSetMetaValueRequest(req)
		if err != nil {
			return err
		}
	case fsmCommand.Type == api.FSMApplyCommand:
		req := &api.ExecuteCommandRequest{}
		if err := proto.Unmarshal([]byte(fsmCommand.Command), req); err != nil {
			return fmt.Errorf("could not unmarshal execute command request: %v", err)
		}
		resp, err = f.srv.handleExecuteCommandRequest(req)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unrecognized fsm command: %v", fsmCommand.Type)
	}

	b, err := proto.Marshal(resp)
	if err != nil {
		return fmt.Errorf("could not marshal response: %v", err)
	}
	return b
}

//Snapshot implements raft.FSM interface.
func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	s := &fsmSnapshot{}
	meta, err := f.srv.meta.AllMeta()
	if err != nil {
		return nil, fmt.Errorf("could not get all metadata: %v", err)
	}
	mb, err := json.Marshal(meta)
	if err != nil {
		return nil, fmt.Errorf("could not marshal metadata: %v", err)
	}
	s.meta = mb

	data, err := f.srv.data.All()
	if err != nil {
		return nil, fmt.Errorf("could not get all data: %v", err)
	}
	db, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("could not marshal data: %v", err)
	}
	s.data = db
	return s, nil
}

//Restore implements raft.FSM interface.
func (f *fsm) Restore(rc io.ReadCloser) error {
	defer rc.Close()

	var l length

	if err := binary.Read(rc, binary.LittleEndian, &l); err != nil {
		return fmt.Errorf("could not read meta length: %v", err)
	}

	mb := make([]byte, l)

	if n, err := io.ReadFull(rc, mb); int64(n) != int64(l) || err != nil {
		return fmt.Errorf("could not read meta")
	}

	meta := make(map[storage.MetaKey]storage.MetaValue)

	if err := json.Unmarshal(mb, &meta); err != nil {
		return fmt.Errorf("could not unmarshal meta: %v", err)
	}

	if err := f.srv.meta.RestoreMeta(meta); err != nil {
		return fmt.Errorf("could restore meta: %v", err)
	}

	if err := binary.Read(rc, binary.LittleEndian, &l); err != nil {
		return fmt.Errorf("could not read data length: %v", err)
	}

	db := make([]byte, l)

	if n, err := io.ReadFull(rc, db); int64(n) != int64(l) || err != nil {
		return fmt.Errorf("could not read data")
	}

	data := make(map[storage.Key]*storage.Value)
	if err := json.Unmarshal(db, &data); err != nil {
		return fmt.Errorf("could unmarshal data: %v", err)
	}

	if err := f.srv.data.Restore(data); err != nil {
		return fmt.Errorf("could not restore data: %v", err)
	}

	return nil
}

func newSetMetaFSMCommand(key, value string) (*api.FSMCommand, error) {
	return newFSMCommand(api.FSMApplyMetadata, &api.UpdateMetadataRequest{Key: key, Value: value})
}

func newExecuteFSMCommand(command string) (*api.FSMCommand, error) {
	return newFSMCommand(api.FSMApplyCommand, &api.ExecuteCommandRequest{Command: command})
}

func newFSMCommand(t api.FSMCommandType, req proto.Message) (*api.FSMCommand, error) {
	cmd := &api.FSMCommand{Type: t}
	b, err := proto.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("could not marhal proto request: %v", err)
	}
	cmd.Command = raw.Raw(b)
	return cmd, nil
}
