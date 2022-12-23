package main

import (
	"fmt"
)

func main() {
	// 跳出循环
	for i := 0; ; i++ {
		if i == 2 {
			goto L1
		}
		fmt.Println(i)
	}
L1:
	fmt.Println("Done")

	// 跳过变量声明，不允许
	// 	goto L2
	// 	j := 1
	// L2:

	// break 跳转到标签处，然后跳过 for 循环
L3:
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if i >= 2 {
				break L3
			}
			if j > 4 {
				break
			}
			fmt.Println(i, j)
		}
	}

	// continue 跳转到标签处，然后执行 i++
L4:
	for i := 0; ; i++ {
		for j := 0; j < 6; j++ {
			if i > 4 {
				break L4
			}
			if i >= 2 {
				continue L4
			}
			if j > 4 {
				continue
			}
			fmt.Println(i, j)
		}
	}
}

// 0
// 1
// Done
// 0 0
// 0 1
// 0 2
// 0 3
// 0 4
// 1 0
// 1 1
// 1 2
// 1 3
// 1 4
// 0 0
// 0 1
// 0 2
// 0 3
// 0 4
// 1 0
// 1 1
// 1 2
// 1 3
// 1 4
