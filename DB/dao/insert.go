package dao

import (
	"DB/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Use struct {
	//Id         uint   `gorm:"AUTO_INCREMENT"`
	gorm.Model
	NodeId       int     `gorm:"size:100"`
	AverCpuUsed  float64 `gorm:"size:100"`
	MaxCpuUsed   float64 `gorm:"size:100"`
	MinCpuUsed   float64 `gorm:"size:100"`
	AverMemUsed  float64 `gorm:"size:100"`
	MaxMemUsed   float64 `gorm:"size:100"`
	MinMemUsed   float64 `gorm:"size:100"`
	AverDiskUsed float64 `gorm:"size:100"`
	MaxDiskUsed  float64 `gorm:"size:100"`
	MinDiskUsed  float64 `gorm:"size:100"`
	//InsertTime *time.Time
}

var db *gorm.DB

func Insert(a model.Aggre) {
	db, err := gorm.Open("mysql", "root:wjh866832@/plm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Errorf("创建数据库连接失败:%v", err)
	}
	defer db.Close()
	db.AutoMigrate(&Use{})
	//// 插入记录
	db.Create(&Use{NodeId: a.Id, MaxCpuUsed: a.MaxValueCPU, MinCpuUsed: a.MinValueCPU, AverCpuUsed: a.AverageValueCPU, MaxMemUsed: a.MaxValueMem, MinMemUsed: a.MinValueMem, AverMemUsed: a.AverageValueMem, MaxDiskUsed: a.MaxValueDisk, MinDiskUsed: a.MinValueDisk, AverDiskUsed: a.AverageValueDisk})
}
