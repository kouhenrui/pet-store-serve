package global

import (
	"golang.org/x/time/rate"
	"log"
)

func BucketInit() {
	l := rate.NewLimiter(1, 5)
	log.Println(l.Limit(), l.Burst())
}

//CPU 使用量
//func CPUCheck() {
//	cores, _ := cpu.Counts(false)
//
//	cpus, err := cpu.Percent(time.Duration(200)*time.Millisecond, true)
//	if err == nil {
//		for i, c := range cpus {
//			fmt.Printf("cpu%d : %f%%\n", i, c)
//		}
//	}
//	a, _ := load.Avg()
//	l1 := a.Load1
//	l5 := a.Load5
//	l15 := a.Load15
//	fmt.Println(l1)
//	fmt.Println(l5)
//	fmt.Println(l15)
//	fmt.Println(cores)
//}

// 内存使用量
//func RAMCheck() {
//	u, _ := mem.VirtualMemory()
//	usedMB := int(u.Used) / MB
//	totalMB := int(u.Total) / MB
//	usedPercent := int(u.UsedPercent)
//	fmt.Printf("usedMB:%d,totalMB:%d,usedPercent:%d", usedMB, totalMB, usedPercent)
//}
