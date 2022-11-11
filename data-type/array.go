package main
import "fmt"

func main(){
	var names [3]string // 3 disini banyaknya element/item dari array tersebut
	names[0] = "Mujahid"
	names[1] = "Ansori"
	names[2] = "Majid"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	// names[3] = "test" akan error

	// atau bisa juga seperti ini
	var values = [3]int{
		1,2,3,
	}

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values)) // mencetak batas maksimal array (3)
}