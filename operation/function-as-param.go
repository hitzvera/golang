package main

import "fmt"


type Filter func(string) string

func main(){
	spamFilter := sensorBadWords
	fmt.Println(sayHelloWithFilter("Anjing", sensorBadWords))
	fmt.Println(sayHelloWithFilter("Mujahid", spamFilter))
}

func sayHelloWithFilter(name string, filter Filter) string{
	return filter(name)
}

func sensorBadWords(name string) string {
	if name == "Anjing" {
		return "*****"
	}
	return name
}