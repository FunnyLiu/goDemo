package main

import (
	"fmt"
)

type Person struct {
	name string
}

type Point struct {
	x, y int
}

func main() {
	p := Person{name: "zhangsan"}

	// 调用方法
	fmt.Println(p.String()) // 1. person name is zhangsan

	// 值接收者，不会改变原始值
	p.Modify()
	fmt.Println(p.String()) // 2. person name is zhangsan
	// 等价于
	(&p).Modify()
	fmt.Println(p.String()) // 3. person name is zhangsan

	// 指针接收者，会改变原始值
	p.ModifyP()
	fmt.Println(p.String()) // 4. person name is lisi
	// 等价于
	(&p).ModifyP()
	fmt.Println(p.String()) // 5. person name is lisi

	// 方法变量
	p1 := Point{1, 2}
	q1 := Point{3, 4}
	f := p1.Add
	fmt.Println(f(q1)) // 6. {4 6}

	// 方法表达式
	f1 := Point.Add
	fmt.Println(f1(p1, q1)) // 7. {4 6}
}

func (p Person) String() string {
	return "person name is " + p.name
}

// 方法在定义的时候，会在 func 和方法名之间增加一个参数，
// 这个参数就是接收者
// 分为值接收者
func (p Person) Modify() {
	p.name = "lisi"
}

// 和指针接收者
func (p *Person) ModifyP() {
	p.name = "lisi"
}

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}
