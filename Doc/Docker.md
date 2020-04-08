# Qué es Docker
Docker es un software que permite crear entornos independientes y aislados para ejecutar aplicaciones. A estos entornos se les llama contenedores permitiendo así la ejecución de aplicaciones en cualquier máquina independientemente del sistema operativo que la máquina tenga instalado.

A diferencia de una máquina virtual, que tiene un sistema operativo completo, docker sólo comparte los recursos de la máquina anfitriona para ejecutar sus entornos. 

Los dos elementos más básicos de Docker son las imágenes y los contenedores.

# Images
Una imagen es una especie de plantilla, una captura del estado de un contenedor. 

Las imágenes se utilizan para crear contenedores y nunca cambian. Éstas se identifican por un ID y un nombre-versión.

# Docker file
Es un archivo de configuración que se utiliza para crear imágenes. En dicho archivo indicamos qué es lo que queremos que tenga la imagen, y los distintos comandos para instalar las herramientas.

# Containers
Son instancias en ejecución de una imagen. A partir de una única imagen, podemos ejecutar varios contenedores. Al poder tener la aplicación en varios contenedores, se podría distribuir los accesos a la aplicación mediante balanceadores.

Como las imágenes no cambian, cuando creas un contenedor a partir de una, y si cambias algo o instalas alguna herramienta durante su ejecución, esos cambios no se verán reflejados en la imagen cuando la ejecución se detenga.

Sin embargo, Docker va trackeando los cambios en los contenedores como si fuera una herramienta de control de versiones, así se puede crear una imagen que contenga dichos cambios.

# Volume
Los volúmenes se utilizan para mantener los datos más allá de la vida útil de su contenedor. Son espacios dentro del contenedor que almacenan datos fuera de él.

Los volúmenes son específicos de cada contenedor. Se pueden crear varios contenedores de una sola imagen y definir el volumen para cada uno y así compartir datos entre contenedores.

# Links
Sirven para enlazar contenedores entre sí, que están dentro de una misma máquina, sin exponer a los contenedores cuáles son los datos de la máquina que los contiene.


# Comandos de Docker

* ```docker --version```: comprueba la version de docker.

* ```docker pull <imagen>```: Descarga una imagen.

* ```docker images```: Lista las imagenes que tenemos descargadas.

* ```docker ps -a```: Muestra los contenedores en ejecucion.

* ```docker rm <contenedor>```: Elimina un contenedor

* ```docker rmi <id>```: Elimina la imagen indicada.

* ```docker run [flags] <imagen>```: Ejecuta una imagen de docker.

* ```docker logs <id>```: Muestra los registros del contenedor indicado.

* ```docker build -t <nombre>```: Construye una imagen.

* ```docker start [flags] <contenedor>```: Ejecuta un contenedor.

* ```docker stop [flags] <contenedor>```: Detiene la ejecucion de un contenedor.

* ```docker kill [flags] <contenedor>```: Detiene abruptamente la ejecucion de un contenedor.

# Docker Compose

Es una herramienta que nos permite definir aplicaciones con multiples contenedores y correrlas con un solo comando.

Requiere de 2 pasos:

1.- Crear un dockerfile que defina el entorno de la aplicacion.

2.- Crear un archivo docker-compose.yml para que corra en conjunto todos los contenedores.

Un archivo docker-compose.yml debe ser de la siguiente forma:

```docker
version: '3.0'
services:
  api:
    build: "./Back"
    ports:
      - "8080:8888"
  front:
    build: "./Front"
    ports:
      - "3000:3000"
```

## Comandos de Docker Compose

* ```docker-compose up```: Levanta la aplicacion.

* ```docker-compose down```: Detiene la aplicacion.

* ```docker-compose build```: Hace un *build* de la aplicacion.

* ```docker-compose logs -f -t```: Lista los registros.

* ```docker-compose ps```: Muestra las aplicaciones en ejecucion.

