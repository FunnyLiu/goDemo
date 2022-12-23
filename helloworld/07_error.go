package main

import (
	"fmt"
)

func main() {
	n, err := echo(10)
	if err != nil {
		fmt.Println("error: " + err.Error())
	} else {
		fmt.Println(n) // 1.  10
	}
}

func echo(param int) (int, error) {
	return param, nil
}
