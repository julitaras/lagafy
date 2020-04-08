package main

import (
	"fmt"
)

func main() {

	var number int

	fmt.Println("Ingrese un numero:")
	fmt.Scan(&number)

	fmt.Printf("%d\t%b\t%#x\n", number, number, number)

	b := number << 1

	fmt.Printf("%d\t%b\t%#x", b, b, b)
}
