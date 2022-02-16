package Http

import (
	"DB/dao"
	"encoding/json"
	"fmt"
	"net/http"
)

type ListResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []int  `json:"data"`
}

func GetList(w http.ResponseWriter, r *http.Request) {
	data := dao.FindNodeList()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	msg, _ := json.Marshal(ListResult{Code: 200, Msg: "成功", Data: data})

	fmt.Fprintf(w, string(msg))
}
