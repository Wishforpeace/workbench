module workbench-gateway

//replace workbench-project => ../project

replace workbench => ../../

replace workbench-user => ../user

go 1.16

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/micro/go-micro => github.com/Lofanmi/go-micro v1.16.1-0.20210804063523-68bbf601cfa4 // to go 1.16

require (
	github.com/gin-gonic/gin v1.9.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/kubernetes v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/wrapper/trace/opentracing v0.0.0-20200119172437-4fe21aa238fd
	github.com/satori/go.uuid v1.2.0
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.12.0
	github.com/swaggo/files v0.0.0-20220728132757-551d4a08d97a
	github.com/swaggo/gin-swagger v1.5.2
	github.com/swaggo/swag v1.8.1
	github.com/teris-io/shortid v0.0.0-20220617161101-71ec9f2aa569
	github.com/willf/pad v0.0.0-20200313202418-172aa767f2a4
	go.uber.org/zap v1.23.0
	workbench v0.0.0-00010101000000-000000000000
	workbench-user v0.0.0-00010101000000-000000000000
)

require (
	github.com/opentracing/opentracing-go v1.2.0
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
)
