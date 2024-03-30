# cdk8s-golang

Infrastructure as code in Go for the Angi takehome project.

This deploys minikube with nginx ingress controller, a redis cluster, a basic webserver, and a cron job to hydrate redis from the PUBG API.

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
kubectl create secret generic pubg-api-token --from-literal=pubg-api-token="<YOUR_TOKEN>"
kubectl create secret generic redis-pass --from-literal=redis-pass="secret" 

minikube tunnel
```

### cdk8s
From the root of this project, synthesis and then apply the charts
```shell
cdk8s synth
kubectl apply -f pubgsa.k8s.yaml #I couldn't get the Service Account working in cdk8s so I created it manually
kubectl apply -f dist/redis.k8s.yaml
kubectl apply -f dist/pubg.k8s.yaml  
kubectl apply -f dist/pubgserver.k8s.yaml  
```

## pubg
This is a cron job that runs the code from the pubg repo. It shoves leaderboard data into redis daily. It can be run manually with `kubectl create job --from=cronjob.batch/pubg  pubg-manual-001` in case the cache needs to be hydrated right away instead of waitng for the daily run.

## pubgserver
This is the light webserver exposing one endpoint `/pubg/leaderboard` that takes one query parameter `accountId` and returns that player's stats from redis.

### Usage
```shell
curl -X GET "localhost:8090/pubg/leaderboard?accountId=account.28b08053492a44659f8bf0517d8c3580"
```

Response:
```json
{"games":"205","rank":"258","wins":"50"}
```
