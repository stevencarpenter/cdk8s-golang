#!/usr/bin/env zsh

brew install minikube
brew install helm

minikube start
minikube addons enable dashboard
minikube addons enable metrics-server

helm install ingress-nginx ingress-nginx/ingress-nginx


minikube dashboard
minikube tunnel


