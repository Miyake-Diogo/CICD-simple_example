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
1. Create a K8S cluster `k3d cluster create {clustername}` (It is in your preference host, in this case are local with K3D).
2. To list all contexts `kubectl config get-contexts`.
3. Enter in K8S cluster's context `kubectl cluster-info --context {clustername}`.
4. See if nodes are ok in cluster with `kubectl get nodes`.
5. Create a app that you need, in this case a simple webserver in go.
6. Make a Dockerfile using multistage build.
7. Build image and send to registry if is necessary (in this case I will maintain local) `docker build -t {myname}/{imagename}:{tag} .`.
8. Run image with command `docker run --rm -p 9090:9090`.
9. Insert in line 16 of file `k8s/deployment.yaml`your image address and in line 22 the port of application.
10. To run deployment in your kubernetes cluster use the command: `kubectl apply -f k8s/deployment.yaml`.
11. To see if deployment is running use `kubectl get pods`.
12. 
13. 
99. To delete cluster `k3d cluster delete {clustername}`.

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

# Reference
Materials used to reference:

[Deploy contínuo com GitOps e ArgoCD](https://www.youtube.com/watch?v=63HGUgQXD1w)
[k3d Docs](https://k3d.io/v5.0.1/)
[ArgoCD](https://github.com/argoproj/argo-cd)
[ArgoCD Get Started](https://argo-cd.readthedocs.io/en/stable/)
