package main

import (
	"fmt"
	"strings"
)

//const PI float64 = 3.14 // global constant

func addNumbers(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	suma, diference := addNumbers(20, 10)
	fmt.Println(suma, " ", diference)

	numero := 5
	resultado := factorial(numero)
	fmt.Println("factorial resultado ", resultado)

	x_d := func(l int, m int) int { return l * m }
	fmt.Println("############")
	fmt.Println(x_d)
	fmt.Printf("%T \n", x_d)
	fmt.Println(x_d(20, 30))

	fmt.Println("#####################ORO##############")
	y_d := func(l int, m int) int { return l * m }(20, 30)
	fmt.Printf("%T \n", y_d)
	fmt.Println(y_d)
	fmt.Println("#####################ORO##############")

	miniscula := func(s string) { fmt.Println(strings.ToLower(s)) }
	fmt.Printf("%T \n", miniscula)
	miniscula("Rachel")

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

	//slice_one := make([]int, 5, 10)
	//fmt.Println(slice_one)
	//fmt.Println(len(slice_one))
	//fmt.Println(cap(slice_one))

	//fmt.Println("#############")
	//arr := [5]string{"a", "b", "c", "d", "e"}
	//slice := arr[:4]
	//fmt.Println(slice)

	//arr := [5]string{"a", "b", "c", "d", "e"}
	//slice := arr[:4]
	//fmt.Println(arr)
	//fmt.Println(slice)
	//slice[1] = "x"
	//fmt.Println(arr)
	//fmt.Println(slice)

	slice := make([]int, 2, 5) // length 2, capacity 5
	fmt.Println(slice)         // [0 0]
	fmt.Println(len(slice))    // 2
	fmt.Println(cap(slice))    // 5

	slice = append(slice, 1, 2, 3) // ahora length 5 → [0 0 1 2 3]
	fmt.Println(slice)
	fmt.Println(len(slice)) // 5
	fmt.Println(cap(slice)) // 5

	slice = append(slice, 4) // excede capacidad → Go crea un arreglo nuevo
	fmt.Println(slice)
	fmt.Println(len(slice)) // 6
	fmt.Println(cap(slice)) // capacidad nueva → Go la define (generalmente duplica)
	fmt.Println(slice)

	codes := map[string]int{"en": 1, "fr": 2}
	value, found := codes["en"]
	fmt.Println(value, found)
	value, found = codes["hh"]
	fmt.Println(value, found)

	languages := map[string]string{"en": "English"}
	fmt.Println(languages)
	valor, found := languages["en"]
	fmt.Println(valor, found)
	languages["it"] = "Italian"
	fmt.Println(languages)

	delete(languages, "en")
	fmt.Println(languages)

	ascii_codes := map[string]int{}
	ascii_codes["A"] = 65
	_, foun := ascii_codes["B"]
	if foun {
		fmt.Println("key B was not found")
	}
	//sumaOfNumbers := addNumbers(2, 3)

	//fmt.Println(sumaOfNumbers)
}
