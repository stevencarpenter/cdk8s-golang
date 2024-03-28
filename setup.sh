#!/usr/bin/env zsh

brew install minikube
brew install helm

## Create kind cluster from config
#kind create cluster --config kind/config.yaml

minikube start
minikube addons enable dashboard
minikube addons enable metrics-server

helm install ingress-nginx ingress-nginx/ingress-nginx
