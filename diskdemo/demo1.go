package main

import (
	"fmt"
	"syscall"
	"time"
)

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}
func main() {
	//mounts, _ := gofstab.ParseSystem()
	//
	//for _, val := range mounts {
	//	//fmt.Printf("%v\n",val.File)
	//	if val.File == "swap" || val.File == "/dev/shm" || val.File == "/dev/pts" || val.File == "/proc" || val.File == "/sys" {
	//		continue
	//	}
	//	disk := DiskUsage(val.File)
	//	//fmt.Printf("All: %.2f GB\n", float64(disk.All)/float64(GB))
	//	//fmt.Printf("Used: %.2f GB\n", float64(disk.Used)/float64(GB))
	//	//fmt.Printf("Free: %.2f GB\n", float64(disk.Free)/float64(GB))
	//	diskall := float64(disk.All) / float64(GB)
	//	diskfree := float64(disk.Free) / float64(GB)
	//
	//	dfpercent := float64(diskfree / diskall)
	//	fmt.Printf("%s %.2f%%\n", val.File, dfpercent*100)
	//}

	disk := DiskUsage("/")

	v, err := time.Parse(time.RFC3339, "2018-01-22T15:24:15Z")
	fmt.Println(v, err)

	fmt.Printf("All: %.2f GB\n", float64(disk.All)/float64(GB))
	fmt.Printf("Used: %.2f GB\n", float64(disk.Used)/float64(GB))
	fmt.Printf("Free: %.2f GB\n", float64(disk.Free)/float64(GB))

}
