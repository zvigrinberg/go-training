package main

import "fmt"

type interface1 interface {
	Method1()
}
type interface2 interface {
	Method2()
}
type interface3 interface {
	Method3()
}
type Interface4 interface {
	interface1
	interface2
	interface3
	Method4()
}

func (implementer Implementer) Method1() {
	fmt.Printf("Method 1 with values:  %s, %s , %d \n " , implementer.first , implementer.second , implementer.size )
}

func (implementer Implementer) Method2() {
	fmt.Printf("Method 2 with values : %s, %s , %d \n " , implementer.first , implementer.second , implementer.size )
}

func (implementer Implementer) Method3() {
	fmt.Printf("Method 3 with values: %s, %s , %d \n " , implementer.first , implementer.second , implementer.size )
}

func (implementer Implementer) Method4() {
	fmt.Printf("Method 4 with values : %s, %s , %d \n  " , implementer.first , implementer.second , implementer.size )
}

type Implementer struct {
	first string
	second string
	size int

}

func interface4Testers(instance Interface4) {
	instance.Method1()
	instance.Method2()
	instance.Method3()
	instance.Method4()

}


func main() {
 implementedInstance := Implementer{
	 first:  "hello",
	 second: "you",
	 size:   9,
 }
 interface4Testers(implementedInstance)

}
