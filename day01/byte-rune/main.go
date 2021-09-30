package main

import "fmt"

func main() {
	// 年 长度为3
	s := "2021年9月30日"
	// 长度
	n := len(s) // 16
	fmt.Printf("长度%v\n", n)

	for i := 0; i < n; i++ {
		// println(s[i])
		// fmt.Printf("%c\n", s[i])

	}
	fmt.Println("--")

	// 字符串遍历
	// for _, c := range s {
	// 	fmt.Printf("字符：%c\n", c)
	// }

	// 字符串修改
	// 转换为另外的变量修改
	s2 := "白萝卜"
	// 把字符串强制转换为一个切片
	//  ['白' '萝' '卜']
	s3 := []rune(s2)
	// 改为字符红
	s3[0] = '红'
	// 把rune切片转换为字符串
	println(string(s3))

	c1 := '红' // 红	int32 rune是他的一个别名
	c2 := "红" // 字符串
	fmt.Printf("c1:%T,c2:%T\n", c1, c2)

	c3 := "H" // 字符串
	c4 := 'h' // int32
	// byte- uint8
	fmt.Printf("c3:%T,c4:%T\n", c3, c4)
	fmt.Printf("数字c4:%d\n", c4)

	// 类型转换
	n1 := 10
	var f float64 = float64(n1)
	fmt.Printf("f:%T\n", f)

}
