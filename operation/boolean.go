package main
import "fmt"

func main(){

	var (
		a = 10
		b = 5
		thisIsTrue = a > b
		thisIsFalse = a < b
	)
	fmt.Println(thisIsFalse)
	fmt.Println(thisIsTrue)

	var testBoolean = thisIsFalse && thisIsTrue // false
	var testBoolean2 = thisIsFalse || thisIsTrue // true

	fmt.Println(!testBoolean && testBoolean2) // true



}