package example

import "fmt"

func ClosingChannel() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 通过接收的第二个值传递判断jobs是否已经关闭
	// 再通过done通道通知main协程任务完成
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
