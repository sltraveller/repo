package main
import "fmt"
/*
func main(){
	// range用于for循环迭代数组，切片(slice)，通道(channel),集合(map)
	nums := []int{2,3,4}
	sum := 0
	for _,num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i,c := range "ab" {
		//枚举unicode字符串，第一个是字符的索引，第二个是字符本身
		fmt.Println(i,c)
	}
}
*/
func main(){
	myfunc(2,3,5)
}
func myfunc(arg...int){//arg是int的slice
	for _,n := range arg{
		fmt.Println(n)
	}
}