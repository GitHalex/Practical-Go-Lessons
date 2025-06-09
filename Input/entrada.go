package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println("Hello World")
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanf("%s ", &name)
	fmt.Println("Hey there: ", name)
	fmt.Println(
		"%T",
		name,
	)

	var grades int = 42
	var message string = "Hello world"
	var isCheck bool = true
	var amount float32 = 5466.54

	fmt.Println("variable grades = %v is of type %T\n", grades, grades)
	fmt.Println("variable message = %v is of type %T\n", message, message)
	fmt.Println("variable isCheck = %v is of type %T\n", isCheck, isCheck)
	fmt.Println("variable amount = %v is of type %T\n", amount, amount)
	fmt.Println("Type %v \n", reflect.TypeOf(message))
	fmt.Println("%q", strconv.Itoa(grades))

}
