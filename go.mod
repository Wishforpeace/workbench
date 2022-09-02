module github.com/Muxi-Backend-Classroom/workbench

replace workbench => ./

go 1.16

require (
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro v1.18.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/smartystreets/goconvey v1.7.2
	github.com/spf13/viper v1.12.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	go.uber.org/zap v1.23.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	workbench v0.0.0-00010101000000-000000000000
)

require github.com/dgrijalva/jwt-go v3.2.0+incompatible
