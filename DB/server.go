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

func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	id := req.Id
	cpuUsed := req.Cpuused
	memUsed := req.Memused
	diskUsed := req.DiskUsed
	timeout := req.Timeout
	temp := model.Staus{int(id), float64(cpuUsed), float64(memUsed), float64(diskUsed), int(timeout)}
	dao.AddNums(temp)
	//fmt.Println(temp)
	resp = &pb.UserResponse{Code: 1, Mesg: "ok"}
	return
}

type SqlInfoService struct{}

var sql = SqlInfoService{}

func (s *SqlInfoService) GetSqlInfo(ctx context.Context, req *pb.SqlRequest) (resp *pb.SqlResponse, err error) {
	fmt.Println("配置更改")
	model.SqlAlteringConfig.Name = req.Name
	model.SqlAlteringConfig.MaxValueCPU = int(req.MaxValueCpu)
	model.SqlAlteringConfig.MinValueCPU = int(req.MinValueCpu)
	model.SqlAlteringConfig.AverageValueCPU = int(req.AvergValueCpue)
	model.SqlAlteringConfig.MaxValueMem = int(req.MaxValueMem)
	model.SqlAlteringConfig.MinValueMem = int(req.MinValueMem)
	model.SqlAlteringConfig.AverageValueMem = int(req.AvergValueMem)
	model.SqlAlteringConfig.MaxValueDisk = int(req.MaxValueDisk)
	model.SqlAlteringConfig.Timeout = int(req.Timeout)
	fmt.Println(model.SqlAlteringConfig)
	resp = &pb.SqlResponse{Code: 1, Mesg: "ok"}
	return
}

type AdminInfoService struct{}

var ad = AdminInfoService{}

func (a *AdminInfoService) GetAdminInfo(ctx context.Context, req *pb.AdminRequest) (resp *pb.AdminResponse, err error) {
	model.Admin = []string{}
	model.Email = []string{}
	for i := 0; i < len(req.Name); i++ {
		model.Admin = append(model.Admin, req.Name[i])
	}
	for i := 0; i < len(req.Email); i++ {
		model.Email = append(model.Email, req.Email[i])
	}

	fmt.Println(model.Admin)
	fmt.Println(model.Email)
	resp = &pb.AdminResponse{Code: 1, Mesg: "ok"}
	return
}

func dayin() {
	for _, v := range model.Node {
		go dao.Aggregation(v)
	}

}

func main() {
	c := cron.New()                   // 新建一个定时任务对象
	c.AddFunc("0 */2 * * * ?", dayin) // 给对象增加定时任务
	c.Start()

	addr := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常%s\n", err)
	}
	fmt.Printf("监听端口%s\n", addr)
	s := grpc.NewServer()
	pb.RegisterUserInfoServiceServer(s, &u)
	pb.RegisterSqlDefaultServiceServer(s, &sql)
	pb.RegisterAdminGetServiceServer(s, &ad)
	s.Serve(listener)
}
