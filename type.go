package main

import "fmt"

func typeprint() {
	var j float32 = 64
	// This is how type of variable got printed on STDOUT
	fmt.Printf("%v, %T\n", j, j)

	var b bool = true
	fmt.Printf("%v, %T\n", b, b)

	

}
