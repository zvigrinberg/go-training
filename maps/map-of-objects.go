package main

import (
	"fmt"

)

type sample struct {
 num1 int

}


func f(p string) {
	fmt.Println("function f parameter:", p)
}

func g(p string, q int) {
	fmt.Println("function g parameters:", p, q)
}

func awesome(text string) int {
	return len(text)
}

func main() {
	m := map[string]interface{}{
		"f": f,
		"g": g,
	}
	m["awesome-func"] = awesome
	for k, v := range m {
		switch k {
		case "f":
			v.(func(string))("astring")
		case "g":
			v.(func(string, int))("astring", 42)
			//function that returns value
		case "awesome-func":
			result := v.(func(string) int)("test me please!")
			fmt.Printf("length of string is : %d \n" ,result)

		}
	}
	//function that has receiver argument value
	m2 := make(map[string]func(sample) int, 0)
	m2["receiver"] = sample.sampleWithReceiver
	theSample := sample{num1: 50}
	receiver := m2["receiver"](theSample)
	fmt.Printf("value returned from method %d \n", receiver)


	
	//random := rand.Int()
	//random%=10
	//for i := 0; i < 10 ; i++ {
	//
	//
	//	}
	}

func (s sample) sampleWithReceiver() int {
      return s.num1
}





