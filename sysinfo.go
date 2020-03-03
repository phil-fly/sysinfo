package main

import (
	"encoding/json"
	"log"
	"sysinfo/syscpu"
	"sysinfo/sysdisk"
	"sysinfo/sysmem"
)


type sysinfo struct {
	DiskInfo []sysdisk.DiskInfo `json:"diskInfo"`
	CpuInfo	 syscpu.CpuInfo		`json:"cpuInfo"`
	MemInfo	 sysmem.MemInfo		`json:"memInfo"`
}

func main(){
	var info sysinfo
	info.DiskInfo = sysdisk.Getdisk()
	info.CpuInfo = syscpu.Getcpu()
	info.MemInfo = sysmem.Getmem()
	bytesData, _ := json.Marshal(info)
	log.Printf("%v",string(bytesData))
}

