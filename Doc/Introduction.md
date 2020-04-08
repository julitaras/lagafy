Otros comandos basicos:

    build     Compila paquetes y dependencias
    clean     Remueve archivos object     
    get       Descarga e instala paquetes y dependencias
    install   Compila e instala paquetes y dependencias
    test      Testea paquetes
    fmt       Formatea el codigo

## Variables

Hay distintas maneras de declarar variables

Por Default:
```go
    var a int 
```

Inferencia de tipo automatica:

```go
    var b = 10
```

Multiples asignaciones de manera simultanea:

```go  
    var c, d = 20, 30
```

Declaracion de variables cortas ( ``:=`` ):

```go
    e, f := 40, 50
```
De esta ultima forma, si la variable ``e`` y ``f`` no existen las crea y les asigna ese valor, si ya existen solo les asigna un nuevo valor

Tambien se pueden hacer transformaciones de tipos de variables:

```go
    i := 42
    f := float64(i)
```

## Funciones

Se pueden expresar de la siguiente manera:
```go
    func AddIntegers(a int, b int) int {
        return a + b
    }   
```

Pueden retornar varios valores:

```go
    func SumDifference(a int, b int) (int, int) {
        return a + b, a - b
    }
```

Y se usa de la siguiente manera:

```go
    func main() {
    
        var sum, diff = SumDifference(10, 20)
        fmt.Println("Sum & Difference of 10 and 20 are (", sum, ",", diff, ")")
    }
```

En la consola aparecera:

    Sum & Difference of 10 and 20 are ( 30 , -10 )




Cuando no prescindimos de algun valor que retorne la funcion, podemos asignar un espacio en blanco (``_``) que representa que ese valor no se va a guardar. Por ejemplo: 

```go
    func main() {
    
        var _, diff = SumDifference(10, 20)
        fmt.Println("Difference of 10 and 20 is", diff)
    }
```

En consola:

    Difference of 10 and 20 is -10

Tambien podemos asignarle un nombre a lo que retorna la funcion:

```go
    func Product(a int, b int) (prod int) {
        prod = a * b
        return
    }
```

## Paquete "fmt"

Este paquete implementa funciones de tipo Input/Output derivadas de C como printf o scanf. El formato de los "verbos" son como los de C pero simplificados

### Los verbos:

    %v: placeholder generico. Convierte automaticamente tu variable en un string con algunas opciones default. Util para imprimir tipos de datos primitivos como strings o numeros.

    %#v: imprime la variable en sintaxis de GO. Significa que se puede copiar el output y pegarlo en el codigo y estaria correcto sintacticamente. Util para trabajar con structs y slices.

    %T: Imprime el tipo de variable.

    %d: Imprime un entero en base decimal.

    %x y %X: imprime enteros en base hexadecimal.

    %f: imprime un float.

    %q: imprime un string citado.

    %p: imprime una direccion de puntero de una variable.

Para ver mas tipos de verbos pueden entrar a <https://godoc.org/fmt>


## Prints

### Print

Permite a los usuarios imprimir datos con formatos basicos.

```go
    const name, age = "Kim", 22
    fmt.Print(name, " is ", age, " years old.\n")
```
Output:

    Kim is 22 years old.

La funcion ``println`` es igual a ``print`` pero agrega una linea de corte


### Printf

Permite a los usuarios imprimir datos formateados.

```go
    var mystring = "Hello world"
    fmt.Printf("The string is %s", mystring)
```
Output:

    The string is Hello world

## Scans

### Scanln / Scan

```go
    package main

    import (
        "fmt"
    )

    func main() {
        fmt.Println("Escribi tu nombre: ")
        var input string
        fmt.Scanln(&input)
        fmt.Println("Hola",input)
    }
```

Si el nombre es "Homero" el output sera:

    Hola Homero


### Scanf

```go
    func main() {
        var st string
        fmt.Scanf("%s", &st)
        fmt.Println("La palabra es", st)
    }
```

Si ingresamos como palabra "arbol" el output sera:

    La palabra es arbol

### Fscan

Con el paquete "os" podemos tomar varios inputs a la vez

```go
    package main

    import (
        "fmt"
        "os"
    )

    func main() {

        fmt.Println("Ingrese su nombre y su edad separado")

        var firstName string
        var age int

        fmt.Fscan(os.Stdin, &firstName, &age)
        
        fmt.Println("Su nombre es", firstName,"y tiene", age, "años.")

    }
```

Donde podremos escribir por consola primero el nombre por ejemplo "Ana", presionamos enter, y despues la edad, por ejemplo "25", seguido de otro enter. El output sera:

    Su nombre es Ana y tiene 25 años.

Para poder saber el tipo de la variable se hace de la siguiente manera:

```go
    package main

    import "fmt"

    func main() {

        nro := 1234
        fmt.Printf("%T\n", nro)
    }
```
``\n`` se usa para que se escriba en la linea de abajo y no todo junto.

En la consola figurara:

    int

