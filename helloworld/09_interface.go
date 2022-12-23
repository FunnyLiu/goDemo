package main

import "fmt"

// 定义接口，包含 Eat 方法
type Duck interface {
	Eat()
}

type Duck1 interface {
	Eat()
	Walk()
}

// 定义 Cat 结构体，并实现 Eat 方法
type Cat struct{}

func (c *Cat) Eat() {
	fmt.Println("cat eat")
}

// 定义 Dog 结构体，并实现 Eat 方法
type Dog struct{}

func (d *Dog) Eat() {
	fmt.Println("dog eat")
}

func (d *Dog) Walk() {
	fmt.Println("dog walk")
}

func main() {
	var c Duck = &Cat{}
	c.Eat() // 1. cat eat

	var d Duck = &Dog{}
	d.Eat() // 2. do eat

	s := []Duck{
		&Cat{},
		&Dog{},
	}
	for _, n := range s {
		n.Eat() // 3. cat eat 4. dog eat
	}

	var c1 Duck1 = &Dog{}
	var c2 Duck = c1
	c2.Eat() // 5. dog eat

	// 类型断言
	var n interface{} = 55
	assert(n) // 6. 55
	var n1 interface{} = "hello"
	// assert(n1) // panic: interface conversion: interface {} is string, not int
	assertFlag(n1) // 不是int类型，不会输出
	// 类型断言
	assertInterface(c) // 7.  &{}

	// 类型查询
	searchType(50)         // 8. Int: 50
	searchType("zhangsan") // 9. String: zhangsan
	searchType(c1)         // 10. dog eat
	searchType(50.1)       // 11.  Unknown type

	// 空接口
	s1 := "Hello World"
	i := 50
	strt := struct {
		name string
	}{
		name: "AlwaysBeta",
	}
	test(s1)   // Type = string, value = Hello World
	test(i)    // Type = int, value = 50
	test(strt) // Type = struct { name string }, value = {AlwaysBeta}
}

func assert(i interface{}) {
	// 类型断言是作用在接口值上的操作，x.(T)
	// 如果是，则输出 x 的值；如果不是，程序直接 panic。
	s := i.(int)
	fmt.Println(s)
}

func assertInterface(i interface{}) {
	// 类型断言是作用在接口值上的操作，x.(T)
	// 如果是，则输出 x 的值；如果不是，程序直接 panic。
	s := i.(Duck)
	fmt.Println(s)
}

func assertFlag(i interface{}) {
	if s, ok := i.(int); ok {
		fmt.Println(s)
	}
}

func searchType(i interface{}) {
	// 语法类似类型断言，只需将 T 直接用关键词 type 替代。
	switch v := i.(type) {
	case string:
		fmt.Printf("String: %s\n", i.(string))
	case int:
		fmt.Printf("Int: %d\n", i.(int))
	case Duck:
		v.Eat()
	default:
		fmt.Printf("Unknown type\n")
	}
}

func test(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}
