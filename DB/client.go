package main

import (
	pb "DB/grpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

func main() {
	conn, err := grpc.Dial("10.243.50.4:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常：%s\n", err)
	}
	defer conn.Close()

	client := pb.NewUserInfoServiceClient(conn)
	req := new(pb.UserRequest)
	req.Id = 4
	req.Cpuused = 79.21
	req.Memused = 29.43
	req.DiskUsed = 54.23
	req.Timeout = time.Now().UnixNano() / 1e6
	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("异常%s", err)
	}
	fmt.Println(resp)
	req1 := new(pb.UserRequest)
	req1.Id = 1
	req1.Cpuused = 69.12
	req1.Memused = 23.89
	req1.DiskUsed = 39.12
	req1.Timeout = time.Now().UnixNano() / 1e6
	client.GetUserInfo(context.Background(), req1)

	req2 := new(pb.UserRequest)
	req2.Id = 2
	req2.Cpuused = 79.12
	req2.Memused = 43.998
	req2.DiskUsed = 59.12
	req2.Timeout = time.Now().UnixNano() / 1e6
	client.GetUserInfo(context.Background(), req2)
	//time.Sleep(time.Duration(10) * time.Second)

}
