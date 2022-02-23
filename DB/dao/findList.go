package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func FindNodeList() []int {
	var Nodelist []int
	db2, err := gorm.Open("mysql", "root:wjh866832@/plm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("创建数据库连接失败:%v", err)
	}
	defer db2.Close()
	var result []NodeSql
	var flat = make(map[int]int)
	db2.Select("node_id").Find(&result)
	for _, val := range result {
		flat[val.NodeId] = val.NodeId
	}
	for key := range flat {
		Nodelist = append(Nodelist, key)
	}
	return Nodelist
}
