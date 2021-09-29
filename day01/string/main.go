package main

import (
	"fmt"
	"strings"
)

// utf8 支持中文
// 双引号
// 单引号为字符 (单独的)'h' '1'
// 1个'A'=1个字节
// 1个'中'=3个字节

// 转义字符
func main() {
	var s string = "Hello, world"
	println(s)
	fmt.Println("E:\\study\\go\\src\\github.com\\studygo")
	h := `
  1
  2
  3`
	fmt.Println(h)
	s2 := `e:\study\go\src\github.com\studygo\day01\string\main.go`
	fmt.Println(s2)

	// 字符串相关操作
	fmt.Println("s长度：", len(s))

	// 字符串拼接 + or fmt.Sprintf

	name := `账单`
	action := `查看`
	name_action := fmt.Sprintf("%s%s", name, action)
	fmt.Println(name_action)

	// 字符串分割
	// 字符串分割：strings.Split(s, "分割符")
	res := strings.Split(s2, "\\")
	// 返回列表list
	fmt.Println(res)

	// 是否包含 strings.Contains
	fmt.Println(strings.Contains(s2, "study"))
	// 前缀 strings.HasPrefix
	fmt.Println(strings.HasPrefix(s2, "e:"))
	// 后缀 strings.HasSuffix
	fmt.Println(strings.HasSuffix(s2, "go1"))
	// 子串出现的位置 strings.Index
	fmt.Println(strings.Index(s2, "go"))
	// 最后一次出现的位置
	fmt.Println(strings.LastIndex(s2, "go"))
	// 拼接字符串
	// 数组的拼接
	fmt.Println(strings.Join(res, "+"))
}
