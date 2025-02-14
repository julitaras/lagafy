package main

import (
	"fmt"
)

func main() {
	x := map[string][]string{
		`eduar_tua`:    []string{`computadoras`, `montañas`, `playa`},
		`carlos_ramon`: []string{`leer`, `comprar`, `música`},
		`juan_bimba`:   []string{`helado`, `pintar`, `bailar`},
	}

	fmt.Println(x)

	x[`luis_perez`] = []string{`trabajar`, `playa`, `cerveza`}

	fmt.Println(x)
	delete(x, `juan_bimba`)

	for llave, valor := range x {
		fmt.Println("Registro:", llave)
		for i, val := range valor {
			fmt.Println("\t", i, val)
		}
	}
}
