#!/usr/bin/env bash

package="github.com/fedejuret/zerossl-cli"

package_split=(${package//\// })
package_name=${package_split[-1]}

platforms=("linux/amd64" "linux/386" "linux/arm" "windows/amd64" "windows/386" "darwin/amd64")

# Manejar la señal de interrupción (SIGINT)
trap 'cleanup' INT

function cleanup() {
    # Detener el spinner
    kill -9 $spinner_pid > /dev/null 2>&1
    wait $spinner_pid > /dev/null 2>&1

    echo 'Script interrupted. Exiting...'
    exit 1
}

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    echo "Running for $GOOS $GOARCH ..."
    
    # Iniciar el spinner
    spinner="/-\|"
    i=0
    while true
    do
        printf "\r[%c] " "${spinner:$i:1}"
        sleep 0.1
        ((i++))
        ((i == ${#spinner})) && i=0
    done &
    spinner_pid=$!

    # Construir el binario para la plataforma actual
    env GOOS=$GOOS GOARCH=$GOARCH go build -o bin/$GOOS/$GOARCH/$output_name $package
    build_status=$?

    # Detener el spinner
    kill -9 $spinner_pid > /dev/null 2>&1
    wait $spinner_pid > /dev/null 2>&1

    # Verificar el resultado de la construcción
    if [ $build_status -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi

    echo "Build completed for $GOOS $GOARCH"
done