Para ver mas documentacion sobre Scan y Print o del paquete "fmt" en general pueden ingresar al link anteriormente mencionado <https://golang.org/pkg/fmt/>

## Paquetes Personalizados

Se pueden crear paquetes con variables y funciones para luego poder ser implementados en otros paquetes, por ejemplo:

Tenemos una carpeta ``proyecto`` que contiene adentro una carpeta llamada ``numeros`` que dentro tiene un archivo llamado ``numeros.go``

Dentro de este archivo, tenemos el siguiente codigo: 

```go
    package numeros

    var numero int = 100

    var Numero2 int = 300

    func Suma() int {
	    return numero + Numero2
    }
```
Para que las variables o funciones se puedan exportar la inicial de las mismas tienen que estar en mayuscula. En este caso, si tratamos de usar ``numero`` en otro paquete dira que no existe, pero ``Numero2`` y ``Suma`` si se podran usar.

Ahora, dentro de la carpeta ``proyecto`` tendremos el ``main.go``, para incluir el paquete anterior tenemos el siguiente codigo:

```go
    package main

    import (
        "fmt"
        "./numeros"
    )

    func main() {

        fmt.Println("El numero es", numeros.Numero2)
        fmt.Println("La suma de los dos numeros es:", numeros.Suma())
    }
```

Tambien se pueden importar packetes desde repositorios

```go
    import (
        "fmt"
        "github.com/sammy/foo"
    )
```

## Instrucciones de control

### If & Else

En el caso del if, la unica diferencia con otros lenguajes es que la condicion no lleva parentesis

```go
    func main() {
        var aa = 10
        if aa == 10 {
            fmt.Println("The number is 10!")
        } else if aa%2 == 0 {
            fmt.Println("Even number!")
        } else {
            fmt.Println("Odd number")
        }
    }
```

En consola: 

    The number is 10!

Tambien el if puede empezar con una instruccion corta antes de que se ejecute:

```go
    package main

    import (
        "fmt"
        "math"
    )

    func pow(x, n, lim float64) float64 {
        if v := math.Pow(x, n); v < lim {
            return v
        }
        return lim
    }

    func main() {
        fmt.Println(
            pow(3, 2, 10),
            pow(3, 3, 20),
        )
    }
```

En este caso ``v``, que es ``x`` potenciado a la ``n``, solo existe dentro del if.

En consola:

    9 20

### For

Lo mismo con el for, tambien puede ir sin parentesis:

```go
    func main() {
    
        for i := 1; i <= 10; i = i + 1 {
            fmt.Print(i, " ")
        }
    }
```

En consola:

    1 2 3 4 5 6 7 8 9 10

#### For con dos variables

En el for se pueden instanciar dos variables al mismo tiempo y asignarles un valor inicial.

```go
    package main

    import "fmt"

    func main() {
        for i, j := 0, 1; i < 10; i, j = i+1, j+1 {
            fmt.Println("i,j", i, j)
        }
    }
```

En consola: 

    i,j 0 1
    i,j 1 2
    i,j 2 3
    i,j 3 4
    i,j 4 5
    i,j 5 6
    i,j 6 7
    i,j 7 8
    i,j 8 9
    i,j 9 10

### While

En el caso del While, en golang, se escribe con un for:

```go
    func main() {
	    sum := 1
	    for ; sum < 1000; {
	    	sum += sum
	    }
	    fmt.Println(sum)
    }
```

Devolvera en consola: 

    1024

### Switch

```go
    func main() {
        fmt.Print("Go runs on ")

        switch os := runtime.GOOS; os {
        case "darwin":
            fmt.Println("OS X.")
        case "linux":
            fmt.Println("Linux.")
        default:
            fmt.Printf("%s.", os)
        }
    }
```

En la consola:

    Go runs on nacl.

## Arrays

```go
    package main

    import "fmt"

    func main() {
        var a [3]int
        fmt.Println(a)

        a[0] = 1
        a[1] = 2
        a[2] = 3
        //a := [3]int{1, 2, 3}
        fmt.Println(a)
    }
```

Se escribira en la consola: 

    [0 0 0]
    [1 2 3]

    Program exited.

## Slices

```go
    package main

    import "fmt"

    func main() {
        a := [...]int{1, 2, 3, 4, 5}
        sa := a[1: 4]

        fmt.Println("Before:", a)
        sa[0] = 22

        fmt.Println("After:", a)
    }
```

Escribe en la consola:

    Before: [1 2 3 4 5]
    After: [1 22 3 4 5]


## Runas

Los literales de runa son solo valores enteros de 32 bits (sin embargo, son constantes sin tipo, por lo que su tipo puede cambiar) Representan puntos de código de Unicode. Por ejemplo, la runa literal 'a' es en realidad el número 97.

```go
    package main

    import "fmt"

    func SwapRune(r rune) rune {
        switch {
        case 97 <= r && r <= 122:
            return r - 32
        case 65 <= r && r <= 90:
            return r + 32
        default:
            return r
        }
    }

    func main() {
        fmt.Println(SwapRune('a'))
    }
```

