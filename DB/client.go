package main

import (
	pb "DB/grpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

type admin struct {
	id    int64
	name  string
	emali string
}

func main() {
	conn, err := grpc.Dial("10.243.50.4:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常：%s\n", err)
	}
	defer conn.Close()
	//接入服务客户端
	client := pb.NewUserInfoServiceClient(conn)
	req := new(pb.UserRequest)
	req.Id = 1
	req.Cpuused = 98
	req.Memused = 32.67
	req.DiskUsed = 32.98
	req.Timeout = time.Now().UnixNano() / 1e6
	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("异常%s", err)
	}
	fmt.Println(resp)

	//管理服务管理员客户端
	client1 := pb.NewUpdateAlertingConfigClient(conn)
	guanliyuan := []admin{
		//{1920, "wu", "569105057@qq.com"},
		{23121, "jiahui", "106463499"},
		{5823, "xiaoshi", "askhfokj"},
	}

	name := []string{}
	email := []string{}
	id := []int64{}
	for _, a := range guanliyuan {
		name = append(name, a.name)
		email = append(email, a.emali)
		id = append(id, a.id)
	}
	req1 := new(pb.AdminRequest)
	req1.Name = name
	req1.Email = email
	req1.Id = id
	response1, err := client1.GetAdminInfo(context.Background(), req1)
	if err != nil {
		fmt.Printf("响应异常%s\n", err)
	}
	fmt.Printf("响应结果%s\n", response1)

	//管理服务配置客户端
	//client2 := pb.NewUpdateAlertingConfigClient(conn)
	req2 := new(pb.SqlRequest)
	req2.MaxValueCpu = 40

	req2.MaxValueDisk = 79
	response2, err := client1.GetSqlInfo(context.Background(), req2)
	if err != nil {
		fmt.Printf("响应异常%s\n", err)
	}
	fmt.Printf("响应结果%s\n", response2)
}
