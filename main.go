package main

import (
	"fmt"
	"grpc-learn/connection"
	"grpc-learn/protobuf"
	"grpc-learn/service/bank"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main(){
	netListen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("Failed To Start ::" + err.Error())
	}

	serverGrpc := grpc.NewServer()

	db := connection.ConnectionDB()

	serviceBank := bank.BankService{
		DB: db,
	}
	protobuf.RegisterBankServiceServer(serverGrpc, &serviceBank)

	fmt.Println("Connected at ", netListen.Addr())

	if err := serverGrpc.Serve(netListen); err != nil {
		log.Fatal("Cant Connect Server " + err.Error())
	}
}
