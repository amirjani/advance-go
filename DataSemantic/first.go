package main

import "fmt"

func main() {
	var (
		number      int
		sentence    string
		floatNumber float32
		isRight     bool
	)

	fmt.Printf("number: \t %T [%v]\n", number, number)
	fmt.Printf("sentence: \t %T [%v]\n", sentence, sentence)
	fmt.Printf("floatNumber: \t %T [%v]\n", floatNumber, floatNumber)
	fmt.Printf("isRight: \t %T [%v]\n", isRight, isRight)
}
