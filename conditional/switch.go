package main

import "fmt"

func main(){
	name := "Mujahid"

	switch name {
	case "mujahid":
		fmt.Println("helo semua")
	case "Mujahid":
		fmt.Println("halo semua ini mujahid")
		fmt.Println("ganteng sekali")
	default:
		fmt.Println("duh ga ada nih")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("kebanyakan tong")
	case false:
		fmt.Println("yes pas cocok!")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("duh panjang teuing")
	case length > 5:
		fmt.Println("duh masih panjang gan")
	default:
		fmt.Println("nah udah pas gening")
	}
}