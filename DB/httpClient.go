package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://10.243.50.4:9090/machine/list")
	if err != nil {
		fmt.Printf("get failed ,err:%v\n", err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read fail err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
