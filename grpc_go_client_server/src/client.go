package main

import (
	"context"
	"fmt"
	"log"
	"time"

	greeter "my-grpc-project/greeter"

	"google.golang.org/grpc"
)

// func main() {
// 	// Establish a connection to the server
// 	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()

// 	// Create a Greeter client
// 	c := greeter.NewGreeterClient(conn)

// 	// Call the SayHello method
// 	name := "World"
// 	resp, err := c.SayHello(context.Background(), &greeter.HelloRequest{Name: name})
// 	if err != nil {
// 		log.Fatalf("could not greet: %v", err)
// 	}

// 	// Print the server response
// 	fmt.Printf("Greeting: %s\n", resp.GetMessage())
// }

func main() {

	// Set up a connection to the server.

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {

		log.Fatalf("did not connect: %v", err)

	}

	defer conn.Close()
	// Create a new User1 client
	client := greeter.NewGreeterClient(conn)
	
	// Prepare the request

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the GetUsers method

	response, err := client.GetUsers(ctx, &greeter.GetUsersRequest{})
	if err != nil {
		log.Fatalf("could not get users: %v", err)
	}
	// Print the response
	for _, user := range response.GetUsers() {
		fmt.Printf(":User  %s, Name: %s, Email: %s\n", user.GetId(), user.GetName(), user.GetEmail())
		fmt.Printf("Password: %s\n", user.GetPassword())
	}
}