Para más informacion visitar el [enlace](https://docs.docker.com/compose/)

# Comandos de Dockerfile

## FROM

El comando *FROM* crea una nueva imagen a partir de otra imagen base. siempre debe situarse al inico del dockerfile. Su sintaxis puede ser de distintas formas:

```docker
FROM <imagen>
```

```
# En el tag se puede aclarar una version especifica o se puede utilizar la mas reciente con 'latest'
FROM <imagen>:<tag>
```

```docker
FROM <imagen>@<digest>
```

Tambien podemos ponerle un alias con el operados *AS* seguido de un nombre.

## MAINTAINER

Sirve para aclarar el nombre y contacto de la persona responsable de esa imagen. Es de caracter opcional y actualmente esta en desuso.

```docker
MAINTANER <nombre> <mail>
```

## RUN

*RUN* ejecuta un comando en el contenedor durante el *build* de la imagen.

```docker
RUN <comando>
```

## CMD

Sirve para ejecutar comandos. A diferencia de *RUN* la ejecucion sucede despues de crear un contenedor.
Su sintaxis viene en tres sabores:

```docker
CMD [<ejecutable>,<parametro1>,<parametro2>]
```

```
# Se ejecuta el ENTRYPOINT con los parametros aclarados
CMD [<parametro1>,<parametro2>]
```

```docker
CMD <comando> <parametro1> <parametro2>
```

## LABEL

Permite agregar metadata a las imagenes.

```docker
LABEL <key>=<value>
```

## EXPOSE

Establece los puertos de conexion.

```docker
EXPOSE <puerto>
```

## ENV

Comando que define variables de entorno.

```docker
ENV <variable>=<valor>
```

## ADD

Copia un archivo o directorio de la maquina *host* en la imagen. En caso de ser un archivo comprimido lo descomprimira en el destino.

```docker
ADD <origen> <destino>
```

## COPY

Muy parecido a *ADD*. Copia un archivo o directorio de la maquina *host* en la imagen. Es mas transparente y tiene menos procesos en la ejecucion.

```docker
COPY <origen> <destino>
```

## ENTRYPOINT

Nos permite especificar el ejecutable que usara el contenedor. Sintaxis:


```docker
ENTRYPOINT <ejecutable>
```

```docker
ENTRYPOINT [<ejecutable>, <parametro>]
```

## VOLUME

Esta instrucción crea un volumen como punto de montaje dentro del contenedor y es visible desde el host anfitrión marcado con otro nombre.

```docker
VOLUME <directorio>
```

## USER

Determina el nombre de usuario a utilizar cuando se ejecuta un contenedor.

```docker
USER <usuario>
```

## WORKDIR

Cambia el directorio por defecto donde ejecutamos los comandos *RUN*, *CMD* y/o *ENTRYPOINT*.

```docker
WORKDIR <directorio>
```

## ARG

Define una variable que el usuario puede pasar en el momento de hacer el *build* mediante el flag *--build-arg <variable>=<valor>*. Tambien permite definir un valor por defecto.

```docker
ARG <variable>[=<valor por defecto>]
```

Las variables definidas mediante *ENV* sobreescriben a las definidas con *ARG* en caso de tener el mismo nombre.

## ONBUILD

El comando *ONBUILD* sirve para definir instrucciones que se ejecutaran cuando se esta imagen sea usada como base para otra imagen.

```docker
ONBUILD <instruccion y parametros>
```

## STOPSIGNAL

Esta instruccion define la señal del sistema que se enviara al container para salir.

```docker
STOPSIGNAL <señal>
```

## HEALTHCHECK

*HEALTHCHECK* le dice a docker como chequear si el container esta funcionando.

```docker
HEALTHCHECK [opcion] CMD <comando>
```

Las distindas opciones son:

1.- ```--interval=<duracion>```

2.- ```--timeout=<duracion>```

3.- ```--start-period=<duracion>```

4.- ```--retries=<intentos>```

Tambien podemos anular el *HEALTHCHECK* que se hereda de la imagen base.

```docker
HEALTHCHECK NONE
```

## SHELL

Permite sobreescribir la consola por defecto.

```docker
SHELL [<ejecutable>, <parametro1>, <parametro2>]
```

Para más informacion visitar el [enlace](https://docs.docker.com/engine/reference/builder/)