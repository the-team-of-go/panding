package main

import (
	pb "DB/grpc/proto"
	"DB/server"
	"fmt"
	"github.com/robfig/cron"
	"google.golang.org/grpc"
	"net"
)

func main() {

	go server.HttpServer()
	c := cron.New()                          // 新建一个定时任务对象
	c.AddFunc("0 */5 * * * ?", server.Dayin) // 给对象增加定时任务
	c.Start()
	addr := "10.243.50.4:8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常%s\n", err)
	}
	fmt.Printf("监听gprc端口%s\n", addr)
	s := grpc.NewServer()
	pb.RegisterUserInfoServiceServer(s, &server.U)
	//pb.RegisterSqlDefaultServiceServer(s, &sql)
	//pb.RegisterAdminGetServiceServer(s, &ad)
	pb.RegisterUpdateAlertingConfigServer(s, &server.Se)

	s.Serve(listener)
}
