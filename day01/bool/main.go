package main

import "fmt"

/**
1.布尔值 默认false
2.不允许将整数强制转换为布尔值
3.布尔值不能参与数值运算 也无法与其他类型转换
*/
func main() {
	b1 := true
	var b2 bool
	fmt.Printf("%T,%T\n", b1, b2)
	fmt.Println(b1, b2)
}
