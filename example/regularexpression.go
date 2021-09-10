package example

import (
	"bytes"
	"fmt"
	"regexp"
)

func RegularExpression() {
	// 是否匹配字符串
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	// 匹配字符串
	fmt.Println(r.MatchString("peach"))

	// 查找匹配的字符串
	fmt.Println(r.FindString("peach punch"))

	// 查找匹配的字符串，并返回开始索引和结束索引
	fmt.Println(r.FindStringIndex("peach punch"))

	// 返回完全匹配和局部匹配的字符串
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// 查找所有匹配项
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// 限制查找次数
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// 直接提供[]byte作为参数
	fmt.Println(r.Match([]byte("peach")))

	// MustCompile用panic替代返回错误
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// 可以用来替换其他值
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// 允许使用函数来转换匹配文本
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
