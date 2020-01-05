package main

import "fmt"

func arrayIntro() {	
	vechiles := [3]string{"car", "bus", "van"}	
	fmt.Printf("Vechiles: %v \n", vechiles)

	var students [3]string
	students[2] = "Akhil"
	students[0] = "Jithesh"
	fmt.Printf("%v\n", len(students))
	// Creating array of arrays (Matrix)
	var myMatrix [3][3]int = [3][3]int{ [3]int{1, 2, 3}, [3]int{3, 4, 5}, [3]int{6,7,8} }
	fmt.Printf("Matrix: %v\n", myMatrix)

	var myOtherMatrix [3][3]int
	myOtherMatrix[0] = [3]int{1, 2, 3}
	myOtherMatrix[1] = [3]int{3, 4, 5}
	myOtherMatrix[2] = [3]int{4, 5, 6}
	fmt.Printf("Matrix: %v", myMatrix)
}