En este caso, la funcion ``SwapRune`` lo que hace es recibir una runa, si es una letra minuscula lo que hace convertirla a mayuscula y viceversa.

En este caso, la consola devolvera:

    A



## Archivos

```go
    package main

    import (
        "fmt"
        "io/ioutil"
    )

    func main() {
        data, err := ioutil.ReadFile("sample.txt")
        if err != nil {
            fmt.Println("File reading error", err)
            return
        }
        fmt.Println("Contents of file:", string(data))

        data := []byte("This is some information!")
        err := ioutil.WriteFile("write_data.txt", data, 0666)
        if err != nil {
            fmt.Println("There has been an error:", err)
            return
        }
    }
```

## Maps

```go
    package main

    import "fmt"

    func main() {
        m := map[string]int{
            "a": 1,
            "b": 2,
            "c": 3,
        }
        for key, value := range m {
            fmt.Println("Key:", key, " Value:", value)
        }
    }
```

En la consola:

    Key: a  Value: 1
    Key: b  Value: 2
    Key: c  Value: 3


## Structs

Son una coleccion de campos tipados que pueden contener metodos pero no son clases estrictamente hablando.

Por ejemplo:

```go
    type Employee struct {
        firstName string
        lastName string
        age int
    }
```

Estos pueden contener otros structs dentro:

```go
    type Address struct {
        city, state string
    }

    type Employee struct {
        firstName, lastName string
        age int
        address Address
    }

    var emp Employee

    emp.firstName = "SomeThing"
    emp.address = Address {
        city  : "AA",
        state : "CO",
    }
```

### Metodos

El receptor del método aparece en su propia lista de parámetros entre la palabra func y el nombre del método:

```go
    type Employee struct {
        firstName string
        lastName string
        age int
    }

    func (e Employee) Print() {
        fmt.Println("Employee Record:")
        fmt.Println("Name:", e.firstName, e.lastName)
        fmt.Println("Age:", e.age)
    }
```

El metodo se llama de la siguiente manera en el main:

```go
    package main

    import "fmt"

    func main() {

        var emp Employee

        emp.firstName = "Carlos"
        emp.lastName = "Jimenez"
        emp.age = 47

        emp.Print()
    }
```

En consola:

    Name: Carlos Jimenez
    Age: 47
    

## Punteros

Golang soporta punteros para actualizar valores pero no admite aritmetica de punteros como en C.

``*`` se usa como prefijo para definir un puntero para de un tipo dado.

```go
    func (e *Employee) updateAge(newAge int) {
        e.age = newAge
    }
    emp := Employee{
        age: 33,
    }

    fmt.Println("Before:", emp.age)
    emp.updateAge(34)
    fmt.Println("After:", emp.age)
```

En consola:

    Before: 33
    After: 34

## Interfaces

Un conjunto de métodos define un interfaz.

Una variable de tipo interfaz puede contener valores de cualquier tipo que implemente esos métodos.

```go
    package main

    import (
        "fmt"
        "os"
    )

    type Reader interface {
        Read(b []byte) (n int, err error)
    }

    type Writer interface {
        Write(b []byte) (n int, err error)
    }

    type ReadWriter interface {
        Reader
        Writer
    }
```

En este ejemplo el interfaz ``Reader`` implementa el metodo ``Read`` y el interfaz ``Writer`` implementa el metodo ``Write``.
Y el interfaz ``ReadWriter`` implementa ambas interfaces.

Una interfaz sin metodos se conoce como interfaz vacia:

```go
    interface{}
```

## Concurrencia

### Go - Routines

Golang provee un mecanismo sencillo para crear un nuevo "thread" *.

Se usa el keyword ``go`` antes de una llamada a una funcion (Go-Routines).

```go
    package main

    import (
        "fmt"
        "runtime"
    )

    func say(s string) {
        for i := 0; i < 5; i++ {
            runtime.Gosched()
            fmt.Println(s)
        }
    }

    func main() {
        go say("world")
        say("hello")
    }
```

En consola:

    hello
    world
    hello
    world
    hello
    world
    hello
    world
    hello

### Channels

Los Channels son pipes para conectar goroutines concurrentes.

Sirven para enviar y recibir valores entre dos goroutines.

    ch <- v    // Envía v por el canal ch.
    v := <-ch  // Recibe del canal ch, y  asigna el valor a v.

Como en los maps o slices, los canales deben crearse antes del primer uso:

```go
    ch := make(chan int)
```

Ejemplo: 

```go
    package main

    import "fmt"

    func sum(a []int, c chan int) {
        sum := 0
        for _, v := range a {
            sum += v
        }
        c <- sum  // send sum to c
    }

    func main() {
        a := []int{7, 2, 8, -9, 4, 0}

        c := make(chan int)
        go sum(a[:len(a)/2], c)
        go sum(a[len(a)/2:], c)
        x, y := <-c, <-c  // receive from c

        fmt.Println(x, y, x + y)
    }
```

En consola:

    -5 17 12