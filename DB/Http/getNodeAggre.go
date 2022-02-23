package Http

import (
	"DB/dao"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type Groud struct {
	Code int         `json:"code"`
	Meg  string      `json:"msg"`
	Data []NodeAggre `json:"data"`
}

type Basic struct {
	Code int         `json:"code"`
	Meg  string      `json:"msg"`
	Data []BasicNode `json:"data"`
}

type NodeAggre struct {
	Id      int   `json:"nodeid"`
	Cpu     Zhi   `json:"cpu"`
	Mem     Zhi   `json:"mem"`
	Disk    Zhi   `json:"disk"`
	Timeout int64 `json:"timeout"`
}

type Zhi struct {
	Max   float64
	Min   float64
	Averg float64
}

type BasicNode struct {
	NodeId  int     `json:"nodeId"`
	Cpu     float64 `json:"cpu"`
	Mem     float64 `json:"mem"`
	Disk    float64 `json:"disk"`
	Timeout int64   `json:"timeout"`
}

func GetInfoList(w http.ResponseWriter, r *http.Request) {
	db4, err := gorm.Open("mysql", "root:wjh866832@/plm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("创建数据库连接失败:%v", err)
	}
	defer db4.Close()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var result []dao.Use
	queryParam := r.URL.Query()
	id := queryParam.Get("id")
	searchId, _ := strconv.Atoi(id)
	lei := queryParam.Get("type")
	start := queryParam.Get("start")
	end := queryParam.Get("end")
	searchBegin, _ := strconv.Atoi(start)
	searchEnd, _ := strconv.Atoi(end)
	if start != "" {
		if end != "" {
			if lei == "group" { //有时间戳查询的聚合数据
				db4.Where("timeout>=? AND timeout<=? ", searchBegin, searchEnd).Find(&result, "node_id = ?", searchId)
				var fanhui []NodeAggre
				for _, v := range result {
					temp := NodeAggre{v.NodeId, Zhi{v.MaxCpuUsed, v.MinCpuUsed, v.AverCpuUsed}, Zhi{v.MaxMemUsed, v.MinMemUsed, v.AverMemUsed}, Zhi{v.MaxDiskUsed, v.MinDiskUsed, v.AverDiskUsed}, v.Timeout}
					fanhui = append(fanhui, temp)
				}
				msg, _ := json.Marshal(Groud{Code: 200, Meg: "成功", Data: fanhui})
				fmt.Fprintf(w, string(msg))
			} else { //有时间戳查询的不聚合的原生数据
				var danDianResult []dao.NodeSql
				db4.Where("timeout>=? AND timeout<=? ", searchBegin, searchEnd).Find(&danDianResult, "node_id = ?", searchId)
				var fanhui []BasicNode
				for _, v := range danDianResult {
					temp := BasicNode{v.NodeId, v.CpuUsed, v.MemUsed, v.DiskUsed, v.Timeout}
					fanhui = append(fanhui, temp)
				}
				msg, _ := json.Marshal(Basic{Code: 200, Meg: "成功", Data: fanhui})
				fmt.Fprintf(w, string(msg))
			}
		} else { //有开始时间戳没结束的聚合数据
			if lei == "group" {
				db4.Where("timeout>=? ", searchBegin).Find(&result, "node_id = ?", searchId)
				var fanhui []NodeAggre
				for _, v := range result {
					temp := NodeAggre{v.NodeId, Zhi{v.MaxCpuUsed, v.MinCpuUsed, v.AverCpuUsed}, Zhi{v.MaxMemUsed, v.MinMemUsed, v.AverMemUsed}, Zhi{v.MaxDiskUsed, v.MinDiskUsed, v.AverDiskUsed}, v.Timeout}
					fanhui = append(fanhui, temp)
				}
				msg, _ := json.Marshal(Groud{Code: 200, Meg: "成功", Data: fanhui})
				fmt.Fprintf(w, string(msg))
			} else { //有开始时间戳没有结束的原生数据
				var danDianResult []dao.NodeSql
				db4.Where("timeout>=?", searchBegin).Find(&danDianResult, "node_id = ?", searchId)
				var fanhui []BasicNode
				for _, v := range danDianResult {
					temp := BasicNode{v.NodeId, v.CpuUsed, v.MemUsed, v.DiskUsed, v.Timeout}
					fanhui = append(fanhui, temp)
				}
				msg, _ := json.Marshal(Basic{Code: 200, Meg: "成功", Data: fanhui})
				fmt.Fprintf(w, string(msg))
			}
		}
	} else { //时间区间不传时聚合数据
		if lei == "group" {
			db4.Limit(100).Order("id desc").Find(&result, "node_id = ?", searchId)
			var fanhui []NodeAggre
			for _, v := range result {
				temp := NodeAggre{v.NodeId, Zhi{v.MaxCpuUsed, v.MinCpuUsed, v.AverCpuUsed}, Zhi{v.MaxMemUsed, v.MinMemUsed, v.AverMemUsed}, Zhi{v.MaxDiskUsed, v.MinDiskUsed, v.AverDiskUsed}, v.Timeout}
				fanhui = append(fanhui, temp)
			}
			msg, _ := json.Marshal(Groud{Code: 200, Meg: "成功", Data: fanhui})
			fmt.Fprintf(w, string(msg))
		} else { //时间区间不传时原生数据
			var danDianResult []dao.NodeSql
			db4.Limit(100).Order("id desc").Find(&danDianResult, "node_id = ?", searchId)
			var fanhui []BasicNode
			for _, v := range danDianResult {
				temp := BasicNode{v.NodeId, v.CpuUsed, v.MemUsed, v.DiskUsed, v.Timeout}
				fanhui = append(fanhui, temp)
			}
			msg, _ := json.Marshal(Basic{Code: 200, Meg: "成功", Data: fanhui})
			fmt.Fprintf(w, string(msg))
		}
	}

}
