package model

type AlertingConfig struct {
	Name             string `json:"name"`
	Timeout          int64  `json:"timeout"`
	MaxValueCPU      int    `json:"maxvaluecpu"`
	MinValueCPU      int    `json:"minvaluecpu"`
	AverageValueCPU  int    `json:"averagevaluecpu"`
	MaxValueMem      int    `json:"maxvaluemem"`
	MinValueMem      int    `json:"minvaluemem"`
	AverageValueMem  int    `json:"averagevaluemem"`
	MaxValueDisk     int    `json:"maxvalueDisk"`
	MinValueDisk     int    `json:"minvalueDisk"`
	AverageValueDisk int    `json:"averagevalueDisk"`
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

var SqlAlteringConfig = AlertingConfig{"sql", 0, 80, 20, 50, 90, 10, 50, 80, 20, 50}

var Admin []string
var Email []string
var one = Staus{1, 65, 23, 49, "danger", 0}
var two = Staus{2, 59, 29, 19, "servious", 0}

//	1: one,
//	2: two,
//	3: one,
//}
