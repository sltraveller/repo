
package main
import "fmt"
/*
type books struct {
	title string
	subject string
	boolid int
}

func main() {
	fmt.Print(books{"go", "com", 1})
	fmt.Print(books{title: "go", subject: "cop"})
}

func main() {
	var book1 books
	var book2 books
	book1.title = "a"
	book1.subject = "wa"
	book1.boolid = 1

	book2.title = "b"
	book2.subject = "wb"
	book2.boolid = 2

	fmt.Printf(book1.title)

	//printbook(book1)

	printbook(&book2)
}

func printbook(book books) {
	fmt.Printf("%s \n", book.title)
	fmt.Printf("%s \n", book.subject)
	fmt.Printf("%d \n", book.boolid)
}


func printbook(book *books) {
	fmt.Printf("%s \n", book.title)
}
*/

type human struct{
	name string
	age int
	weight int
}

type student struct{
	human
	speciality string
}

func main(){
	mark := student{human{"mark",22,122},"q"}
	fmt.Println(mark.name)
	fmt.Println(mark.speciality)
	mark.speciality = "w"
	fmt.Println(mark.speciality)
}