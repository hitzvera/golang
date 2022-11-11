package main
import "fmt"

func main(){

	name := "Mujahid"

	if name == "Mujahid"{
		fmt.Println(true)
	} else if name == "Ansori" {
		fmt.Println("Hello ansori")
	} else {
		fmt.Println(false)
	}

	// short statement and condition
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjag")
	} else {
		fmt.Println("Nama sudah benar")
	}
	

}