package main

import (
	"fmt"
)

func main() {
	G()
}

func G() {
	defer func() {
		fmt.Println("c") //4.  c
	}()
	F()
	fmt.Println("继续执行") // 3.  继续执行
}

func F() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err) // 1.  捕获异常：a
		}
		fmt.Println("b") // 2.  b
	}()
	panic("a")
}
