package main

//empty interface is something similar to Object class in java, can be used by code for any type if the type is not known.

import "fmt"

type test_struct struct {
	name1 string
	name2 string
}

func main() {
	var i interface{}
	printMe(i)

	i = 20
	printMe(i)

	i = "test"
	printMe(i)

	i = test_struct{
		name1: "test",
		name2: "me",
	}
	printMe(i)
}

func printMe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

