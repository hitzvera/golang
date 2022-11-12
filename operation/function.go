package main

import "fmt"

func main(){
	sayHello()
	sayHello2("Mujahid", "Ansori")
	fmt.Println(add(1,2))

	firstName, secondName := getFullName() // jika pengen menghiraukan beberapa value maka bisa pake ' _ '
	fmt.Println(firstName, secondName)

	namaPertama, namaKedua, namaKetiga := getCompleteName()
	fmt.Println(namaPertama)
	fmt.Println(namaKedua)
	fmt.Println(namaKetiga)
}

func sayHello(){
	fmt.Println("Hello")
}

func sayHello2(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}

func add(a int, b int) int {
	return a + b
}

// return multiple value
func getFullName() (string, string){
	return "Mujahid", "Ansori"
}

// name return value
func getCompleteName() (firstName, secondName, lastName string) {
	firstName = "Mujahid"
	secondName  = "Ansori"
	lastName = "Majid"

	return // otomatis nge return firstname, secondName, lastName,
}