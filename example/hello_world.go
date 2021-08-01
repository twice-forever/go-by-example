package example

import "fmt"

// 通过fmt包中的方法Println打印hello world
// 其中fmt包实现了格式化I/O函数，主要由Printing和Scanning组成
// 通过调用Println函数，打印hello world到输出，并且会在操作数之间插入空格，并在最后加入一个换行符
func HelloWorld() {
	fmt.Println("hello world")
}
