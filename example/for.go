package example

import "fmt"

func For() {
	// 单条件循环
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 经典的初始/条件/后续for循环
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 不带条件的for循环会一直持续下去，直到在循环体内使用break或return跳出循环
	for {
		fmt.Println("loop")
		break
	}

	// 使用continue进入下一个循环
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
