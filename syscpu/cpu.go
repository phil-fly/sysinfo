package syscpu

import (
	"github.com/shirou/gopsutil/cpu"
	"time"
)

type CpuInfo struct {
	ModelName string    `json:"modelName"`
	Usage     []float64 `json:"usage"`  //使用率
	Cores     int       `json:"cores"`  //CPU核数
	CpuNum    int       `json:"cpuNum"` //物理cpu个数
}

func Getcpu() CpuInfo {
	var cpuinfo CpuInfo
	cores, _ := cpu.Percent(time.Second, false)
	cpuStats, _ := cpu.Info()

	for _, cpuStat := range cpuStats {
		cpuinfo.ModelName = cpuStat.ModelName
	}
	cpuinfo.Usage = cores
	cpuinfo.Cores = len(cpuStats)
	cpuinfo.CpuNum = len(cores)

	return cpuinfo
}
