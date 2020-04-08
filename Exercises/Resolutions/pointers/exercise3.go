package main

import "fmt"

func swap(px, py *int) {
	tempx := *px
	tempy := *py
	*px = tempy
	*py = tempx
}
func main() {
	x := int(1)
	y := int(2)
	fmt.Println("x es", x)
	fmt.Println("y es", y)
	swap(&x, &y)

	fmt.Println("x ahora es", x)
	fmt.Println("y ahora es", y)
}
