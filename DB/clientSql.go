package main

import (
	pb "DB/grpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("10.243.50.4:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常：%s\n", err)
	}
	fmt.Println("启动更改配置")
	defer conn.Close()
	//client := pb.NewSqlDefaultServiceClient(conn
	client1 := pb.NewUpdateAlertingConfigClient(conn)
	req := new(pb.SqlRequest)
	req.MaxValueCpu = 50

	req.MaxValueDisk = 79

	response, err := client1.GetSqlInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("响应异常%s\n", err)
	}
	fmt.Printf("响应结果%s\n", response)

}
