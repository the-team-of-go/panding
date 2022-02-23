package Http

import (
	"DB/dao"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type NodeInfoResult struct {
	Code int      `json:"code"`
	Meg  string   `json:"msg"`
	Info NodeInfo `json:"data"`
}

type NodeInfo struct {
	Id       int
	CpuUsed  float64
	MemUsed  float64
	DiskUsed float64
	Status   string
	Timeout  int64
}

func GetNodeNow(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	id := queryParam.Get("id")
	searchId, _ := strconv.Atoi(id)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	db3, err := gorm.Open("mysql", "root:wjh866832@/plm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("创建数据库连接失败:%v", err)
	}
	defer db3.Close()
	temp := dao.FindNodeList()
	flag := make(map[int]int)
	for _, v := range temp {
		flag[v] = v
	}
	_, ok := flag[searchId]
	if ok {
		var result dao.NodeSql
		db3.Last(&result, "node_id = ?", searchId)
		nodeinfo := NodeInfo{result.NodeId, result.CpuUsed, result.MemUsed, result.DiskUsed, result.Grade, result.Timeout}
		msg, _ := json.Marshal(NodeInfoResult{Code: 200, Meg: "成功", Info: nodeinfo})
		fmt.Fprintf(w, string(msg))
	} else {
		msg, _ := json.Marshal(NodeInfoResult{Code: 404, Meg: "不存在此结点", Info: NodeInfo{}})
		fmt.Fprintf(w, string(msg))
	}
}
