package example

import "fmt"

func Slices() {
	// 初始化切片
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 设置值和取值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// 切片长度
	fmt.Println("len:", len(s))

	// append方法增加新值
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copy操作
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 切片操作
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 切片操作遵循左闭右包原则
	l = s[:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("sl3:", l)

	// 声明切片
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 多维切片，且内部切片长度可以不同
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
