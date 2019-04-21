package main

import "fmt"

/*
	数组 和 切片
*/
func main() {
	//arr()
	slice()
}

/*
	数组
	go的数组是一个定长的序列
	使用操作符 [] 来索引
    数组是定长
*/
func arr() {
	var buffer [20]byte
	var grid1 [3][3]int

	grid1[1][0], grid1[1][1], grid1[1][2] = 8, 6, 2
	grid2 := [3][3]int{{4, 3}, {8, 6, 2}}
	cities := [...]string{"Shanghai", "Mumbai", "Istanbul", "Beijing"}
	cities[len(cities)-1] = "Karachi"

	fmt.Println(buffer)
	fmt.Println(grid1)
	fmt.Println(grid2)
	fmt.Println(cities)
}

/*
	切片
	内置函数 make() 用于创建切片、映射和通道
	[]Type{} 和 make([]Type,0) 等价，两者都创建一个空切片，我们可以使用内置的append()函数来有效的增加切片的容量
	
*/
func slice() {
	s := []string{"A", "B", "C", "D", "E", "F", "G"}
	t := s[:5]           // [A B C D E]
	u := s[3 : len(s)-1] // [D E F]
	fmt.Println(s, t, u)

	u[1] = "X"
	// 因为 s t  u 都是同一个底层数组的引用，其中一个改变会影响到其他所有指向该相同数组的任何其他引用
	fmt.Println(s, t, u)
}
