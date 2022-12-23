package main

import (
	"fmt"
)

func main() {
	// defer 语句的执行是按调用 defer 语句的倒序执行。
	defer func() {
		fmt.Println("first") // 5.  first
	}()

	defer func() {
		fmt.Println("second") // 4.  second
	}()

	fmt.Println("done") // 1. done

	fmt.Println(triple(4)) //3. 12
}

func double(x int) (result int) {
	// defer 语句在 return 语句之后执行
	defer func() {
		fmt.Printf("double(%d) = %d\n", x, result) //2. double(4) = 8
	}()

	return x + x
}

func triple(x int) (result int) {
	defer func() {
		result += x
	}()

	return double(x)
}
