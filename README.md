# sysinfo
获取系统信息：cpu  磁盘  内容等


### 下载安装
```
go get	"github.com/phil-fly/sysinfo/syscpu"
go get	"github.com/phil-fly/sysinfo/sysmem"
go get	"github.com/phil-fly/sysinfo/sysdisk"
```

### 使用示例

```
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
```

### 运行结果
```
{
    "diskInfo":[
        {
            "path":"/",
            "total":"849 GiB",
            "used":"300 GiB"
        },
        {
            "path":"/boot",
            "total":"1014 MB",
            "used":"186 MB"

        },
        {
            "path":"/home",
            "total":"72 GiB",
            "used":"59 MB"
        }
        ],
    "cpuInfo":{
        "modelName":"Intel(R) Xeon(R) CPU E3-1220 v6 @3.00GHz",
        "usage":[12.72727272672295],
        "cores":4,
        "cpuNum":1
    },
    "memInfo":{
        "total":"15 GiB",
        "used":"3 GiB"
    }
}
```