package server

import (
	"DB/Http"
	"fmt"
	"net/http"
)

func HttpServer() {
	http.HandleFunc("/machine/list", Http.GetList)
	http.HandleFunc("/machine/info/now", Http.GetNodeNow)
	http.HandleFunc("/machine/info/list", Http.GetInfoList)
	go func() {
		fmt.Println("监听http端口：10.243.50.4:9090")
	}()
	http.ListenAndServe("10.243.50.4:9090", nil)
}
