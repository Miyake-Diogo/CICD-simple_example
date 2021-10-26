# CICD-simple_example
Simple example using Git actions + Argo CD + K8S + Docker and  GO lang

# Intro


# Pre reqs
* Have an ArgoCD account and Installed. 
* Docker installed.
* Have k3D intalled.
* Kubectl installed.
* Have a Github account.
* Unlimited of curiosity to learn new things.

# Steps
1. Create a K8S cluster `k3d cluster create {clustername}` (It is in your preference host, in this case are local with K3D)
2. To list all contexts `kubectl config get-contexts`.
2. Enter in K8S cluster's context `kubectl cluster-info --context {clustername}`
3. See if nodes are ok in cluster with `kubectl get nodes`
4. Create a app that you need, in this case a simple webserver in go.
5. Make a Dockerfile using multistage build.
6. Build image and send to registry if is necessary (in this case I will maintain local) `docker build -t {myname}/{imagename}:{tag} .`.
7. Run image with command `docker run --rm -p 9090:9090`.
99. To delete cluster `k3d cluster delete {clustername}`

# About Dockerfile
Explain about dockerfile for generating go image

```Dockerfile
# Multistage build
# First image compile a binary in GO
FROM golang:1.17 as build 

WORKDIR /app
COPY . . 
RUN CGO_ENABLED=0 go build -o server main.go
# Second image catch a binary from first and generate a simple and light image to use in webserver

FROM alpine:3.12
WORKDIR /app
COPY --from=build /app/server .
CMD ["./server"]
```