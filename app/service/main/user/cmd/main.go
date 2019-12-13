package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"mall/app/service/main/account/api"
)

func main() {
	// TODO use client from account api from http request
	conn, err := grpc.Dial("localhost:9013", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("what htefuc %s ", err.Error())
	}
	defer conn.Close()
	client := api.NewAccountClient(conn)
	c := context.Background()
	a := &api.Id{Id: 11}
	f, err := client.Read(c, a)
	if err != nil {
		fmt.Errorf("error %s ", err.Error())
	}
	fmt.Println(f)
	fmt.Println(f.GetModel().GetId())
}
