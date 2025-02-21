# Notes


## 61. Golang: Empieza a escribir tu API

Guiding Principles of REST


https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj

    docker ps -l
    docker run --rm -dti --net host --name host --name golang golang bash
    docker rm -fv 88ac6b0b29df

Editar el codigo desde visual estudio y tenerlo disponible desde docker (Es una ayuda pero es mala practica) Si el contenedor es malicioso podria afectar el host.

    docker run --rm -dti -v $PWD/:/go --net host --name golang golang bash

    docker exec -it golang bash

    curl localhost:9001/fg  (LA API sigue funcionando, se deberia obtener un 404)


##



##


##



##



##


## 