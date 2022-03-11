# go-grpc-helloworld [![Go Report Card](https://goreportcard.com/badge/github.com/alexferl/go-grpc-helloworld)](https://goreportcard.com/report/github.com/alexferl/go-grpc-helloworld)

A small boilerplate gRPC app based on the Go [helloworld](https://github.com/grpc/grpc-go/tree/master/examples/helloworld)
with [12-factor](https://12factor.net/).

### Building & Running locally
```shell script
make run
```

### Testing
With [gRPCurl](https://github.com/fullstorydev/grpcurl):

```shell
echo '{"name": "World"}' | grpcurl -plaintext -d @ localhost:50051 helloworld.Greeter/SayHello
{
  "message": "Hello World"
}
```

Health check:

```shell
grpcurl -plaintext localhost:50052 grpc.health.v1.Health/Check
{
  "status": "SERVING"
}
```

### Usage
```shell
Usage of ./server:
      --app-name string          The name of the application. (default "app")
      --bind-address ip          The IP address to listen at. (default 127.0.0.1)
      --bind-port uint           The port to listen at. (default 50051)
      --env-name string          The environment of the application. Used to load the right configs file. (default "local")
      --graceful-timeout uint    Timeout for graceful shutdown. (default 30)
      --health-bind-address ip   The IP address to listen at. (default 127.0.0.1)
      --health-bind-port uint    The port to listen at. (default 50052)
      --log-level string         The granularity of log outputs. Valid log levels: 'panic', 'fatal', 'error', 'warn', 'info', 'debug' and 'trace'. (default "info")
      --log-output string        The output to write to. 'stdout' means log to stdout, 'stderr' means log to stderr. (default "stdout")
      --log-writer string        The log writer. Valid writers are: 'console' and 'json'. (default "console")
```

### Building & Running Docker container
```shell
make build-docker
make run-docker
```
