package main

import (
	"DB/Http"
	"net/http"
)

func main() {
	http.HandleFunc("/machine/list", Http.GetList)
	http.HandleFunc("/machine/info/now", Http.GetNodeNow)
	http.HandleFunc("/machine/info/list", Http.GetInfoList)
	http.ListenAndServe("10.243.50.4:9090", nil)
}
