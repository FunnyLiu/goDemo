package main

import (
	"fmt"
)

func main() {
	i := 1
	// 只有条件
	for i <= 3 {
		fmt.Println(i) // 1 2 3
		i = i + 1
	}

	// 有变量初始化和条件
	for j := 7; j <= 9; j++ {
		fmt.Println(j) // 7 8 9
	}

	// 死循环
	for {
		fmt.Println("loop")
		break
	}

	// 遍历数组
	a := [...]int{10, 20, 30, 40}
	for i := range a {
		fmt.Println(i) // 0 1 2 3
	}
	for i, v := range a {
		fmt.Println(i, v) // 0 10, 1 20, 2 30, 3 40
	}

	// 遍历切片
	s := []string{"a", "b", "c"}
	for i := range s {
		fmt.Println(i) // 0 1 2
	}
	for i, v := range s {
		fmt.Println(i, v) // 0 a, 1 b, 2 c
	}

	// 遍历字典
	m := map[string]int{"a": 10, "b": 20, "c": 30}
	for k := range m {
		fmt.Println(k) // a b c
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
