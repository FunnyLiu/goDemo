package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// RWMutex 属于经典的单写多读模型，当读锁被占用时，会阻止写，但不阻止读。而写锁会阻止写和读。

	var rwMutex sync.RWMutex
	wg := sync.WaitGroup{}

	Data := 0
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func(t int) {
			// 第一次运行后，写解锁。
			// 循环到第二次时，读锁定后，goroutine 没有阻塞，同时读成功。
			fmt.Println("Locking")
			// 1. 连续10个Locking
			rwMutex.RLock()
			defer rwMutex.RUnlock()
			fmt.Printf("Read data: %v\n", Data)
			wg.Done()
			time.Sleep(2 * time.Second)
		}(i)
		go func(t int) {
			// 写锁定下是需要解锁后才能写的
			rwMutex.Lock()
			defer rwMutex.Unlock()
			Data += t
			fmt.Printf("Write Data: %v %d \n", Data, t)
			wg.Done()
			time.Sleep(2 * time.Second)
		}(i)
	}

	wg.Wait()
}
