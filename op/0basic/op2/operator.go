package main

/*
func main() {
	var a = 21
	var b = 10
	var c int
	c = a + b
	println(c)
	a--
	println(a)
}

func main() {
	var a = 10
	if a < 10 {
		print(a)
	} else {
		a++
		print(a)
	}
}

func main() {
	var grade string = "B"
	var marks int = 70
	switch marks {
		case 90: grade = "A"
		case 80,70: grade = "B"
		case 60: grade = "C"
		default: grade = "D"
	}
	switch {
		case grade == "A": print("1")
		case grade == "B": print("2")
		default: print("3")
	}
}
import "fmt"
func main() {
	var b int = 3
	var a int 
	number := [3]int{1,2}
	for a := 0 ; a < b ; a++ {
		fmt.Println(a)
	}
	for a < b {
		a ++
		fmt.Println(a)
	}
	for i,x := range number {
		fmt.Println(i,x)
	}
}

import "fmt"
func main() {
	var i, j int
	for i = 1; i < 10; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf(" %d*%d=%d", j, i, i*j)
		}
		print("\n")
	}
}
*/

import "math"
func pow(a,b,c float64) float64 {
	if v := math.Pow(a,b); v < c{
		return v
	}
	return c
}
func main(){
	print(
		int(pow(2,5,1)),
		int(pow(2,1,5)),
	)
}