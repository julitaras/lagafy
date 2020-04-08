package main

import "fmt"

type Persona struct {
	nombre   string
	apellido string
}

func main() {
	v := Persona{}
	p := &v
	p.nombre = "María"
	p.apellido = "Perez"
	fmt.Println(v.nombre, v.apellido)
}
