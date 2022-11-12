package main

import "fmt"

// functoin di go sebagai first class citizen
func main(){
	variableGoodBye := getGoodBye
	fmt.Println(variableGoodBye("Mujahid"))
}

func getGoodBye(name string) string{
	return "Good bye " + name
}