package example

import (
	"fmt"
	"time"
)

func workers(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func WorkPool() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	result := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go workers(w, jobs, result)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-result
	}
}
