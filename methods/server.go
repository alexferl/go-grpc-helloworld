package methods

import pb "google.golang.org/grpc/examples/helloworld/helloworld"

// Server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedGreeterServer
}
