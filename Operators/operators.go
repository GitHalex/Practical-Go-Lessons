package main

import "fmt"

func main() {
	var a, b float64 = 24.4, 3.0
	fmt.Println(a / b)
	fmt.Println(int(a) % int(b))

	var i int = 100
	switch i {
	case 10:
		fmt.Println("10")
	case 100, 200:
		fmt.Println("i is either 100 or 200")
	}

}
