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


## 62. Ultimos detalles



## 63. Crea un Dokcerfile para tu aplicacion en golang.

https://dev.to/andrioid/slim-docker-images-for-your-go-application-11oo


diegoall@pho3nix:~/MAESTRIA_ING/k8s-hands-on$ file app
app: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), dynamically linked, interpreter /lib64/ld-linux-x86-64.so.2, BuildID[sha1]=a51bbf528cad26d298d5c9d8292c47c60fa3e683, with debug_info, not stripped


root@pho3nix:/home/diegoall/MAESTRIA_ING/k8s-hands-on# ldd app
        linux-vdso.so.1 (0x00007ffef1515000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x000076e238600000)
        /lib64/ld-linux-x86-64.so.2 (0x000076e238a00000)


        docker build -t k8s-hands-on:0.0.1 -f Dockerfile .

        docker run -d -p 9001:9001 --name k8s-hands-on k8s-hands-on

        docker rm -fv k8s-hands-on  (Buena practica recomendado)
        

##



##



##


## 