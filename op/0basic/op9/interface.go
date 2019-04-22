package main
import "fmt"

type phone interface {
	call()
}

type nokias struct {
}
func (nokia nokias) call() {
	fmt.Printf("a")
}

type iphone struct {
}
func (iphone iphone) call() {
	fmt.Printf("b")
}

func main() {
	var phone phone
	phone = new(nokias)
	phone.call()

	phone = new(iphone)
	phone.call()
}