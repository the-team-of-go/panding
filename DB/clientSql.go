package main

import (
	pb "DB/grpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常：%s\n", err)
	}
	fmt.Println("启动更改配置")
	defer conn.Close()
	client := pb.NewSqlDefaultServiceClient(conn)
	req := new(pb.SqlRequest)
	req.MaxValueCpu = 50
	req.MinValueCpu = 20
	req.AvergValueCpue = 45
	req.MaxValueMem = 50
	req.MinValueMem = 20
	req.AvergValueMem = 45
	req.MaxValueDisk = 90
	req.Timeout = 0
	response, err := client.GetSqlInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("响应异常%s\n", err)
	}
	fmt.Printf("响应结果%s\n", response)

}
