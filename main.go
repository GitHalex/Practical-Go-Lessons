package main

//los packeges son la forma de GO de organizar y reutilizar codig
import (
	"fmt"
)

//const PI float64 = 3.14 // global constant

func main() {

	//fmt.Println("Hello World")
	//User input
	var nombreEntradaInput string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &nombreEntradaInput)
	fmt.Println("Hey there, ", nombreEntradaInput)

	var otroNombreEntrada string
	var is_muggle bool

	fmt.Print("Enter your name & are you a muggle: ")
	fmt.Scanf("%s %t", &otroNombreEntrada, &is_muggle)

	fmt.Println(otroNombreEntrada, is_muggle)
}
