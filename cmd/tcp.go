package main

import (
	"fmt"
	"github.com/wywwwwei/IMServer/Service"
	"log"
	"net"
)

func setTCP(){
	address := fmt.Sprintf("%s:%d",Service.SERVER__IP,Service.SERVER_TCP_PORT)
	addr,err:=net.ResolveTCPAddr("tcp",address)
	fmt.Println(addr)
	if err!=nil{
		log.Fatalf("Create tcp socket failed: %s",err)
	}
	listener,err:=net.ListenTCP("tcp",addr)
	if err!=nil{
		log.Fatalf("Listen tcp port failed: %s",err)
	}
	defer listener.Close()
	fmt.Printf("Listening tcp on port:%d\n",Service.SERVER_TCP_PORT)
	for{
		conn,err := listener.Accept()
		if err!=nil{
			log.Printf("Accept tcp error: %s",err)
		}
		go Service.TcpHandler(conn)
	}
}
