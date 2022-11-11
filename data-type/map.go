package main
import "fmt"

func main(){
	person := map[string]string{
		"name":"Mujahid",
		"address":"Bandung",
	}
	person["title"] = "Student"


	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["title"])

	newArray := make(map[string]string)

	newArray["title"] = "Harry potter"
	newArray["author"] = "J.K Rowling"
	newArray["terbit"] = "2000"
	fmt.Println(newArray)

	delete(newArray, "terbit")
	fmt.Println(newArray)
}