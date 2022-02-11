package dao

import (
	"DB/model"
	"fmt"
)

func Aggregation(a []model.Staus) int {
	if len(a) == 0 {
		return 0
	}
	var sumCpu float64 = 0
	var sumMen float64 = 0
	var sumDisk float64 = 0
	var maxCpuUsed float64 = 0
	var minCpuUsed float64 = 100
	var maxMemUsed float64 = 0
	var minMemUsed float64 = 100
	var maxDiskUsed float64 = 0
	var minDiskUsed float64 = 100
	var avergCpu float64 = 0
	var avergMem float64 = 0
	var averDisk float64 = 0
	for i := 0; i < len(a); i++ {
		if a[i].CpuUsed > maxCpuUsed {
			maxCpuUsed = a[i].CpuUsed
		}
		if a[i].CpuUsed < minCpuUsed {
			minCpuUsed = a[i].CpuUsed
		}
		if a[i].MenUsed > maxMemUsed {
			maxMemUsed = a[i].MenUsed
		}
		if a[i].MenUsed < minMemUsed {
			minMemUsed = a[i].MenUsed
		}
		if a[i].DiskUsed > maxDiskUsed {
			maxDiskUsed = a[i].DiskUsed
		}
		if a[i].DiskUsed < minDiskUsed {
			minDiskUsed = a[i].DiskUsed
		}
		sumDisk += a[i].DiskUsed
		sumCpu += a[i].CpuUsed
		sumMen += a[i].MenUsed

	}
	averDisk = sumDisk / float64(len(a))
	avergCpu = sumCpu / float64(len(a))
	avergMem = sumMen / float64(len(a))
	result := model.Aggre{a[0].Id, a[0].Timeout, maxCpuUsed, minCpuUsed, avergCpu, maxMemUsed, minMemUsed, avergMem, maxDiskUsed, minDiskUsed, averDisk}
	fmt.Println(result)
	Insert(result)
	delete(model.Node, a[0].Id)

	return 1
}
