package main

import "fmt"

var testThis int = 0

func variableShadowing() {

	// testThis := 1
	// The above line of code will throw an error.
	// This feature is known as variable shadowing.
	// But we can reinitialize it in other was as below

	var testThis int = 7
	fmt.Println(testThis)

	// Sometimes variable shadowing can create unexpected results

	j := 5

	if true {
		j := 1
		j++
	}

	fmt.Println(j)
}
