# Unit test 🔩
Para tener un Endpoint completo es necesario hacer el respectivo test para probar el correcto funcionamiento de su lógica.

Para ello debemos crear un usecase_test, será donde tendremos todos nuestros tests asociados a esa entidad y sus endpoints, mockearemos una base de datos y la consultaremos por medio de un repository_mock que es el encargado de crear una simulación de cómo funciona el repository original.

Nuestro usecase_test lo ubicaremos dentro de la carpeta “usecase” de la entidad cuyo Endpoint vamos a testear.

Dentro de nuestro test tendremos que:
1. Crear nuestra base de datos mockeada.
2. Agregar nuestras entidades de prueba.
3. Crear nuestro mock de usecase.
4. Llamar al endpoint que queremos testear. 
5. Hacer los casos que necesitamos comprobar, por ej, que no de error, que lo que devuelve el método sea lo que esperamos, etc.

Nuestro repository-mock lo ubicaremos dentro de la carpeta “repository” de la entidad cuyo Endpoint vamos a testear.

Crearemos una func que será la que posea la lógica igual a lo que hace nuestro repository original pero interactuando sobre la base de datos mockeada en el usecase-mock. 
Su nombre, sus parámetros y lo que devuelve deberá ser igual al original.

Cuando ya tengamos nuestro test hecho, si tenemos la extensión para go podremos correrlo desde la opción "run test" que aparecerá encima de cada test o, con el comando go test


A continuación un ejemplo de la entidad travel, Endpoint “GetById”

``travel_usecase_test.go``

```go
    package usecase

    import (
        "api-dashboard/models"
        "api-dashboard/travel/repository"
        _ "api-dashboard/travel/repository"
        "context"
        "testing"
        "time"
    )

    func TestTravelUsecase_GetById(t *testing.T) {

    mockDB := []models.Travel{}
    travel := models.Travel{Departure: time.Now().UTC().Add(time.Minute * time.Duration(20)), ID: 1, Reservations: []*models.Reservation{}}
    travel2 := models.Travel{Departure: time.Now().UTC().Add(time.Minute * time.Duration(45)), ID: 2, Reservations: []*models.Reservation{}}
    mockDB = append(mockDB, travel)
    mockDB = append(mockDB, travel2)
    mockRepo := repository.NewTravelRepositoryMock(&mockDB)

    timeoutContext := time.Duration(15 * time.Second)

    useCase := NewTravelUsecase(mockRepo, timeoutContext)

    ctx, cancel := context.WithTimeout(context.Background(), timeoutContext)
    defer cancel()

    foundTravel, foundError := useCase.GetById(ctx, "1")
    notFoundTravel, notFoundError := useCase.GetById(ctx, "3")

    //found travel
    if foundTravel == nil {
        t.Fatal("Expecting travel not to be nil", foundTravel)
    }
    if foundError != nil {
        t.Error("Expecting error to be nil", foundError)
    }
    if foundTravel.Reservations == nil {
        t.Error("Expecting reservations to be an array", foundTravel.Reservations)
    }

    //not found travel
    if notFoundTravel != nil {
        t.Fatal("Expecting travel not to be nil", notFoundTravel)
    }
    if notFoundError == nil {
        t.Error("Expecting error to be nil", notFoundError)
    }

    }
```

``travel_repository_mock.go``

```go
    package repository

    import (
        "api-dashboard/models"
        "api-dashboard/travel"
        "context"
        "errors"
        "strconv"
        "time"
    )

    type travelRepositoryMock struct {
        travels *[]models.Travel
    }

    //NewTravelRepositoryMock creates repository
    func NewTravelRepositoryMock(t *[]models.Travel) travel.Repository {
        travels := t
        return &travelRepositoryMock{travels: travels}
    }

    func (tr *travelRepositoryMock) GetById(ctx context.Context, id string) (*models.Travel, error) {
        idTravel, _ := strconv.Atoi(id)

    for _, p := range *tr.travels {
        if p.ID == uint(idTravel) {
            return &p, nil
        }
    }

    return nil, errors.New("No hay ningún viaje con ese Id.")
    }
```
