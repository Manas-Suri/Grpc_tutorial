package main

import (
	"context"
	"fmt"
	"log"
	"time"

	greeter "grpc_go_client_server/greeter"

	"google.golang.org/grpc"
)


func main() {

	fmt.Println("Will try to greet the world")
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
