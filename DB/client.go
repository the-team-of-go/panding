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
	defer conn.Close()

	client := pb.NewUserInfoServiceClient(conn)
	req := new(pb.UserRequest)
	req.Id = 1
	req.Cpuused = 19
	req.Memused = 49
	req.DiskUsed = 90
	req.Timeout = 0
	client.GetUserInfo(context.Background(), req)

	req1 := new(pb.UserRequest)
	req1.Id = 2
	req1.Cpuused = 19
	req1.Memused = 59
	req1.DiskUsed = 80
	req1.Timeout = 0
	client.GetUserInfo(context.Background(), req1)

	req2 := new(pb.UserRequest)
	req2.Id = 4
	req2.Cpuused = 89
	req2.Memused = 19
	req2.DiskUsed = 82
	req2.Timeout = 0
	client.GetUserInfo(context.Background(), req2)
	//if err != nil {
	//	fmt.Printf("异常%s", err)
	//}
	//fmt.Println(resp)
}
