package sysmem

import (
	"github.com/shirou/gopsutil/mem"
	"fmt"
	"strconv"
)


type MemInfo struct {
	Total             string  `json:"total"`
	Used              string  `json:"used"`
}

func Getmem()(MemInfo){
	v, _ := mem.VirtualMemory()
	return MemInfo{
		Total:handle(v.Total),
		Used:handle(v.Used),
	}
}

func handle(data uint64) string {
	if data/1024 < 1 {
		return fmt.Sprintf("%v B",data)
	}
	if data/1024/1024 < 1 {
		return fmt.Sprintf("%v KB",strconv.FormatUint(data/1024, 10))
	}
	if data/1024/1024/1024 < 1 {
		return fmt.Sprintf("%v MB",strconv.FormatUint(data/1024/1024, 10))
	}


	return fmt.Sprintf("%v GiB",strconv.FormatUint(data/1024/1024/1024, 10))
}