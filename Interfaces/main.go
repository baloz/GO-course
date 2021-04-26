package main
import "fmt"

type bot interface{
	getGreeting() string
}
type englishbot struct{}
type spanishbot struct{}
func main(){
	eb := englishbot{}
	sb := spanishbot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishbot) getGreeting() string{
	return "Hi There!"
}
func (sb spanishbot) getGreeting() string{
	return "Hola!"
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}