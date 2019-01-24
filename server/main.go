package main

import (
	"log"
	"net"

	pb "github.com/doutori/grpc-sample/api/pb"
	"github.com/doutori/grpc-sample/api/services"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	catService := &services.MyCatService{}

	// 実行したい実処理をseverに登録する
	pb.RegisterCatServer(s, catService)

	s.Serve(listenPort)
}
