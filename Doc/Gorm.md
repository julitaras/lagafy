# Gorm

Para la persistencia de datos utilizamos *Gorm*, una libreria ORM para go. Esta nos permite interactuar de forma amigable y sencilla con la base de datos. En nuestro caso el motor de base de datos utilizado es MySql

## Instalación:

Antes que nada debemos instalar la libreria mediante el siguiente comando

     go get -u github.com/jinzhu/gorm

Luego, para poder hacer uso del ORM necesitamos importar los paquetes:

```go
    import (
        "github.com/jinzhu/gorm"
        _ "github.com/jinzhu/gorm/dialects/mysql"
    )
```

Aclaracion: el segundo paquete solo es necesario importarlo cuando hagamos la conexion a la base de datos

## Setup:

Luego de instalar e importar lo primero que debemos hacer es conectarnos a la base de datos. Dentro del archivo main.go generamos la conexion

```go
    client, err := gorm.Open("mysql", "user:password@(localhost)/dbname?charset=utf8&parseTime=True&loc=Local")
```

Recordar siempre de disponer de un manejo de error en caso de que algo falle

Esta instancia es uno de los parametros que utiliza *SetRoutes()*

## Auto Migration

Con gorm podemos migrar de forma automatica los modelos de las entidades ahorrandonos asi tener que crear las tablas manualmente. El metodo *AutoMigrate()* creará tablas, columnas faltantes e índices faltantes, pero no cambiará el tipo de columna existente o eliminará columnas no utilizadas. La configuracion de las migraciones van en el archivo main.go

Para migrar un modelo usaremos la siguiente sintaxis:

```go
	client.AutoMigrate(&models.Travel{}, &models.Passenger{}, &models.Reservation{})
```

Cuando usemos una foreign key lo especificaremos de la siguiente forma:

```go
    client.Model(&models.Reservation{}).AddForeignKey("travel_id", "travels(id)", "RESTRICT", "RESTRICT")
```

Tambien sera necesario agregar en el modelo de las entidades los tags *gorm:"primary_key"* en la clave primaria y *gorm:"foreignkey:TravelID"* si tiene una entidades relacionada.

Para mas informacion sobre tags puede visitar el siguiente [link](http://gorm.io/docs/models.html#Supported-Struct-tags)

En caso de una relacion de *uno a muchos*:

```go
	client.Model(&models.Travel{}).Related(&models.Reservation{})
```

Aclaracion: los modelos deben contar con los campos *CreatedAt*, *UpdatedAt* del tipo time.Time y *DeletedAt* del tipo *time.Time

## CRUD

En nuestra arquitectura las operaciones crud del ORM se utilizan en los repository's. En la nomenclatura de los ejemplos se hace referencia al tipo *gorm.DB mediante la abstraccion repository compuesta por la primera letra de la entidad seguida de una r, un punto y el plural del modelo

### Create:

El metodo *Create()* recibe la direccion de memoria de la entidad a guardar

```go
    tr.travels.Create(&t)
```

### Get:

Para este caso encadenamos 3 metodos:
* *Where()* nos permite establecer condiciones
* *Preload()* trae los modelos asociados
* *Find()* todo lo que cumpla las condiciones se guarda en la direccion de memoria pasada como parametro

```go
    tr.travels.Where("id = ?", id).Preload("Reservations").Preload("Reservations.Passenger").Find(result)
```

En vez de *Find()* tambien se puede usar *First()*, la diferencia es que este ultimo solo trae el primer resultado que cumpla las condiciones

### GetOrCreate:

El metodo *FirstOrCreate()* guarda en la direccion de memoria pasada como parametro la primera coincidencia y en caso de no haber ninguna la crea

```go
    pr.passengers.Where("email = ?", email).FirstOrCreate(&pg)
```

### Delete:

Si el modelo tiene el campo *DeletedAt* el metodo *Delete()* funciona como un borrado logicoy no nos traera los datos aunque estos sigan estando en la base de datos

```go
    tr.travels.Where("id = ?", id).Delete(&models.Travel{})
```

### Update:

El metodo *Update()* actualiza solo los campos aclarados aunque tambien se le puede pasar un struct completo

```go
    tr.travels.Where("id = ?", t.ID).Update(t)
```

Para informacion adicional visitar la [documentacion oficial](http://gorm.io/docs/)