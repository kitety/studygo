package main

import (
	"fmt"
	"math"
)

// 浮点数
func main() {
	// 32位浮点数
	// math.MaxFloat32
	println(math.MaxFloat32)

	f1 := 1.23456789
	// float 64 默认go语言中的小数为64的
	fmt.Printf("%T\n", f1)
	f2 := float32(f1)
	fmt.Printf("%T\n", f2)

	// 不能直接赋值，因为类型不一致
	// f2 = f1

}
