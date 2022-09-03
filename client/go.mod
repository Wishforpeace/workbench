module workbench-client

replace workbench => ../

replace workbench-user => ../microservice/user

go 1.16

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/micro/go-micro => github.com/Lofanmi/go-micro v1.16.1-0.20210804063523-68bbf601cfa4 // to go 1.16

require (
	github.com/micro/go-micro v1.18.0
	workbench-user v0.0.0-00010101000000-000000000000
)
