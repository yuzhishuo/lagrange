package store

import (
	"fmt"
	"strings"

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

	getSnapshot() ([]byte, error)
}

// NewStore create the raft store
func NewStore(cfg cfg.StoreCfg) (Store, error) {

	if cfg.Memory {
		return newMemoryStore()
	}

	proposeC := make(chan string)
	// defer close(proposeC)
	confChangeC := make(chan raftpb.ConfChange)
	// defer close(confChangeC)
	var kvs Store
	fmt.Println(cfg.DataPath)
	getSnapshot := func() ([]byte, error) { return kvs.getSnapshot() }
	commitC, errorC, snapshotterReady := newRaftNode(cfg.RaftId, strings.Split(cfg.RaftCluster, ","), cfg.RaftJoin, getSnapshot, proposeC, confChangeC)

	kvs = newPersistentStore(cfg, <-snapshotterReady, proposeC, commitC, errorC)
	// kvs = newKVStore(<-snapshotterReady, proposeC, commitC, errorC)
	go func() {
		serveHttpKVAPI(kvs, cfg.RaftPort, confChangeC, errorC)
	}()
	return kvs, nil
}
