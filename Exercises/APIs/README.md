# APIs exercise

1. Hacer un endpoint que devuelva `` "Hello Wold" ``. 

    - Armar un archivo ``router`` en donde pondremos la respectiva URL y el ``metodo`` que nos traera el "Hello World".

    - En un archivo aparte, creamos el ``metodo`` que nos devolvera el StatusCode y el mensaje correspondiente.

2. Modelar una entidad ``Task`` de un "to do list" que contenga:

    - ID
    - Titulo
    - Descripcion
    - Fecha de creación
    - Fecha límite
    - Check
    - isDeleted (Delete logico)

3. Hacer un endpoint que cree una tarea a traves del modelo ``Task``.

    - Colocar la URL en el archivo router, creado anteriormente.
    - En otro archivo, crear el metodo que nos devolvera el StatusCode y donde se hara el llamado al ``usecase``.
    - Crear el archivo ```task_usecase``, donde pondremos la "logica de negocio". En este caso, donde llamaremos al metodo del ``repository`` que nos cree la task.
    - Crear el archivo ```task_repository``, donde haremos un mock con los datos.
    - Crear las interfaces correspondientes para el usecase y el repository.

4. Hacer un endpoint ``PUT`` que modifique el titulo y la descripción de una ``Task`` 

    - Armar en el archivo ``router`` la respectiva URL, donde el ``ID`` sea el parámetro.
    - En otro archivo, crear el metodo que nos devolvera el StatusCode y donde se hara el llamado al ``usecase``.
    - En el archivo ``task_usecase``, hacer la "logica de negocio", donde llamaremos al metodo del ``repository`` para que actualice los datos de la ``task``.
    - Hacer en el archivo ``task_repository`` un mock con los datos.
    - Crear en las interfaces correspondientes los metodos
    - En caso de que no se encuentre la task, devolver el error ``"No se ha podido modificar la tarea"``.

5. Hacer un endpoint que devuelva todas las tareas ``Task`` (GET)

    - Armar en el archivo ``router`` la respectiva URL
    - En otro archivo, crear el metodo que nos devolvera el StatusCode y donde se hara el llamado al ``usecase``.
    - En el archivo ``task_usecase``, hacer la "logica de negocio", donde llamaremos al metodo del ``repository`` para que traiga todas las ``tasks``.
    - Hacer en el archivo ``task_repository`` un mock con los datos.
    - Crear en las interfaces correspondientes los metodos
    
6. Hacer un endpoint que devuelva una ``Task`` especifica (GET)

    - Armar en el archivo ``router`` la respectiva URL, en donde el ``ID`` sea el parametro.
    - Crear el metodo que nos devolver el StatusCode y donde se hara el llamado al ``usecase``.
    - En el archivo ``task_usecase``, llamaremos al metodo del ``repository`` para que traiga la ``task`` con el ``ID`` especificado
    - Hacer en el archivo ``task_repository`` un mock con los datos.
    - Crear en las interfaces correspondientes los metodos

7. Hacer un endpoint ``DELETE`` el cual elimine de forma lógica una ``task``.

    - Armar en el archivo ``router`` la respectiva URL, en donde el ``ID`` sea el parametro.
    - Crear el metodo que nos devolver el StatusCode y donde se hara el llamado al ``usecase``.
    - En el archivo ``task_usecase``, llamaremos al metodo del ``repository`` para que elimine la ``task`` según el ``ÌD`` pasado como parámetro.
    - Hacer en el archivo ``task_repository`` un mock con los datos.
    - Crear en las interfaces correspondientes los metodos
    - En caso de no encontrar el elemento, devolver el error ``"No se ha encontrado la tarea"``.

