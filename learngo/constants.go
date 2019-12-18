package main

import "fmt"

const a = iota	// Enumerated constant

const (
	_ = iota // This is done since if we initiazed a var and doesn't assign a value to it, BY default its gonna create an issue
	success	
	failed
)

func constants_learn() {

	const intConst int = 10
	const stringConst string = "This is a string const"
	const floatConst float64 = 3.14
	const boolConst bool = false	
	var status int

	fmt.Printf("%v\n", status == success)

	// myConst = 27  // Uncommenting this line is gonna throw an error
	fmt.Printf("%v, %T\n", intConst, intConst)
	fmt.Printf("%v, %T\n", stringConst, stringConst)
	fmt.Printf("%v, %T\n", floatConst, floatConst)
	fmt.Printf("%v, %T\n", boolConst, boolConst)

}
