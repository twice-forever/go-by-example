package example

import "fmt"

// 变量，通过var声明
// 声明之后没有赋值，会初始化为零值
// 通过:=简写赋值
// go会通过赋值类型，自动判断变量的类型
func Variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}
