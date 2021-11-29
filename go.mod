module github.com/matrixorigin/talent-challenge/matrixbase/distributed

go 1.17

// replace go.etcd.io/etcd/server/v3 v3.5.1 => /home/luluyuzhi/go/src/github.com/etcd-io/etcd/server
replace (
	go.etcd.io/etcd/api/v3 => /home/luluyuzhi/go/src/github.com/etcd-io/etcd/api
	go.etcd.io/etcd/client/pkg/v3  => /home/luluyuzhi/go/src/github.com/etcd-io/etcd/client/pkg
	go.etcd.io/etcd/client/v2  => /home/luluyuzhi/go/src/github.com/etcd-io/etcd/client/v2
	go.etcd.io/etcd/client/v3  => /home/luluyuzhi/go/src/github.com/etcd-io/etcd/client/v3
	go.etcd.io/etcd/etcdctl/v3   => /home/luluyuzhi/go/src/github.com/etcd-io/etcd/etcdctl
	go.etcd.io/etcd/etcdutl/v3   => /home/luluyuzhi/go/src/github.com/etcd-io/etcd/etcdutl
	go.etcd.io/etcd/pkg/v3   => /home/luluyuzhi/go/src/github.com/etcd-io/etcd/pkg
	go.etcd.io/etcd/raft/v3   => /home/luluyuzhi/go/src/github.com/etcd-io/etcd/raft
	go.etcd.io/etcd/server/v3   => /home/luluyuzhi/go/src/github.com/etcd-io/etcd/server
	go.etcd.io/etcd/tests/v3   => /home/luluyuzhi/go/src/github.com/etcd-io/etcd/tests
)

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

require (
	github.com/DataDog/zstd v1.4.5 // indirect
	github.com/andybalholm/brotli v1.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/cockroachdb/errors v1.8.1 // indirect
	github.com/cockroachdb/logtags v0.0.0-20190617123548-eb05cc24525f // indirect
	github.com/cockroachdb/redact v1.0.8 // indirect
	github.com/cockroachdb/sentry-go v0.6.1-cockroachdb.2 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.2.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/klauspost/compress v1.11.7 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.26.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	go.etcd.io/etcd/api/v3 v3.5.0 // indirect
	go.etcd.io/etcd/pkg/v3 v3.5.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/exp v0.0.0-20200513190911-00229845015e // indirect
	golang.org/x/net v0.0.0-20211123203042-d83791d6bcd9 // indirect
	golang.org/x/sys v0.0.0-20211123173158-ef496fb156ab // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	google.golang.org/genproto v0.0.0-20211118181313-81c1377c94b1 // indirect
	google.golang.org/grpc v1.41.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
