#!/bin/bash
echo "Starting minikube"
minikube start

echo "Making Cert"
openssl req -new -newkey rsa:4096 -days 365 -nodes -x509 \
    -subj "/C=US/ST=Denial/L=Fairfax/O=Dis/CN=www.sgune.com" \
    -keyout tls-key.key  -out tls-cert.crt

echo "Creating k8s tls secret"
kubectl create secret tls sgune-cert --key tls-key.key --cert tls-cert.crt 

echo "cleaning stuff"
rm -f tls-key.key tls-cert.crt

echo "creating nginx-ingress"
kubectl apply -f ingress-controller 

echo "deploying stuff"
skaffold run 