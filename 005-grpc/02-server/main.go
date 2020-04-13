package main

import (
	"context"
	"fmt"
	"net"

	"github.com/dearrudam/udemy-course-go-project/005-grpc/02-server/echo"
	grpc "google.golang.org/grpc"
)

// EchoServer is the implementation of the gRPC service
type EchoServer struct{}

// Echo method is required for the gRPC EchoService Interface
func (s *EchoServer) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	fmt.Printf("Receiving EchoRequest = %v\n", req.GetMessage())
	return &echo.EchoResponse{
		Response: fmt.Sprintf("My Echo: %v", req.GetMessage()),
	}, nil
}
func main() {

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	srv := &EchoServer{}

	echo.RegisterEchoServiceServer(s, srv)

	fmt.Printf("Now serving at %v\n", listener.Addr().String())
	err = s.Serve(listener)
	if err != nil {
		panic(err)
	}

}
