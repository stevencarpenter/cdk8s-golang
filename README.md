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

# Install some helm charts and the nginx controller
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm install ingress-nginx ingress-nginx/ingress-nginx

# Create the secrets in k8s 
kubectl create secret generic pubg-api-token --from-literal=pubg-api-token=""
kubectl create secret generic redis-pass --from-literal=redis-pass=""
```

### cdk8s
From the root of this project, synthesis and then apply the charts
```shell
cdk8s synth
kubectl apply -f dist/redis.k8s.yaml
kubectl apply -f dist/pubg.k8s.yaml  
kubectl apply -f dist/pubgserver.k8s.yaml  
```