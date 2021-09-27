package main

import "fmt"

// 批量声明
const (
	statusOk = 200
	notFound = 404
)

// 批量声明，某一行没有写等号，默认和上一行一致
const (
	n1 = 100
	n2
)

// iota 常量计数器
const (
	a1 = iota //0
	a2        //1
	a3        //2
)

const (
	b1 = iota // 0
	b2        //1
	_         //2
	b3        //3
)

// 插队
const (
	c1 = iota //0
	c2 = 100
	c3 = iota //2 (c2新增一行 计数器加一)
	c4
)

// 多个常量声明在一行
// 新加一行常量声明 iota 加1
const (
	d1, d2 = iota + 1, iota + 2 // 1 2
	d3, d4 = iota + 1, iota + 2 // 2 3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	// 常量
	const PI = 3.14
	// 常量不能修改
	fmt.Println(PI)
	fmt.Println("n1", n1)
	fmt.Println("n2", n2)

	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
	fmt.Println("a3", a3)

	fmt.Println("b1", b1)
	fmt.Println("b2", b2)
	fmt.Println("b3", b3)

	fmt.Println("c1", c1)
	fmt.Println("c2", c2)
	fmt.Println("c3", c3)
	fmt.Println("c4", c4)

	fmt.Println("d1", d1)
	fmt.Println("d2", d2)
	fmt.Println("d3", d3)
	fmt.Println("d4", d4)
}
