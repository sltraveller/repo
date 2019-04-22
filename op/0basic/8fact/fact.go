package main
import "fmt"
/*
// 递归函数
func factorial(n uint64)(result uint64) {
	if (n > 0) {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

func main() {
	var i int = 4
	fmt.Println(factorial(uint64(i)))
}
*/

// 斐波那契数列

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}
}