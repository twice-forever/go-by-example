package example

import "fmt"

func Arrays() {
	// 声明一个数组
	var a [5]int
	fmt.Println("emp:", a)

	// 使用array[index] = value来指定数组的值，或者用array[index]来取值
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 打印数组的长度
	fmt.Println("len:", len(a))

	// 声明并初始化数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 声明并初始化一个二维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
