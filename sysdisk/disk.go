package sysdisk

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"strconv"
)

type DiskInfo struct {
	Path  string `json:"path"`
	Total string `json:"total"`
	Used  string `json:"used"`
}

func Getdisk() []DiskInfo {
	var nodelist []DiskInfo

	parts, _ := disk.Partitions(false)
	for _, part := range parts {
		u, err := disk.Usage(part.Mountpoint)
		check(err)
		node := DiskInfo{
			Path:  u.Path,
			Total: handle(u.Total),
			Used:  handle(u.Used),
		}
		nodelist = append(nodelist, node)
	}
	return nodelist
}

func handle(data uint64) string {
	if data/1024 < 1 {
		return fmt.Sprintf("%v B", data)
	}
	if data/1024/1024 < 1 {
		return fmt.Sprintf("%v KB", strconv.FormatUint(data/1024, 10))
	}
	if data/1024/1024/1024 < 1 {
		return fmt.Sprintf("%v MB", strconv.FormatUint(data/1024/1024, 10))
	}

	return fmt.Sprintf("%v GiB", strconv.FormatUint(data/1024/1024/1024, 10))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
