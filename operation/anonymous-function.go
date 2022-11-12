package main

import "fmt"

type Blacklist func(string) bool
func main(){
	registerUser("Mujahid", func(name string) bool {
		return name == "Mujahid"
	})
}

func registerUser (name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

