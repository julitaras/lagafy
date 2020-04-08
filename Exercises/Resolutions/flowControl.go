package main

import (
	"fmt"
)

func exerciseOne() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

func exerciseTwo() {
	var year int

	fmt.Println("Ingrese su año de nacimiento:")
	fmt.Scan(&year)

	for year <= 2020 {
		fmt.Println(year)
		year++
	}
}

func exerciseThree() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("Cuando dividimos %v entre 4, el resto (también módulo) es %v\n", i, i%4)
	}
}

func esPar(number int) bool {
	if number%2 == 0 {
		return true
	}

	return false
}

func exerciseFour() {
	var number int

	fmt.Println("Ingrese un numero:")
	fmt.Scan(&number)

	if esPar(number) {
		if number > 0 {
			fmt.Printf("El numero %v es par y mayor a 0", number)
		} else if number < 0 {
			fmt.Printf("El numero %v es par y menor a 0", number)
		} else {
			fmt.Printf("El numero %v es par y igual a 0", number)
		}
	} else {
		if number > 0 {
			fmt.Printf("El numero %v es impar y mayor a 0", number)
		} else {
			fmt.Printf("El numero %v es impar y menor a 0", number)
		}
	}
}

func main() {
	var opcion int

	fmt.Println("Ingrese la opcion que quiere realizar:")
	fmt.Println("\tOpcion 1: Se imprimiran todos los numeros del 1 al 100")
	fmt.Println("\tOpcion 2: Se imprimiran todos los años que ha vivido")
	fmt.Println("\tOpcion 3: Se imprimira el resto de dividir entre 4, cada numero entre 10 y 100")
	fmt.Println("\tOpcion 4: Se imprimira si el numero ingresado es par o impar y si es mayor, menor o igual a 0")
	fmt.Scan(&opcion)

	switch opcion {
	case 1:
		exerciseOne()
	case 2:
		exerciseTwo()
	case 3:
		exerciseThree()
	case 4:
		exerciseFour()
	}

}
