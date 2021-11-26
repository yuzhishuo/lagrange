module github.com/matrixorigin/talent-challenge/matrixbase/distributed

go 1.15

replace go.etcd.io/etcd/server/v3 v3.5.1 => /home/luluyuzhi/go/src/github.com/etcd-io/etcd/server

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/cockroachdb/pebble v0.0.0-20211125005712-9791c0f4c052
	github.com/gin-gonic/gin v1.6.3
	github.com/stretchr/testify v1.7.0
	github.com/valyala/fasthttp v1.19.0
	go.etcd.io/etcd/client/pkg/v3 v3.5.1
	go.etcd.io/etcd/raft/v3 v3.5.1
	go.etcd.io/etcd/server/v3 v3.5.1
	go.uber.org/zap v1.17.0
)
