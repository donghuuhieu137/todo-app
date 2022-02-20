package main

import (
	"fmt"
	"log"
	"net"

	"todo-app/handlers"
	"todo-app/proto/todo"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50001")

	h := handlers.New()

	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()
	todo.RegisterTodoServiceServer(s, h)

	fmt.Println("Gateway service is running...")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
