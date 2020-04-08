package main

import "fmt"

func main() {
	a := 3
	p := &a
	fmt.Println(a, *p)
	*p = *p + 23
	fmt.Println(a, *p)
	a = a - 10
	fmt.Println(a, *p)
	*p = *p * 2
	fmt.Println(a, *p)
	a = a / 3
	fmt.Println(a, *p)

}
