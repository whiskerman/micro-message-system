module micro-message-system

go 1.16

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/asim/go-micro/plugins/broker/kafka/v3 v3.0.0-20210924081004-8c39b1e1204d
	github.com/asim/go-micro/plugins/registry/etcd/v3 v3.0.0-20210924081004-8c39b1e1204d
	github.com/asim/go-micro/plugins/registry/kubernetes/v3 v3.0.0-20210924081004-8c39b1e1204d
	github.com/asim/go-micro/plugins/transport/grpc/v3 v3.0.0-20210924081004-8c39b1e1204d
	github.com/asim/go-micro/plugins/wrapper/breaker/hystrix/v3 v3.0.0-20210924081004-8c39b1e1204d
	github.com/asim/go-micro/plugins/wrapper/ratelimiter/ratelimit/v3 v3.0.0-20210924081004-8c39b1e1204d
	github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3 v3.0.0-20210924081004-8c39b1e1204d
	github.com/asim/go-micro/v3 v3.6.1-0.20210831082736-088ccb50019c
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.4
	github.com/go-acme/lego/v4 v4.4.0
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v1.8.5
	github.com/gorilla/websocket v1.4.2
	github.com/jinzhu/gorm v1.9.16
	github.com/juju/ratelimit v1.0.1
	github.com/lib/pq v1.10.3 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/olivere/elastic v6.2.37+incompatible
	github.com/opentracing/opentracing-go v1.2.0
	github.com/satori/go.uuid v1.2.0
	github.com/smartystreets/assertions v1.1.1 // indirect
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/urfave/cli/v2 v2.3.0
	google.golang.org/genproto v0.0.0-20210821163610-241b8fcbd6c8 // indirect
	google.golang.org/grpc v1.40.0
	gopkg.in/go-playground/validator.v8 v8.18.2
)
