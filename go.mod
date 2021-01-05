module boxin

go 1.14

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.3.0
)

require (
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fatih/color v1.10.0 // indirect
	github.com/garyburd/redigo v1.6.2
	github.com/gin-gonic/gin v1.6.3
	github.com/go-gomail/gomail v0.0.0-20160411212932-81ebce5c23df
	github.com/go-siris/siris v7.4.0+incompatible // indirect
	github.com/golang/protobuf v1.4.3
	github.com/jinzhu/gorm v1.9.16
	github.com/labstack/echo v3.3.10+incompatible // indirect
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/oxequa/interact v0.0.0-20171114182912-f8fb5795b5d7 // indirect
	github.com/oxequa/realize v2.0.2+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a
	github.com/valyala/fasttemplate v1.2.1 // indirect
	go.mongodb.org/mongo-driver v1.4.4
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	google.golang.org/protobuf v1.25.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/urfave/cli.v2 v2.0.0-00010101000000-000000000000 // indirect
)
