package main

import (
	pb "DB/grpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
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
	fmt.Println("启动更改配置")
	defer conn.Close()
	//client := pb.NewAdminGetServiceClient(conn)
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
	req := new(pb.AdminRequest)
	req.Name = name
	req.Email = email
	req.Id = id
	response, err := client1.GetAdminInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("响应异常%s\n", err)
	}
	fmt.Printf("响应结果%s\n", response)

}
