package main

import "fmt"

func Add(x, y int, ch chan int) {
	z := x + y
	ch <- z
}

// 通道支持三个主要操作：send，receive 和 close。
// ch <- x // 发送
// x = <-ch // 接收
// <-ch // 接收，丢弃结果

// close(ch) // 关闭
func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

// 类型 chan<- int 是一个只能发送的通道，类型 <-chan int 是一个只能接收的通道。
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	// 使用 make 创建通道：
	// make 函数接受两个参数，第二个参数是可选参数，表示通道容量。不传或者传 0 表示创建了一个无缓冲通道。
	// 无缓冲通道上的发送操作将会阻塞，直到另一个 goroutine 在对应的通道上执行接收操作
	// 所以，无缓冲通道是一种同步通道
	n := make(chan int)
	s := make(chan int)

	go counter(n)
	go squarer(s, n)
	printer(s)
	// 0 1 4 9 16 25 36 49 64 81

	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go Add(i, i, ch)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
		//随机的： 0 8 2 4 6 12 10 14 16
	}
}
