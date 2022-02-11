package main

import (
	"DB/common"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("10.243.105.17:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常：%s\n", err)
	}
	//fmt.Println("启动更改配置")
	defer conn.Close()
	c := common.NewEmailServiceClient(conn)
	r, err := c.SendEmail(context.Background(), &common.GetEmailRequest{
		Sender:    "1834960035@qq.com",
		Recipient: "stackoverflow520@163.com",
		CpuUsed:   "%30",
		MemUsed:   "%40",
		Grade:     "serious",
	})
	if err != nil {
		log.Fatalf("could not sendEmail: %v", err)
	}
	fmt.Println(r.Code)
	fmt.Println(r.Info)
}
