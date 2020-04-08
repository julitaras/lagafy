# About GO

## Instalación  🔧

1. Descargar la version para tu OS desde [https://golang.org/dl](https://golang.org/dl).
2. Abrir el instalador y seguir las [instrucciones](https://golang.org/doc/install).
    
    Go quedara instalado por default en:
    
        Unix: /usr/local/go
        Windows: C:\Go

3. Workspace

    La variable de entorno GOPATH especifica la ubicacion del entorno de trabajo.

    Por defecto es:
    
        Unix: $HOME/go
        Windows: %USERPROFILE%\go

    El workspace contiene:

        GOPATH
            |_ BIN: Contiene los archivos ejecutables
            |_ SRC: Contiene los fuentes organizados en paquetes
            |_ PKG: Contiene los paquetes en formato binario

    Es conveniente agregar [GOPATH](https://golang.org/doc/code.html#GOPATH)/bin al PATH.

4. Para chequear que la instalacion esta correctar tipear en la consola 
    
    ```go
        $ go version
    ```

5.  Para chequear que todas las variables de entorno estan correctas tipear en la consola

    ```go
        $ go env 
    ```
        GOROOT = [Directorio de instalacion]
        GOPATH = [Directorio del Workspace, tiene que ser distinto al de instalacion]

- **Recomendacion: Seguir el [Tour of Go](https://tour.golang.org/list)**

# Hello World 🌎

Todo programa de GO esta hecho de paquetes.

Los programa se empienzan a ejecutar con el paquete ``main``

En este programa utilizamos el paquete ``"fmt"``

```go
    package main

    import (
	"fmt"
    )

    func main() {
        fmt.Println("Hello World")
    }
```

Corremos en la consola, prados en la carpeta donde se encuentra nuestro archivo ``.go``

```go
      $ go run [nombre del archivo]
```

Como respuesta en la consola nos aparece:

    Hello World


# **Para poder hacer una API hay cosas que tenemos que aprender de [Go language](Introduction.md)**


# Lagafy - ¿Cómo hicimos nuestras Apis ?

Utilizamos una libreria llamada ``gin``, con la que podemos acceder a las peticiones HTTP.

#### Gin Web Framework - Caracteristicas
- Speed
- Crash-Free
- Routing
- JSON Validation
- Error Management
- Built-In Rendering

**Para leer mas sobre [Gin](https://github.com/gin-gonic/gin)**

---------
## Hello World 🌎

Para ir entendiendo como utilizar la librería Gin.
El ¡Hello World! Con Gin sería algo parecido a esto:

```go
    package main

    import "github.com/gin-gonic/gin"

    func main() {
        r := gin.Default()
        r.GET("/", func(c *gin.Context) {
            c.JSON(200, gin.H{
                "message": "¡Hello World!",
            })
        })
        r.Run() // En el puerto 8080
    }
```
---

## ¿Cómo nos organizamos para implementar Gin?

Identificamos los modelos a utilizar. Por lo que, por cada entidad teniamos una carpeta con su respectivo nombre, cada una con su ``usecase`` (Logica de negocio) y su ``repository`` (Llamado a la DB, servicios, etc)

Para hacer un mejor entendimiento:
        
    travel
    |_repository
        |_travel_respository.go
        |_travel_repository_mock.go
    |_usecase
        |_travel_usecase.go
        |_travel_usecase_test.go
    |_respository.go 
    |_usecase.go

Siendo ``respository.go`` la interfaz de ``travel_respository.go`` del paquete ``travel`` y a su vez ``usecase.go`` la interfaz de ``travel_usecase.go``

## Donde hacemos el ruteo?

Colocamos las rutas en un archivo aparte (``router.go``), cada una en su respectivo handler. Separamos los handlers por entidad.

A su vez por cada entidad tenemos un router, es decir en el caso de la entidad Travel tenemos un archivo llamado ``travel_router.go``. En el cual se encuentran todos los metodos que se llaman en el router. Estos metodos se encargan de hacer el llamado al usecase y dar la response con el StatusCode correspondiente. En este archivo tambien colocamos los comentarios que corresponden a swagger, para leer mas sobre esto: [Swagger](Swagger.md) 

En el archivo ``router.go``:

El handler recibe un grupo de ruta, en nuestro caso seria "/api" y por otro lado recibe la interfaz del usecase.

```go
    func NewTravelHandler(g *gin.RouterGroup, uc travel.Usecase) {

        g.GET("/travel/:id", travelHandler.GetByID)
    }
```

En el archivo ``travel_router.go``:

```go
    func (p *travelHandler) GetByID(c *gin.Context) {
        id := c.Param("id")

        v, err := p.Usecase.GetById(c, id)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
                "message": err.Error(),
            })
            return
        }
        c.JSON(http.StatusOK, v)
        return
    }
```
Este metodo nos llevaria entonces al metodo ``GetByID`` que se encuentra en el archivo ``travel_usecase``. Basicamente se encarga de llamar al metodo que se encuentra en el ``travel_repository``

Archivo  ``travel_usecase``:

```go
    func (uc *travelUsecase) GetById(ctx context.Context, id string) (*models.Travel, error) {
        res, err := uc.tr.GetById(ctx, id)
        if err != nil {
            return nil, err
        }
        return res, nil
    }
```

Archivo ``travel_repository``:

Estamos buscando en la tabla Travels que se encuentra en la DB, el viaje correspondiente al id.

```go
    func (tr *travelRepository) GetById(ctx context.Context, id string) (*models.Travel, error) {
        result := &models.Travel{}
        if err := tr.travels.Where("id = ?", id).Where("template = ?", false).Preload("Reservations").Preload("Reservations.Passenger").Find(result).Error; err != nil {
            return nil, err
        } else {
            return result, nil
        }
    }
```

**Para entender mejor, leer [Gorm](Gorm.md)**

## Error Handling
- Los errores en Golang no son excepciones.
- Por convencion, si una funcion puede fallar deber retornar un tipo error.
- El tipo error contiene la infomacion del problema.  
- Si es nil quiere decir que no hubo errores

    Podemos observar  en los ejemplos anteriores como hacemos el manejo de errores.

## Acerca de la ``seguridad`` en nuestras APIs: [Leer aqui](Security.md)

# Unit test ⚙️
No podemos olvidarnos de los test. Los mismos lo desarrollamos en el archivo ``travel_usecase_test.go`` siendo el caso de la entidad ``Travel``

Podemos encontrar la explicacion de los mismos: [Test](Test.md)

---

# Git Hooks

Con Git podemos crear ramas de desarrollo, registrar cambios y tener un control absoluto sobre las versiones. Sin embargo, es posible automatizar este proceso. La automatización de Git funciona a nivel de programas y deployment. Y para eso existen los Hooks.

Los Hooks de Git son scripts de shell que se ejecutan automáticamente antes o después de que Git ejecute un comando importante como Commit o Push. Para que un Hook funcione, es necesario otorgarle al sistema Unix los permisos de ejecución. Mediante el uso de estos scripts, podemos automatizar ciertas cosas.

Encontraremos una pequeña guia para agregar Hooks a nuestro código: [Hooks](GitHooks.md)


# Docker 🐳

Para poder levantar nuestra aplicacion de forma independiente utilizamos ``Docker``.

Para saber de que se trata leer este archivo: **[Docker](Docker.md)**