#!/usr/bin/env zsh

brew install minikube
brew install helm

minikube start
minikube addons enable dashboard
minikube addons enable metrics-server

helm install ingress-nginx ingress-nginx/ingress-nginx

kubectl create secret generic pubg-api-token --from-literal=pubg-api-token=""
kubectl create secret generic redis-pass --from-literal=redis-pass=""

minikube dashboard
minikube tunnel


