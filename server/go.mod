module github.com/kurtosis-tech/example-api-server/server

go 1.15

require (
	github.com/kurtosis-tech/example-api-server/api/golang v0.0.0-20211101150102-edd2e13c322a
	github.com/kurtosis-tech/example-datastore-server/api/golang v0.0.0-20211027202930-da3ee244008e
	github.com/kurtosis-tech/minimal-grpc-server/golang v0.0.0-20210921153930-d70d7667c51b
	github.com/palantir/stacktrace v0.0.0-20161112013806-78658fd2d177
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)
