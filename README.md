# cpulimiteimporta
Para revisar como funciona un artefacto con un limite bajo y con un limite alto de cpu.

Los archivos fueron tomados del repo:
https://github.com/chiahan1123/docker-hpa-example

Gracias!

# TODO
Mejorar la documentacion................

## Instalar


### Conectar


# Prueba

kubectl apply -f manifests/manifests.yaml


Para poder hacer la prueba se va a ocupar hey. https://github.com/rakyll/hey 

## Ejecutando hey

Prendiendo un container con ubuntu:
kubectl run hey --image=ubuntu -i --tty --rm

instalar hey:
apt update; apt install wget -y

wget https://hey-release.s3.us-east-2.amazonaws.com/hey_linux_amd64

mv hey_linux_amd64 hey

chmod +x hey

./hey -n 500000 -c 128 http://stress-app

