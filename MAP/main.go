package main
import "fmt"

func main(){
	colors := map[string] string {
		"red" : "#ff0000",
		"green" : "#00ff00",
		"blue" : "#0000ff",
	}
	fmt.Println(colors)

	cols := make(map[string] string)
	cols["key1"] = "val1"
	fmt.Println(cols)
	
	delete(cols, "key1")
	fmt.Println(cols)
}