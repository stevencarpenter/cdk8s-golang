# cdk8s-golang

Infrastructure as code in Go for the Angi takehome project.

This deploys minikube with nginx ingress controller, a redis cluster, a basic webserver, and a cron job to hydrate redis from the PUBG api.

## Setup

### Local K8s with Ingress
```shell
#!/usr/bin/env zsh

brew install minikube
brew install helm

minikube start
minikube addons enable dashboard
minikube addons enable metrics-server

helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm install ingress-nginx ingress-nginx/ingress-nginx
```

### cdk8s


##