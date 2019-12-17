package main

import "fmt"

var (
	studentFname string
	studentLname string
	studentClass int
	studentAge   int
)

func variable() {
	var i int
	i = 5
	fmt.Println(i)

	var j int = 7
	fmt.Println(j)

	x := 10
	fmt.Println(x)

	studentFname = "Akhil"
	fmt.Println(studentFname)

}
