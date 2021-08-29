package example

import "fmt"

type rect struct {
	width, height int
}

// *rect类型的接收器
func (r *rect) area() int {
	return r.width * r.width
}

// 值类型的接受器
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func Methods() {
	r := rect{width: 10, height: 5}

	// 调用结构体方法
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// 调用方法时，Go会自动处理值和指针之间的转换
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
