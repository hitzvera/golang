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


}