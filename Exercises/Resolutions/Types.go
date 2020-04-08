package types

import (
	"fmt"
)

type numero int

var x numero
var y int

func main() {
	var number numero
	fmt.Println(x)
	fmt.Printf("El tipo de x es: %T\n", x)

	fmt.Println("Ingrese un numero:")
	fmt.Scan(&number)
	x = number
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
