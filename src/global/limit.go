package global

import (
	"golang.org/x/time/rate"
)

var (
	Limiter *rate.Limiter
)

func StartLimit() {
	Limiter = rate.NewLimiter(10, 100) //令牌桶大小为100,并以每秒10个Token的速率向桶中放置Token
	//fmt.Println("令牌存放速率：", Limiter.Limit(), "令牌桶大小：", Limiter.Burst())
	//log.Println("令牌存放速率：", Limiter.Limit(), "令牌桶大小：", Limiter.Burst())
	//fmt.Println(limiter.Limit(), limiter.Burst())
}
func LimitReset() {
	Limiter.Reserve()
}

//	func wait() {
//		limiter = rate.NewLimiter(10, 100) //令牌桶大小为100,并以每秒10个Token的速率向桶中放置Token
//		fmt.Println(limiter.Limit(), limiter.Burst())
//
//		c, _ := context.WithCancel(context.TODO())
//		for {
//			limiter.Wait(c)
//			time.Sleep(200 * time.Millisecond)
//			fmt.Println(time.Now().Format("2016-01-02 15:04:05.000"))
//		}
//	}
//func Allow() {
//	//limit := rate.Every(100 * time.Millisecond)
//	//limiter := rate.NewLimiter(limit, 10)
//
//	for {
//		if limiter.AllowN(time.Now(), 2) {
//			fmt.Println(time.Now().Format("2016-01-02 15:04:05.000"))
//		} else {
//			time.Sleep(6 * time.Second)
//		}
//	}
//}

//
//func release() {
//	limiter := rate.NewLimiter(10, 100)
//	for {
//		r := limiter.ReserveN(time.Now(), 20)
//		time.Sleep(r.Delay())
//		fmt.Println(time.Now().Format("2016-01-02 15:04:05.000"))
//	}
//}
