package main

import "fmt"

func main() {

	const a = 5
	fmt.Println(a)

	const (
		b = "hellop"
		c = 3
	)

	// a = 6;
	// tidak bisa diinisiasi ulang karena const

	fmt.Println(b)
	fmt.Println(c)
}
