package main
import "fmt"
// var map_variable map[key_data_type]value_data_type
// map_variable := make(map[key_data_type]value_data_type)

func main() {
	var counts map[string]string
	counts = make(map[string]string)

	counts["a"] = "q"
	counts["b"] = "w"
	counts["c"] = "e"

	for count := range counts {
		fmt.Println(count)
	}
	ca, ok := counts["c"]
	fmt.Println(ca)
	fmt.Println(ok)

	delete(counts, "b")
	for count := range counts {
		fmt.Println(count)
	}
}