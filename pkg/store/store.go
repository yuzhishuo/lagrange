package store

import (
	"fmt"
	"unsafe"

	"github.com/matrixorigin/talent-challenge/matrixbase/distributed/pkg/cfg"
	"go.etcd.io/etcd/raft/v3/raftpb"
)

// Store the store interface
type Store interface {
	// Set set key-value to store
	Set(key []byte, value []byte) error
	// Get returns the value from store
	Get(key []byte) ([]byte, error)
	// Delete remove the key from store
	Delete(key []byte) error
}

// NewStore create the raft store
func NewStore(cfg cfg.StoreCfg) (Store, error) {

	if cfg.Memory {
		return newMemoryStore()
	}

	proposeC := make(chan string)
	defer close(proposeC)
	confChangeC := make(chan raftpb.ConfChange)
	defer close(confChangeC)
	var kvs Store
	fmt.Println(cfg.DataPath)
	getSnapshot := func() ([]byte, error) { return (*persistentStore)(unsafe.Pointer(&kvs)).getSnapshot() }
	commitC, errorC, snapshotterReady := newRaftNode(1, []string{"http://127.0.0.1:12379"}, false, getSnapshot, proposeC, confChangeC)

	var err error
	if kvs, err = newPersistentStore(cfg, <-snapshotterReady, proposeC, commitC, errorC); err != nil {
		return nil, err
	}

	return kvs, nil
}
