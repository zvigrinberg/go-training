package main

import "fmt"

type Base struct {
	innerField int
	field2 string
}


type Outer struct { // Outer is the embedding struct
	Base                      // Base is the embedded struct
	field string
	field2 string
}

func (base Base) Describe() string {
	return fmt.Sprintf("base inner field's value=%d", base.innerField)
}

func main() {
	//direct assignment of innerField from Outer
	son := Outer{}
	son.innerField = 5
	son.field = "Hello there"
    fmt.Printf("This is the embeded value accesssed from Outer struct: %d" , son.innerField)

// assignment while declaring the outer struct, need to specify the whole path to the field this time, no shortcuts
   outerStruct := Outer{Base: Base{innerField: 10,field2: "I'm inside!"},field: "Hello Again!",field2: "I'm outside!"}
   //but accessing the value has shortcut of course, or when wanting to directly assign to the inner field
   fmt.Printf("Again With different value , this is the embeded value accesssed from Outer struct: %d" , outerStruct.innerField)

   //Calling method of the embedded struct from the embedding struct
   fmt.Println(outerStruct.Describe())

   //Shadowing field of embedded struct with field of embedding struct that has the same name
	fmt.Printf("This is the outer value accesssed from Outer struct: %s" , outerStruct.field2)
    fmt.Printf("This is the inner value accesssed from Outer struct: %s" , outerStruct.Base.field2)

}