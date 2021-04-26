package main
import "fmt"

func main(){
	colors := map[string] string {
		"red" : "#ff0000",
		"green" : "#00ff00",
		"blue" : "#0000ff",
	}
	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string] string){
	for key, value := range c{
		fmt.Println("Color : ",key,"has code : ", value)
	}
}