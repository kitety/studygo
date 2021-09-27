package main

import "fmt"

// 小驼峰命名法
// 声明一个变量
var name string
var age int
var isOk bool

// 批量声明
var (
	name1 string //''
	age1  int    // 0
	isOk1 bool   // false
)

func main() {
	name = "张三"
	age = 20
	isOk = true
	// go语言变量声明了必须使用
	// %s  字符串

	fmt.Printf("name=%v,age=%v,isOk=%v\n", name, age, isOk)
	// println 加入换行符

	// 声明并赋值
	var s1 string = "good"
	fmt.Println(s1)
	// 类型推导（根据值推到类型）
	var s2 = "good"
	fmt.Println(s2)

	// 简短变量声明 只在函数里面使用
	s3 := "good"
	fmt.Println(s3)

	// 匿名变量 _ 占位
	// 同一个作用域（{}）不能重复声明同名变量



}
