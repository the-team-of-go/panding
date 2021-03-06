package predict

import (
	"DB/common"
	"DB/model"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

func Predict(a model.Staus) {
	if model.Sending {
		if a.CpuUsed >= float64(model.SqlAlteringConfig.MaxValueCPU) || a.MenUsed > float64(model.SqlAlteringConfig.MaxValueMem) || a.DiskUsed > float64(model.SqlAlteringConfig.MaxValueDisk) {
			fmt.Printf("danger,cpuUsed:%f memUsed:%f diskUsed:%f\n", a.CpuUsed, a.MenUsed, a.DiskUsed)
			conn, err := grpc.Dial("10.243.105.17:8181", grpc.WithInsecure())
			if err != nil {
				fmt.Printf("连接异常：%s\n", err)
			}
			//fmt.Println("启动更改配置")
			defer conn.Close()
			c := common.NewEmailServiceClient(conn)
			for i := 0; i < len(model.Email); i++ {
				model.Wg.Add(1)
				go func(i int) {
					defer model.Wg.Done()
					_, err := c.SendEmail(context.Background(), &common.GetEmailRequest{
						Sender:    "1834960035@qq.com",
						Recipient: model.Email[i],
						CpuUsed:   strconv.FormatFloat(a.CpuUsed, 'f', 4, 64),
						MemUsed:   strconv.FormatFloat(a.MenUsed, 'f', 4, 64),
						DiskUsed:  strconv.FormatFloat(a.DiskUsed, 'f', 4, 64),
						Timestamp: a.Timeout,
						Grade:     "danger",
					})
					if err != nil {
						//log.Fatalf("could not sendEmail: %v", err)
						fmt.Printf("sendingerr:%s\n", err)
					}
					//fmt.Println(r.Code)
					//fmt.Println(r.Info)
				}(i)
			}
			go func() {
				model.Wg.Add(1)
				fmt.Println(1)
				model.Sending = false
				time.Sleep(time.Duration(2) * time.Minute)
				model.Sending = true
				fmt.Println(2)
				defer model.Wg.Done()
			}()
		} else if a.CpuUsed >= 65 || a.MenUsed > 65 || a.DiskUsed > 65 {
			fmt.Printf("servious,cpuUsed:%f memUsed:%f diskUsed%f\n", a.CpuUsed, a.MenUsed, a.DiskUsed)
			conn, err := grpc.Dial("10.243.105.17:8181", grpc.WithInsecure())
			if err != nil {
				fmt.Printf("连接异常：%s\n", err)
			}
			//fmt.Println("启动更改配置")
			defer conn.Close()
			c := common.NewEmailServiceClient(conn)
			for i := 0; i < len(model.Email); i++ {
				model.Wg.Add(1)
				go func(i int) {
					defer model.Wg.Done()
					_, err := c.SendEmail(context.Background(), &common.GetEmailRequest{
						Sender:    "1834960035@qq.com",
						Recipient: model.Email[i],
						CpuUsed:   strconv.FormatFloat(a.CpuUsed, 'f', 4, 64),
						MemUsed:   strconv.FormatFloat(a.MenUsed, 'f', 4, 64),
						DiskUsed:  strconv.FormatFloat(a.DiskUsed, 'f', 4, 64),
						Timestamp: a.Timeout,
						Grade:     "danger",
					})
					if err != nil {
						//log.Fatalf("could not sendEmail: %v", err)
						fmt.Printf("sendingerr:%s\n", err)
					}
					//fmt.Println(r.Code)
					//fmt.Println(r.Info)
				}(i)
			}
			go func() {
				model.Wg.Add(1)
				model.Sending = false
				time.Sleep(time.Duration(2) * time.Minute)
				model.Sending = true
				defer model.Wg.Done()
			}()
		} else {
			fmt.Printf("normal,cpuUsed:%f memUsed:%f diskUsed:%f\n", a.CpuUsed, a.MenUsed, a.DiskUsed)
			//fmt.Println("启动更改配置")
		}
		model.Wg.Wait()
	}

}
