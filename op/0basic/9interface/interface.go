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



/*
type salarycalculator interface{
	calculatesalary() int
}

type contract struct{
	empid int
	basicpay int
}

type permanent struct{
	empid int
	basicpay int
	jj int
}

func (p permanent) calculatesalary() int{
	return p.basicpay+p.jj
}
func (c contract) calculatesalary() int {
	return c.basicpay
}

func totalexpense(s []salarycalculator){
	expense := 0
	for _,v := range s{
		expense += v.calculatesalary()
	}
	fmt.Printf("%d",expense)
}
func main(){
	pemp1 := permanent{1,222,333}
	pemp2 := permanent{2,333,444}
	cemp1 := contract{5,666}
	employes := []salarycalculator{pemp1,pemp2,cemp1}
	totalexpense(employes)
}
*/