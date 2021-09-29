package main

import "fmt"

func main() {
	// 整型
	var i1 = 10
	fmt.Printf("%d\n", i1)
	// 十进制转换为二进制
	fmt.Printf("%b\n", i1)
	// 十进制转换为八进制
	fmt.Printf("%o\n", i1)
	// 十进制转换为十六进制
	fmt.Printf("%x\n", i1)

	i2 := 077
	fmt.Printf("%d\n", i2)

	i3 := 0x1234567
	fmt.Printf("%d\n", i3)

	// 查看变量的类型
	fmt.Printf("%T\n", i3)

	// 声明int8的类型
	// 明确指定类型
	i4 := int8(9)
	fmt.Printf("%d\n", i4)
	fmt.Printf("%T\n", i4)

}
