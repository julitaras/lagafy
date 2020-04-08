# Swagger
Es un framework para documentar APIs Rest que te pe permite  describir, producir, consumir y visualizar APIs desde muy diferentes fuentes.
Pasos para poder usarlo
1)	Descargar ``swag`` para go
    
        $ go get -u github.com/swaggo/swag/cmd/swag

2)	Descargar ``gin-swagger``

```go
        import "github.com/swaggo/gin-swagger" // gin-swagger middleware
        import "github.com/swaggo/files" // swagger embed files
```
3)	Dentro de ``main.go`` añadir comentarios describiendo características del proyecto y la configuración tales como nombre, título, descripción, el host y el tipo de seguridad.

    Ejemplo:

```go
        // @title Swagger Lagafy API
        // @version 1.0
        // @description This is a sample server celler server.
        // @termsOfService http://swagger.io/terms/

        // @contact.name API Support
        // @contact.url http://www.swagger.io/support
        // @contact.email support@swagger.io

        // @license.name Apache 2.0
        // @license.url http://www.apache.org/licenses/LICENSE-2.0.html

        // @host localhost:8888
        // @BasePath /api/

        // @securityDefinitions.apikey ApiKeyAuth
        // @in header
        // @name Authorization
```
Y dentro de la función ``main`` poner:

```go
    r := gin.Default()
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
 ```

4)	Añadir comentarios en los que describa las características de la API tales como el título o un tag para dividir las apis según la entidad o algún otro parámetro que se elija. 

    En los comentarios también se describe los parámetros o el modelo que se espera recibir así como el que se devuelve, la uri, el tipo de request y el tipo de seguridad con el que cuenta.

    Ejemplo de una API para crear un viaje.

```go
        // @Summary Create new travel
        // @Tags travels
        // @Accept  json
        // @Produce  json
        // @Param travel body models.Travel true "Create a travel"
        // @Success 200 {object} models.Travel
        // @Router /travel/ [post]
        // @Security ApiKeyAuth
 ```
 
5)	Luego, correr el comando ``“swag init”``, el cual generará los archivos necesarios para que swagger recopile la información en los comentarios anteriormente puestos. 

    ``Aclaración:`` se debe ejecutar el comando cada vez que se modifiquen estos o se agreguen comentarios nuevos.

6)	Luego, levantar la aplicación e ir a [http://localhost:8888/swagger/index.html](http://localhost:8888/swagger/index.html) en este caso.
