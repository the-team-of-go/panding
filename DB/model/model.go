package model

import (
	"sync"
)

type AlertingConfig struct {
	Id               int64  `json:"id"`
	Name             string `json:"name"`
	Timeout          int64  `json:"timeout"`
	MaxValueCPU      int64  `json:"maxvaluecpu"`
	MinValueCPU      int64  `json:"minvaluecpu"`
	AverageValueCPU  int64  `json:"averagevaluecpu"`
	MaxValueMem      int64  `json:"maxvaluemem"`
	MinValueMem      int64  `json:"minvaluemem"`
	AverageValueMem  int64  `json:"averagevaluemem"`
	MaxValueDisk     int64  `json:"maxvalueDisk"`
	MinValueDisk     int64  `json:"minvalueDisk"`
	AverageValueDisk int64  `json:"averagevalueDisk"`
}

type Aggre struct {
	Id               int     `json:"id"`
	Timeout          int64   `json:"timeout"`
	MaxValueCPU      float64 `json:"maxvaluecpu"`
	MinValueCPU      float64 `json:"minvaluecpu"`
	AverageValueCPU  float64 `json:"averagevaluecpu"`
	MaxValueMem      float64 `json:"maxvaluemem"`
	MinValueMem      float64 `json:"minvaluemem"`
	AverageValueMem  float64 `json:"averagevaluemem"`
	MaxValueDisk     float64 `json:"maxvalueDisk"`
	MinValueDisk     float64 `json:"minvalueDisk"`
	AverageValueDisk float64 `json:"averagevalueDisk"`
}

type Staus struct {
	Id       int
	CpuUsed  float64
	MenUsed  float64
	DiskUsed float64
	Grade    string
	Timeout  int64
}

var Node = make(map[int][]Staus)

var SqlAlteringConfig = AlertingConfig{000, "sql", 0, 10, 20, 50, 90, 10, 50, 80, 20, 50}

var Admin []string
var Email []string = []string{"as.59@qq.com", "569105057@qq.com", "linf_z@163.com", "1054932842@qq.com"}
var Id []int64
var one = Staus{1, 65, 23, 49, "danger", 0}
var two = Staus{2, 59, 29, 19, "servious", 0}

var Wg sync.WaitGroup

var Sending bool = true
