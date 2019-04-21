package main

import "fmt"

// 结构体
type composer struct {
	name      string
	birthYear int
}

func main() {
	// composer类型值
	antonio := composer{"an", 1999}
	// 指向 composer 的指针
	agents := new(composer)
	agents.name, agents.birthYear = "Agnes Zim", 199
	// 指向 composer 的指针
	julia := &composer{}
	julia.name, julia.birthYear = "Jilia", 190
	// 指向 composer 的指针
	augusta := &composer{"Augusta", 1847}
	fmt.Println(antonio)
	fmt.Println(agents, julia, augusta)
	println(julia)
}
