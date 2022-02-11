package main

import (
	pb "DB/grpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type admin struct {
	name  string
	emali string
}

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常：%s\n", err)
	}
	fmt.Println("启动更改配置")
	defer conn.Close()
	client := pb.NewAdminGetServiceClient(conn)
	guanliyuan := []admin{
		{"wu", "569105057@qq.com"},
		{"jiahui", "106463499"},
		{"xiaqian", "8010190301"},
	}
	name := []string{}
	email := []string{}
	for _, a := range guanliyuan {
		name = append(name, a.name)
		email = append(email, a.emali)
	}
	req := new(pb.AdminRequest)
	req.Name = name
	req.Email = email
	response, err := client.GetAdminInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("响应异常%s\n", err)
	}
	fmt.Printf("响应结果%s\n", response)

}
