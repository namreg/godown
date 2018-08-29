package server

import (
	"encoding/binary"
	"fmt"

	"github.com/hashicorp/raft"
)

type fsmSnapshot struct {
	meta []byte
	data []byte
}

type length int64

func (fs *fsmSnapshot) Persist(sink raft.SnapshotSink) (err error) {
	defer func() {
		if err == nil {
			err = sink.Close()
		} else {
			err = sink.Cancel()
		}
	}()
	if err := binary.Write(sink, binary.LittleEndian, length(len(fs.meta))); err != nil {
		return fmt.Errorf("could not write meta length: %v", err)
	}
	if _, err := sink.Write(fs.meta); err != nil {
		return fmt.Errorf("could not write meta: %v", err)
	}
	if err := binary.Write(sink, binary.LittleEndian, length(len(fs.data))); err != nil {
		return fmt.Errorf("could not write data length: %v", err)
	}
	if _, err := sink.Write(fs.data); err != nil {
		return fmt.Errorf("could not write data: %v", err)
	}
	return nil
}

//Release is invoked when we are finished with the snapshot.
func (fs *fsmSnapshot) Release() {
	//noop
}
