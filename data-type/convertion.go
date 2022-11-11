package main

import "fmt"

func main(){
	var a int32 = 1000000
	var b = int64(a)
	var c = int16(a)

	fmt.Println(b, c)

	var name = "Mujahid"
	var m = name[0] // bakal jadi byte
	fmt.Println(m)
	var n = string(m)
	fmt.Println(n)

}