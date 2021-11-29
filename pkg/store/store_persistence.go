package store

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"sync"

	"github.com/cockroachdb/pebble"
	"github.com/matrixorigin/talent-challenge/matrixbase/distributed/pkg/cfg"
	"go.etcd.io/etcd/raft/v3/raftpb"
	"go.etcd.io/etcd/server/v3/etcdserver/api/snap"
)

// TO DO: close db connection

type kv struct {
	Key string
	Val string
}

type persistentStore struct {
	mu sync.RWMutex

	proposeC    chan<- string // channel for proposing updates
	db          *pebble.DB
	snapshotter *snap.Snapshotter
}

func newPersistentStore(
	cfg cfg.StoreCfg,
	snapshotter *snap.Snapshotter,
	proposeC chan<- string,
	commitC <-chan *commit,
	errorC <-chan error) Store {

	db, err := pebble.Open(cfg.DataPath, &pebble.Options{})
	if err != nil {
		return nil
	}
	s := &persistentStore{
		db:          db,
		snapshotter: snapshotter,
		proposeC:    proposeC,
	}

	snapshot, err := s.loadSnapshot()
	if err != nil {
		log.Panic(err)
	}
	if snapshot != nil {
		log.Printf("loading snapshot at term %d and index %d", snapshot.Metadata.Term, snapshot.Metadata.Index)
		if err := s.recoverFromSnapshot(snapshot.Data); err != nil {
			log.Panic(err)
		}
	}
	// read commits from raft into kvStore map until error
	go s.readCommits(commitC, errorC)

	return s
}

func (s *persistentStore) Set(key []byte, value []byte) error {
	log.Printf("%s %s %s\n", "set", key, value)
	return s.Propose("set", string(key), string(value))
}

func (s *persistentStore) Get(key []byte) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, closer, err := s.db.Get(key)
	if err != nil {
		return nil, err
	}

	log.Printf("%s %s %s\n", "get", key, value)
	if err := closer.Close(); err != nil {
		log.Fatal(err)
	}

	return value, nil
}

func (s *persistentStore) Delete(key []byte) error {
	log.Printf("deleting %s\n", key)
	return s.Propose("del", string(key), string(""))
}

func (s *persistentStore) getSnapshot() ([]byte, error) {
	// s.mu.RLock()
	// defer s.mu.RUnlock()
	snapshot := s.db.NewSnapshot()
	defer snapshot.Close()
	options := make(map[string]string)
	// Count the keys at this snapshot.
	iter := snapshot.NewIter(nil)
	count := 0
	for iter.First(); iter.Valid(); iter.Next() {
		count++
		log.Printf("getSnapshot  No. %d : %s %v", count, iter.Key(), iter.Value())
		options[string(iter.Key())] = string(iter.Value())
	}

	return json.Marshal(options)
}

func (s *persistentStore) loadSnapshot() (*raftpb.Snapshot, error) {
	snapshot, err := s.snapshotter.Load()
	if err == snap.ErrNoSnapshot {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return snapshot, nil
}

func (s *persistentStore) Propose(ops string, k string, v string) error {
	var buf bytes.Buffer
	options := make(map[string]kv)
	options[ops] = kv{k, v}
	if err := gob.NewEncoder(&buf).Encode(options); err != nil {
		return err
	}
	s.proposeC <- buf.String()
	return nil
}

func (s *persistentStore) readCommits(commitC <-chan *commit, errorC <-chan error) {
	for commit := range commitC {
		if commit == nil {
			// signaled to load snapshot
			snapshot, err := s.loadSnapshot()
			if err != nil {
				log.Panic(err)
			}
			if snapshot != nil {
				log.Printf("loading snapshot at term %d and index %d", snapshot.Metadata.Term, snapshot.Metadata.Index)
				if err := s.recoverFromSnapshot(snapshot.Data); err != nil {
					log.Panic(err)
				}
			}
			continue
		}

		for _, data := range commit.data {
			options := make(map[string]kv)
			dec := gob.NewDecoder(bytes.NewBufferString(data))
			if err := dec.Decode(&options); err != nil {
				log.Fatalf("raftexample: could not decode message (%v)", err)
			}

			for ops, kv := range options {
				if ops == "set" {
					log.Printf("readCommits setting %s %s\n", kv.Key, kv.Val)
					if err := s.db.Set([]byte(kv.Key), []byte(kv.Val), pebble.Sync); err != nil {
						log.Panic(err)
					}
				}
				if ops == "del" {
					if err := s.db.Delete([]byte(kv.Key), pebble.Sync); err != nil {
						log.Panic(err)
					}
				}
			}
		}
		close(commit.applyDoneC)
	}
	if err, ok := <-errorC; ok {
		log.Fatal(err)
	}
}

func (s *persistentStore) recoverFromSnapshot(snapshot []byte) error {
	var store map[string]string
	if err := json.Unmarshal(snapshot, &store); err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	for k, v := range store {
		if err := s.db.Set([]byte(k), []byte(v), pebble.Sync); err != nil {
			return (err)
		}
	}
	return nil
}
