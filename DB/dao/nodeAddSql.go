package dao

import (
	"DB/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

type NodeSql struct {
	Id     uint `gorm:"AUTO_INCREMENT"`
	NodeId int  `gorm:"size:100"`
	//gorm.Model

	CpuUsed  float64 `gorm:"size:100"`
	MemUsed  float64 `gorm:"size:100"`
	DiskUsed float64 `gorm:"size:100"`
	Grade    string  `gorm:"size:50"`
	Timeout  int64   `gorm:"size:100"`
	//InsertTime *time.Time
}

func InsertNode(a model.Staus) {
	db1, err := gorm.Open("mysql", "root:wjh866832@/plm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("创建数据库连接失败:%v", err)
	}
	defer db1.Close()
	db1.AutoMigrate(&NodeSql{})
	//// 插入记录
	db1.Create(&NodeSql{NodeId: a.Id, CpuUsed: a.CpuUsed, MemUsed: a.MenUsed, DiskUsed: a.DiskUsed, Grade: a.Grade, Timeout: a.Timeout})
}
