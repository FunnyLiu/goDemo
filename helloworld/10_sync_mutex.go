package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 互斥锁
	// 当一个 goroutine 获取了 Mutex 后，其他 goroutine 不管读写，只能等待，直到锁被释放
	var mutex sync.Mutex
	wg := sync.WaitGroup{}

	// 主 goroutine 先获取锁
	fmt.Println("Locking  (G0)") // 1. Locking  (G0)
	mutex.Lock()
	fmt.Println("locked (G0)") // 2. locked (G0)

	wg.Add(3)
	for i := 1; i < 4; i++ {
		go func(i int) {
			// 由于主 goroutine 先获取锁，程序开始 5 秒会阻塞在这里
			fmt.Printf("Locking2 (G%d)\n", i)
			// 345三个结果是随机
			// 3. Locking2 (G3)
			// 4. Locking2 (G1)
			// 5. Locking2 (G2)
			mutex.Lock()
			fmt.Printf("locked2 (G%d)\n", i)

			time.Sleep(time.Second * 2)
			mutex.Unlock()
			fmt.Printf("unlocked2 (G%d)\n", i)
			// 8 10 12之间间隔两秒
			// 8. locked2 (G3)
			// 9. unlocked2 (G3)
			// 10. locked2 (G1)
			// 11. unlocked2 (G1)
			// 12. locked2 (G2)
			// 13. unlocked2 (G2)
			wg.Done()
		}(i)
	}

	// 主 goroutine 5 秒后释放锁
	time.Sleep(time.Second * 5)
	fmt.Println("ready unlock (G0)")
	// 5 6 之间等待5秒
	// 6. ready unlock (G0)
	mutex.Unlock()
	fmt.Println("unlocked (G0)")
	// 7. unlocked (G0)

	wg.Wait()
}
