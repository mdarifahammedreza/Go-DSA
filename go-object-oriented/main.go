package main

import "fmt"

type class struct{
	Name string
	Age int
}
// *********************
// func(c class) this is called receiver and it is too type , Value Receiver vs Pointer Receiver
//-----------------
func(c class) SayHello(){
	fmt.Println("Hello",c.Name)
}

func main(){

	//object created
	p:= class{
		Name:"reza",
		Age:26,
	}
	p.SayHello()
	fmt.Println(p)
}
