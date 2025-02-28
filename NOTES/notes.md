# Notes


## 61. Golang: Empieza a escribir tu API

Guiding Principles of REST


https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj

    docker ps -l
    docker run --rm -dti --net host --name host --name golang golang bash
    docker rm -fv 88ac6b0b29df

Editar el codigo desde visual estudio y tenerlo disponible desde docker (Es una ayuda pero es mala practica) Si el contenedor es malicioso podria afectar el host.

    docker run --rm -dti -v $PWD/:/go --net host --name golang golang bash  (Unido a la red de host para no exponer el puerto 9090)

    docker exec -it golang bash

    curl localhost:9001/fg  (LA API sigue funcionando, se deberia obtener un 404)


## 62. Ultimos detalles



## 63. Crea un Dockerfile para tu aplicacion en golang.

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
        

## 64. Tip: ¿No puedes construir un Dockerfile por que no tienes Docker instalado?

    ricardoandre97/backend-k8s-hands-on:v1

## 65. Escribe manifiestos de Kubernetes para desplegar tu aplicación.

    kubectl get deployments.apps backend-k8s-hands-on -o yaml | grep -i Pull

    kubectl get deployments.apps backend-k8s-hands-on -o yaml | grep -i Pull -C 12

**Recordar que si se usa minikube es necesario utilizar un comando:**

    minikube image load k8s-hands-on (este no sirve o validar)
    eval $(minikube -p minikube docker-env)

    kubectl get pods -l app=backend

        diegoall@pho3nix:~/MAESTRIA_ING/k8s-hands-on/backend$ kubectl get svc
    NAME                              TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
    backend-k8s-hands-on              ClusterIP   10.111.106.142   <none>        80/TCP           22m


    kubectl get deployments



> Se debe ver el pod qeu esta resolviendo el servicio, es decir se esta haciendo la solicitud por el puerto 80. Cuando se hace esto, el servicio lo envia a un pod en el puerto 9090 y ese pod regresa la informacion con la hora.

O de esta otra forma:

> curl http://10.108.131.36:80


**El servicio ClusterIP que has configurado solo es accesible dentro del clúster de Kubernetes,** lo que significa que no puedes acceder a él directamente desde tu navegador. Para poder acceder a tu aplicación en Minikube desde el navegador, necesitas cambiar el tipo de servicio a **NodePort** o usar un inclusión de puerto con Minikube.


En teoria si accedo a esa IP por el puerto 80 deberia obtener respuesta del deployment.

    diegoall@pho3nix:~/MAESTRIA_ING/k8s-hands-on$ minikube service backend-k8s-hands-on
    |-----------|----------------------|-------------|--------------|
    | NAMESPACE |         NAME         | TARGET PORT |     URL      |
    |-----------|----------------------|-------------|--------------|
    | default   | backend-k8s-hands-on |             | No node port |
    |-----------|----------------------|-------------|--------------|
    😿  service default/backend-k8s-hands-on has no node port
    ❗  Services [default/backend-k8s-hands-on] have type "ClusterIP" not meant to be exposed, however for local development minikube allows you to access this !


El mensaje indica que el servicio tiene el tipo ClusterIP, lo que significa que no está expuesto fuera del clúster. Para acceder desde el navegador, tienes varias opciones:


**Opción 1: Usar minikube tunnel (Recomendado)**

Este comando creará un túnel que permitirá acceder al servicio como si tuviera una IP pública. 

    diegoall@pho3nix:~/MAESTRIA_ING/k8s-hands-on$ minikube tunnel
    Status:
            machine: minikube
            pid: 1530921
            route: 10.96.0.0/12 -> 192.168.59.100
            minikube: Running
            services: []
        errors: 
                    minikube: no errors
                    router: no errors
                    loadbalancer emulator: no errors

Luego, obtén la IP del servicio con:

    kubectl get svc backend-k8s-hands-on

Deberías ver una salida similar a esta:

    NAME                   TYPE        CLUSTER-IP      EXTERNAL-IP     PORT(S)        AGE
    backend-k8s-hands-on   ClusterIP   10.111.106.142 <minikube-ip>   80/TCP         5d

Pero no ... sale otra cosa ...

    diegoall@pho3nix:~/MAESTRIA_ING/k8s-hands-on$ kubectl get svc backend-k8s-hands-on
    NAME                   TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)   AGE
    backend-k8s-hands-on   ClusterIP   10.111.106.142   <none>        80/TCP    5d

Accede desde el navegador a:

    http://<minikube-ip>:80

> Paila, no me entrego la EXTERNAL-IP

Pero si entrega otras IPs:

    route: 10.96.0.0/12 -> 192.168.59.100

    diegoall@pho3nix:~/MAESTRIA_ING/k8s-hands-on$ curl http://192.168.59.100
    Application is up and running

Desde el browser ejecutando el script:

http://localhost:9001/

    {"time":"2025-02-27T23:32:21.186791186-05:00","hostname":""}

Pero desde minikube usando el tunnel recibo:

    Application is up and running

No se de donde sale este mensaje:

Sera que la imagen esta desactualizada?

    REPOSITORY                                 TAG               IMAGE ID       CREATED         SIZE
    k8s-hands-on                               latest            a1fb73b301d0   6 days ago      15.4MB



REVISAR LUEGO!!!

NO SE SABE QUIEN ESTA RESPONDIENDO:     Application is up and running


**Opción 2: Cambiar el tipo del servicio a NodePort**

Si prefieres una solución más permanente, cambia el Service a NodePort editando el YAML:

    minikube ip


http://192.168.59.100:30080/  (NO FUNCIONO)


**Opción 3: *Usar kubectl port-forward (Temporal)*

kubectl port-forward svc/backend-k8s-hands-on 8080:80

    diegoall@pho3nix:~/MAESTRIA_ING/k8s-hands-on$ kubectl port-forward svc/backend-k8s-hands-on 8080:80
    Forwarding from 127.0.0.1:8080 -> 9090
    Forwarding from [::1]:8080 -> 9090
    Handling connection for 8080
    Handling connection for 8080
    E0227 23:51:57.243751 1568811 portforward.go:413] an error occurred forwarding 8080 -> 9090: error forwarding port 9090 to pod 115becc2cfe250ca8a67cf986d8e54743104d678bcabbf8d4bde2cbcfdd71c7c, uid : exit status 1: 2025/02/28 04:51:57 socat[347190] E connect(5, AF=2 127.0.0.1:9090, 16): Connection refused
    error: lost connection to pod

## 66. Aprender a consumir el servicio que creaste.




## 67. Nota: ¿NO puedes ver el servicio Backend?