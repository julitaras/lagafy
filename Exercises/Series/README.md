# Series

Dada una cadena de dígitos, genere todas las subcadenas contiguas de longitud "n" en esa cadena

Por ejemplo, la cadena "49142" tiene las siguientes series de 3 dígitos:

- 491
- 914
- 142

Y las siguientes series de 4 dígitos:

- 4914
- 9142

Tenga en cuenta que las series sólo necesitan ocupar *posiciones adyacentes*; los dígitos no necesitan ser *numéricamente consecutivos*

## Correr los test

Para correr los test usa el comando `go test` desde dentro del directorio del ejercicio.

Si los test contienen benchmarks, los puedes correr con `-bench`
flag:

```go
    go test -bench .
```