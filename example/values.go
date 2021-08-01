package example

import (
	"fmt"
)

// 打印字符串，浮点数计算，布尔值操作
func Values() {
	fmt.Println("go" + "lang")

	fmt.Println("1.0 + 1.0 =", 1.0+1.0)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
