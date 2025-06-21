package main

import "fmt"

//const PI float64 = 3.14 // global constant

func main() {

	//fmt.Println("Hello World")
	//User input
	//var nombreEntradaInput string
	//fmt.Print("Enter your name: ")
	//fmt.Scanf("%s", &nombreEntradaInput)
	//fmt.Println("Hey there, ", nombreEntradaInput)

	//var otroNombreEntrada string
	//var is_muggle bool
	//var amount float32 = 5466.54

	//var i int = 42
	//var s string = strconv.Itoa(i)

	//fmt.Print("Enter your name & are you a muggle: ")
	//fmt.Scanf("%s %t", &otroNombreEntrada, &is_muggle)

	//fmt.Println(otroNombreEntrada, is_muggle)
	//fmt.Printf("variable de otroNombreEntrada = %v is of type %T\n ", otroNombreEntrada, otroNombreEntrada)
	//fmt.Printf("variable amount = %v is of type %T\n ", amount, amount)
	//fmt.Printf("%q", s)

	var city string = "Kolkata"
	var city_2 string = "Calcutta"
	fmt.Println(city == city_2)
	fmt.Println(city != city_2)

	//var i int = 1
	//i++
	//fmt.Println(i)
	//fmt.Println("###########################")
	//var x, y int = 12, 25
	//z := x & y
	//fmt.Println(z)

	//var x, y int = 100, 90
	//fmt.Println(x & y)
	//fmt.Println(x | y)

	//var x, y int = 100, 90
	//fmt.Println((x +

	var x, y int = 100, 90
	fmt.Println(!(((x + y) >> 2) == 47))

}
