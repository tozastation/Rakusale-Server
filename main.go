package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listenPort, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	// catService := &service.MyCatService{}
	// // 実行したい実処理をseverに登録する
	// pb.RegisterCatServer(server, catService)
	server.Serve(listenPort)
}
