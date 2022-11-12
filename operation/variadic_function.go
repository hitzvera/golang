package main

import "fmt"

func main(){

	firstSum := sumAll(1,2,32,4,2,1,21,3,)
	secondSum := sumAll(1,2,3)
	fmt.Println(firstSum, secondSum)

	// bisa juga dimasukin ke slice
	testSlice := []int {1,2,3,4}

	thirdSum := sumAll(testSlice...)

	fmt.Println(thirdSum)

}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}