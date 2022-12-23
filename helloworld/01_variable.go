package main

import "fmt"

// 全局变量
var gg = "global"

func main() {
	// 作用域
	fmt.Println(gg) // 输出 global
	gg = "local"
	fmt.Println(gg) // 输出 local

	// 第一种方式：通过var来声明
	var a = "initial" // 声明单个变量

	var d = true // 声明布尔值

	// 以组的方式声明多个变量
	var (
		b1, c1 int // 默认会赋予0
		b2, c2 = 3, 4
	)
	// 第二种方式：使用短变量声明方式
	// go会自己判断类型
	f := "short"
	// 什么多个值
	g, h := 5, "alwaysbeta"

	fmt.Println(a)
	fmt.Println(d)
	fmt.Println(b1, c1, b2, c2)
	fmt.Println(f, g, h)

	// 赋值 使用=
	var m, n int
	m = 9
	n = 10
	fmt.Println(m, n) // 9, 10
	// 直接交换值
	m, n = n, m
	fmt.Println(m, n) // 10, 9

	// 如果有不需要的变量，使用空标识符 _ 来忽略
	// 空标识符
	r := [5]int{1, 2, 3, 4, 5}
	for _, v := range r {
		// fmt.Println(i, v)
		// fmt.Println(v)   // 定义 i 但不用会报错 i declared but not used
		fmt.Println(v) // 忽略索引
	}
}
