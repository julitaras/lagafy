package main

import (
	"fmt"
)

func main() {

	var number int
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println("Tenemos el siguiente slice: ")
	fmt.Println(x)
	fmt.Println("Ingrese un numero para agregar al slice: ")
	fmt.Scan(&number)

	x = append(x, number)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}
