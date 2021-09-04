package example

import (
	"fmt"
	"time"
)

func RateLimiting() {
	// 基本速率限制
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 允许3个爆发事件
	burstyLimiter := make(chan time.Time, 3)

	// 将当前时间送入通道中
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 每隔200h毫秒放入新值
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// 建立5个请求
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// 执行请求
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
