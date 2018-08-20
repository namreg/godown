package server

import (
	"fmt"
	"io"

	"github.com/centrifugal/centrifugo/libcentrifugo/raw"
	"github.com/gogo/protobuf/proto"

	"github.com/hashicorp/raft"
	"github.com/namreg/godown-v2/internal/api"
)

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

func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}

func (f *fsm) Restore(io.ReadCloser) error {
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
