package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/doutori/grpc-sample/pb"
	"google.golang.org/grpc"
)

func main() {
	//sampleなのでwithInsecure
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()

	client := pb.NewCatClient(conn)
	message := &pb.GetMyCatMessage{"tama"}
	res, err := client.GetMyCat(context.TODO(), message)

	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
