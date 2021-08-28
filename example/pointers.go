package example

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func Pointers() {
	i := 1
	fmt.Println("initial:", i)

	// 值传递，不会改变传值变量的值
	zeroval(i)
	fmt.Println("zeroval:", i)

	// 指针传递，能改变传值变量的值
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 打印内存地址
	fmt.Println("pointer:", &i)
}
