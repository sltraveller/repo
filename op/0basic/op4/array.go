package main
import "fmt"
/*
// 数组
func main() {
	var balance = [5]float32{1,2,3,4}
	sala := balance[3]
	print(sala)

	var a = [...]float32{1,2,3}
	print(a[:2])
	
	//var ptr *int
	//print(ptr)
}


func main() {
	// var number []int
	var number = make([]int,3,5)
	printslice(number)
}
func printslice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}*/

func main() {
	/*
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printslice(numbers)
	fmt.Printf("numbers[1:4]==%v\n", numbers[1:4])
	fmt.Printf("numbers[:3]==%v\n", numbers[:3])
	fmt.Printf("numbers[3:]==%v\n", numbers[3:])
	
	numbers1 := make([]int,0,5)
	printslice(numbers1)
	numbers2 := numbers[:2]
	printslice(numbers2)
	numbers3 := numbers[2:5]
	printslice(numbers3)
	*/
	var number []int
	printslice(number)
	number = append(number, 0)
	printslice(number)
	number = append(number, 1,2,3)
	printslice(number)

	numbers := make([]int, len(number), (cap(number))*2)

	copy(numbers, number)
	printslice(numbers)
}
func printslice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}