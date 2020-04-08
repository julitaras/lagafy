# Git Hooks

Los Hooks se encuentran en el directorio oculto ``.git/hooks`` del repositorio el cual posee ejemplos de cada hook.

Nos centraremos en el script de ``pre-commit`` ya que será el cual nos avise al momento de commitear si nuestro código cumple con las reglas pactadas.

Para empezar debemos saber que si utilizamos el directorio ``.git/hooks`` para guardar nuestros Hooks no se verán los cambios en el repositorio si no que quedarán de forma local.
Para poder versionarlos y que toda persona que baje el repositorio tenga los mismos Hooks deberemos hacer un nuevo directorio, por ej ``.githooks``

A continuación ejecutar el siguiente comando:


    git config core.hooksPath .githooks


Esto hará que Git sepa que los Hooks están alojados en ese directorio y pueda ejecutarlos.

Ahora bien, podemos crear nuestro propios scripts bash dentro de pre-commit o sobre un archivo .sh y ejecutarlo desde nuestro archivo pre-commit.
Para ambos casos deberemos ejecutar el comando:

    chmod +x pre-commit 
o 

    chmod +x rutaDelArchivo

Esto lo hacemos para darle permisos de ejecución[ al script.

En el siguiente repositorio encontraremos varios scripts enfocados para Golang que podremos usar para nuestros proyectos: [Git Hooks](https://github.com/vitessio/vitess/tree/master/misc/git/hooks)


Más ejemplos e información:
- [Useful Git hooks for developing in Go](https://discuss.dgraph.io/t/useful-git-hooks-for-developing-in-go/690)
- [Go Development Workflow With Git Hooks](https://tutorialedge.net/golang/improving-go-workflow-with-git-hooks/)

