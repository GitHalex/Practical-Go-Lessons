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

	var grades [3]int
	fmt.Println(grades)

	var numeros [3]int = [3]int{10, 20, 30}
	fmt.Println(numeros)

	enteros := [3]int{10, 20}
	fmt.Println(enteros)

	n := [...]int{10, 20, 30}
	fmt.Println(n)

	var fruits [2]string = [2]string{"apples", "oranges"}
	fmt.Println(len(fruits))

	for index, element := range fruits {
		fmt.Println(index, "=>", element)
	}

	//arr := [5]bool{true, true, true, true}
	//for i := 0; i < len(arr); i++ {
	//	if arr[i] {
	//		fmt.Println(i)
	//	}
	//}

	//arreglo := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	//slice := arreglo[1:8]
	//fmt.Println(slice)
	//
	//sub_slice := slice[0:3]
	//fmt.Println(sub_slice)

	slice_one := make([]int, 5, 10)
	fmt.Println(slice_one)
	fmt.Println(len(slice_one))
	fmt.Println(cap(slice_one))

	//fmt.Println("#############")
	//arr := [5]string{"a", "b", "c", "d", "e"}
	//slice := arr[:4]
	//fmt.Println(slice)

	arr := [5]string{"a", "b", "c", "d", "e"}
	slice := arr[:4]
	fmt.Println(arr)
	fmt.Println(slice)
	slice[1] = "x"
	fmt.Println(arr)
	fmt.Println(slice)

}
