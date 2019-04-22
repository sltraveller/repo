package main

import (
	"fmt"
	"unsafe"
)
/*
var (
	q int
	d bool
)
// 一般用于声明全局变量
func main() {
	fmt.Println("hello")
	var a = "run"
	fmt.Println(a)

	var b int
	fmt.Println(b)

	var d bool
	fmt.Println(d)

	var e string
	fmt.Println(e)

	var f int
	c := 1	// 声明新的变量,只能出现在函数体中
	f = c 
	c = 3
	fmt.Println(f,c)
	fmt.Println(c,q)

}
*/
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
	d   //定义常量组为赋值，使用上一行的值
)
func main() {
	const length int = 5
	const width = 5
	var area int
	area = length * width
	fmt.Println("面积为：", area)
	println(a,b,c,d)

	const (
		a = iota	// 0
		b			// 1
		c = "ha"	// ha
		d			// ha
		e = iota	// 4
		f			// 5
	)
	println(a,b,c,d,e,f) 
}