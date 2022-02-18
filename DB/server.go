package main

import (
	"DB/dao"
	pb "DB/grpc/proto"
	"DB/model"
	"context"
	"fmt"
	"github.com/robfig/cron"
	"google.golang.org/grpc"
	"net"
)

type UserInfoService struct{}

var u = UserInfoService{}

func (u *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	id := req.Id
	cpuUsed := req.Cpuused
	memUsed := req.Memused
	diskUsed := req.DiskUsed
	timeout := req.Timeout
	var grade string
	if cpuUsed >= float64(model.SqlAlteringConfig.MaxValueCPU) || memUsed >= float64(model.SqlAlteringConfig.MaxValueMem) || diskUsed >= float64(model.SqlAlteringConfig.MaxValueDisk) {
		grade = "danger"
	} else if cpuUsed >= 65 || memUsed >= 65 || diskUsed >= 65 {
		grade = "serious"
	} else {
		grade = "normal"
	}

	temp := model.Staus{int(id), float64(cpuUsed), float64(memUsed), float64(diskUsed), grade, timeout}
	go dao.AddNums(temp)
	go dao.InsertNode(temp)
	//model.LastestNodeStatus[temp.Id] = temp
	//fmt.Println(model.LastestNodeStatus)
	//fmt.Println(temp)
	resp = &pb.UserResponse{Code: 1, Mesg: "ok"}
	return
}

type ser struct {
	pb.UnimplementedUpdateAlertingConfigServer
}

var se = ser{}

//type SqlInfoService struct{}

//var sql = SqlInfoService{}

func (s *ser) GetSqlInfo(ctx context.Context, req *pb.SqlRequest) (resp *pb.SqlResponse, err error) {
	fmt.Println("配置更改")
	if req.Id != 0 {
		model.SqlAlteringConfig.Id = req.Id
	}
	if req.Name != "" {
		model.SqlAlteringConfig.Name = req.Name
	}
	if req.MaxValueCpu != 0 {
		model.SqlAlteringConfig.MaxValueCPU = req.MaxValueCpu
	}
	if req.MaxValueMem != 0 {
		model.SqlAlteringConfig.MaxValueMem = req.MaxValueMem
	}
	if req.MaxValueDisk != 0 {
		model.SqlAlteringConfig.MaxValueDisk = req.MaxValueDisk
	}
	if req.PeriodTime != 0 {
		model.SqlAlteringConfig.Timeout = req.PeriodTime
	}
	fmt.Println(model.SqlAlteringConfig)
	var data []string
	resp = &pb.SqlResponse{Code: 1, Msg: "ok", Data: data}
	return
}

//type AdminInfoService struct{}
//
//var ad = AdminInfoService{}

func (s *ser) GetAdminInfo(ctx context.Context, req *pb.AdminRequest) (resp *pb.AdminResponse, err error) {
	model.Admin = []string{}
	model.Email = []string{}
	model.Id = []int64{}
	for i := 0; i < len(req.Name); i++ {
		model.Admin = append(model.Admin, req.Name[i])
	}
	for i := 0; i < len(req.Email); i++ {
		model.Email = append(model.Email, req.Email[i])
	}
	for i := 0; i < len(req.Id); i++ {
		model.Id = append(model.Id, req.Id[i])
	}
	fmt.Println(model.Admin)
	fmt.Println(model.Email)
	fmt.Println(model.Id)
	resp = &pb.AdminResponse{Code: 1, Msg: "ok"}
	return
}

func dayin() {
	for _, v := range model.Node {
		go dao.Aggregation(v)
	}

}

func main() {
	c := cron.New()                   // 新建一个定时任务对象
	c.AddFunc("0 */5 * * * ?", dayin) // 给对象增加定时任务
	c.Start()

	addr := "10.243.50.4:8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常%s\n", err)
	}
	fmt.Printf("监听端口%s\n", addr)
	s := grpc.NewServer()
	pb.RegisterUserInfoServiceServer(s, &u)
	//pb.RegisterSqlDefaultServiceServer(s, &sql)
	//pb.RegisterAdminGetServiceServer(s, &ad)
	pb.RegisterUpdateAlertingConfigServer(s, &se)

	s.Serve(listener)
}
