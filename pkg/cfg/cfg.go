package cfg

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Cfg cfg
type Cfg struct {
	API   APICfg   `toml:"api"`
	Store StoreCfg `toml:"store"`
}

// APICfg api cfg
type APICfg struct {
	Addr string `toml:"addr"`
}

// StoreCfg store cfg
type StoreCfg struct {
	Memory   bool   `toml:"memory"`
	DataPath string `toml:"dataPath"`

	RaftAddr    string `toml:"raftAddr"`
	RaftId      int    `toml:"raftId"`
	RaftCluster string `toml:"raftCluster"`
	RaftPort    int    `toml:"raftPort"`
	RaftJoin    bool   `toml:"raftJoin"`
}

// MustParseCfg parse cfg, exit if has any error
func MustParseCfg(file string) Cfg {
	c := Cfg{}
	_, err := toml.DecodeFile(file, &c)
	if err != nil {
		log.Fatalf("parse cfg %s failed with %+v", file, err)
	}
	return c
}
