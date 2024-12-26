package main

import (
	"context"
	"fmt"
	"log"
	greeter "my-grpc-project/greeter"
	"net"

	"google.golang.org/grpc"
)

// server is used to implement greeter.GreeterServer
type server struct {
	greeter.UnimplementedGreeterServer
}

// SayHello implements greeter.GreeterServer
func (s *server) SayHello(ctx context.Context, in *greeter.HelloRequest) (*greeter.HelloReply, error) {
	return &greeter.HelloReply{Message: "Hello, " + in.GetName()}, nil
}




func (s *server) GetUsers(ctx context.Context, in *greeter.GetUsersRequest) (*greeter.GetUsersResponse, error){
	users := []*greeter.User{
		{
			Id:        "1",
			Name:      "John Doe",
			Email:     "manassuri98.ms",
			Password:  "manas123",
			Phone:     "123-456-7890",
			Address:   "123 Main St",
			Role:      "user",
			Status:    "active",
			CreatedAt: "2024-08-09",
			UpdatedAt: "2024-08-09",
		},
		{
			Id:        "2",
			Name:      "Manas Suri",
			Email:     "manassuri98.ms",
			Password:  "manas1234556",
			Phone:     "123-456-7890",
			Address:   "123 Main St",
			Role:      "user",
			Status:    "active",
			CreatedAt: "2024-08-09",
			UpdatedAt: "2024-08-09",
		},

	}
	return &greeter.GetUsersResponse{Users: users},nil
}


func main() {
	// Create a TCP listener
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server
	s := grpc.NewServer()
	greeter.RegisterGreeterServer(s, &server{})

	// Start the server
	fmt.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
