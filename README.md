# k8s-hands-on


    HOSTNAME=$(hostname) go run main.go

    go build -o app ./main.go


Opción 1: Pasar la variable al ejecutar el binario

    HOSTNAME=$(hostname) ./app


Opción 2: Exportar la variable antes de ejecutar

    export HOSTNAME=$(hostname)
    ./app

Opción 3: Incrustar la variable en tiempo de compilación

    go build -o app -ldflags "-X 'main.Hostname=$(hostname)'" ./main.go  NO FUNCIONO
