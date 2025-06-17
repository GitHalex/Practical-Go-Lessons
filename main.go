package main

//los packeges son la forma de GO de organizar y reutilizar codig
import (
	"fmt"
	"strconv"
)

const PI float64 = 3.14 // global constant

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
	fmt.Println("el valor %v", nombre)

	fmt.Println("#############################")
	var radius float64 = 5.154
	var area float64

	area = PI * radius * radius
	fmt.Printf("Radius: %.2f \nPI:%.1f \n", radius, PI)
	fmt.Printf("Area of Circle is : %f", area)
	fmt.Println("Thank You")

}
