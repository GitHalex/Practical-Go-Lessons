package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println((4 <= 6))
	name := "Alex"
	fmt.Println("Hello World " + name)

	var s string = "Hola"
	fmt.Println("valor de cadena => " + s)
	var i int = 100
	fmt.Println("valor entero => " + strconv.Itoa(i))
	var b bool = true
	fmt.Println("valor booleano => " + strconv.FormatBool(b))
	var f float64 = 77.90
	fmt.Println("valor float => " + strconv.FormatFloat(f, 'f', -1, 64))
	var nombre string = "CloudLex"
	fmt.Println("valor nombre => ", nombre, " ", "Otro")

}
