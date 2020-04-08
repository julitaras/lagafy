# Ejercicio integrador inicial

## Crear una funcion Reverse que devuelva un string invertido al pasado por parametro.

Hints:

- La funcion len devuelve la cantidad de elementos de un arreglo.

- El tipo de dato rune es una representacion de un code point (valor numerico de un caracter en un codificacion determinada).
- La funcion de convertion string se puede usar para convertir un rune a un caracter, y viceversa.

```go
    r := []rune(s)
    s := string(r)
```