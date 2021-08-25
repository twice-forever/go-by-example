package example

import "fmt"

func Range() {
	// 遍历切片
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 获取索引
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// 遍历map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 只遍历key值
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 遍历字符串是使用unicode码点
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
