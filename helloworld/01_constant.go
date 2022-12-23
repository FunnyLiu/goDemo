package main

import "fmt"

func main() {
	// 无类型整型常量
	const n = 500000000

	// 用编译阶段即可计算出值的表达式来赋值
	const d = 3e20 / n
	fmt.Println(d)
	// 无类型浮点常量
	const zero = 0.0

	// 无类型整型和字符串常量
	const a, b, c = 3, 4, "foo"
	fmt.Println(a, b, c)

	// 多个常量
	const (
		size int64 = 1024
		eof        = -1 // 无类型整型常量
	)
	fmt.Println(size, eof)

	// 常量声明还有可以使用常量生成器 iota，它不会显示写出常量的值，而是从 0 开始，逐项加 1。
	// iota 在每个 const 开头被重置为 0。
	// 从 0 值开始，逐项加 1
	const (
		a0 = iota // 0
		a1 = iota // 1
		a2 = iota // 2
	)
	fmt.Println(a0, a1, a2)

	// 简写，表达式相同，可以省略后面的
	const (
		b0 = iota // 0
		b1        // 1
		b2        // 2
	)
	fmt.Println(b0, b1, b2)

	const (
		bb         = iota      // 0
		cc float32 = iota * 10 // 10
		dd         = iota      // 2
	)
	fmt.Println(bb, cc, dd)
}
