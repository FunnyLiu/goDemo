package main

import (
	"fmt"
	"time"
)

func main() {
	// 并不阻塞
	go spinner(100 * time.Millisecond)
	const n = 45
	// 计算斐波那契数列
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN) // Fibonacci(45) = 1134903170
}

func spinner(delay time.Duration) {
	for {
		// 转小菊花
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
