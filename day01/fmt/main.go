package main

import "fmt"

// 占位符
func main() {
	n := 100
	n1 := "100"
	// %T 类型
	// %v 值
	// %d 十进制
	// %b 二进制
	// %o 八进制
	// %x 十六进制
	// %f 浮点数
	// %s 字符串

	// # 会在字符串的时候加个引号
	fmt.Printf("%#v", n1)
	fmt.Printf("%#v", n)
}
