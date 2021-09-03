package example

import (
	"fmt"
	"sync"
	"time"
)

func workerss(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func WaitGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go workerss(i, &wg)
	}

	wg.Wait()
}
