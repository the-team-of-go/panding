package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

var db4 *gorm.DB

func main() {
	//resp, _ := http.Get("http://127.0.0.1:9090/machine/info/now?id=2")
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	//db4, err := gorm.Open("mysql", "root:wjh866832@/plm?charset=utf8&parseTime=True&loc=Local")
	//if err != nil {
	//	fmt.Errorf("创建数据库连接失败:%v", err)
	//}
	//defer db4.Close()
	//var result []dao.Use
	//db4.Limit(2).Order("id desc").Find(&result, "node_id = ?", 3)
	////db4.Limit(2).Where(&result, "node_id = ?", 2)
	////db4.Limit(2).Last()
	////db4.Last(&result, "node_id = ?", 2)
	//db4.Where("timeout>=? AND timeout<=? ", "1644850500", "1644852600").Find(&result, "node_id = ?", 3)
	//fmt.Println(result)
	//var temp float64 = 23.2310
	//fmt.Println(dao.Decimal(temp))
	var time = time.Now().UnixNano() / 1e6
	fmt.Println(time)
}
