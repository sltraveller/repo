package main

import (
	"fmt"
	"reflect"
)

/*
a. reflect.TypeOf()，获取变量的类型，返回reflect.Type类型 
b. reflect.ValueOf()，获取变量的值，返回reflect.Value类型 
c. reflect.Value.Kind()，获取变量的类别，返回一个常量 
d. reflect.Value.Interface()，转换成interface{}类型
*/

//基本类型
// func main() {
// 	var x float64 = 3.4
// 	fmt.Println("type:", reflect.TypeOf(x))	//type: float64
// 	v := reflect.ValueOf(x)
// 	fmt.Println("value:", v)	//value: 3.4
// 	fmt.Println("type:", v.Type())	//type: float64
// 	fmt.Println("kind:", v.Kind())	//kind: float64
// 	fmt.Println("value:", v.Float())	//value: 3.4
// 	fmt.Println(v.Interface())	//3.4
// 	fmt.Printf("value is %5.2e\n", v.Interface())	//value is 3.40e+00
// 	y := v.Interface().(float64)
// 	fmt.Println(y)	//3.4
// }

//获取结构体
// type Student struct {
// 	Name  string
// 	Age   int
// 	Score float32
// }

// func test(b interface{}) {
// 	t := reflect.TypeOf(b)
// 	fmt.Println(t)

// 	v := reflect.ValueOf(b)
// 	fmt.Println(v)

// 	k := v.Kind()
// 	fmt.Println(k)

// 	iv := v.Interface()
// 	fmt.Println(iv)

// 	stu, ok := iv.(Student)
// 	if ok {
// 		fmt.Printf("%v %T\n", stu, stu)
// 	}
// }

// func main() {
// 	var a Student = Student{
// 		Name:  "stu01",
// 		Age:   18,
// 		Score: 92,
// 	}
// 	test(a)
// }

// main.Student
// {stu01 18 92}
// struct
// {stu01 18 92}
// {stu01 18 92} main.Student


//Elem
func main() {

	var b int = 1
	b = 200
	testInt(&b)
	fmt.Println(b)
}

//fv.Elem()用来获取指针指向的变量
func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	val.Elem().SetInt(100)
	c := val.Elem().Int()

	fmt.Printf("get value  interface{} %d\n", c)
	fmt.Printf("string val:%d\n", val.Elem().Int())
}

// get value  interface{} 100
// string val:100
// 100