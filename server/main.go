package main

import (
	"log"
	"net"

	pb "github.com/doutori/grpc-sample/pb"
	"github.com/doutori/grpc-sample/server/services"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Starting Server!")

	s := grpc.NewServer()
	catService := &services.MyCatService{}

	// 実行したい実処理をseverに登録する
	pb.RegisterCatServer(s, catService)

	s.Serve(listenPort)
}
