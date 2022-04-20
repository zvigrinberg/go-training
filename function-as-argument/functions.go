package main

import "fmt"

type function1 func(int)

type function2 func(string)

func myFunction(num int) {
	fmt.Printf("\nnum is %v", num)
}
func myFunction2(num int) {
	fmt.Printf("\nnum is %v", num)
}
func test(f function1, val int) {
	f(val)
}

func test2(f function2, val string) {
	f(val)
}

func myFunction3(something string) {
	fmt.Printf("\nsomething is %v", something)
}

func main() {
	test(myFunction, 123)
	test(myFunction2, 321)
	test2(myFunction3, "something")

}

