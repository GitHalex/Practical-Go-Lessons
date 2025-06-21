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

	fmt.Println("###############")
	//var a, b string = "kolkata", "Kolkata"
	//if a == b {
	//	fmt.Println("strings are equal")
	//} else {
	//	fmt.Println("strings are not equal")
	//}
	//fmt.Println("thank you!")

	var i int = 10
	switch i {
	case -5:
		fmt.Println("i is -5")
	case 10:
		fmt.Println("i is 10")
		fallthrough
	case 20:
		fmt.Println("i is 20")
		fallthrough
	default:
		fmt.Println("default")
	}

	day := "wednesday"
	switch day {
	case "monday":
		fmt.Println("monday")
	case "tuesday":
		fmt.Println("tuesday")
	case "wednesday":
		fmt.Println("wednesday")
		fallthrough
	case "thursday":
		fmt.Println("thursday")
		fallthrough
	case "friday":
		fmt.Println("friday")
	case "saturday", "sunday":
		fmt.Println("weekend")
	default:
		fmt.Println("default")
	}

}
