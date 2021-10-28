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
9. Push image in your DockerHub account (You can use any registry that you wish), first make login in your account using  `docker login`, second run this command: `docker push {myname}/{imagename}:{tag}`.
10. Insert in line 16 of file `k8s/deployment.yaml`your image address and in line 22 the port of application.
11. To run deployment in your kubernetes cluster use the command: `kubectl apply -f k8s/deployment.yaml`.
12. To see if deployment is running use `kubectl get pods`, if you not add -n and namespace name kubernetes show only for default namespace.
13. Make a service in `k8s/service.yaml` with data about deployments, addin a port and name for your app, in my case i use goaap.
14. Run command `kubectl apply -f k8s/service.yaml`.
15. To show services you can use command `kubectl get svc`.
16. To test if service are running use command: `kubectl port-forward svc/goapp 9090:9090`, after this you can access in your localhost:9090 at webrowser.
Nice you have a application running in kubernetes.  

# Kustomize
In real world we have a necessity to use a versioning of images, to use it with more easy way, kustomize are tool that help you to make it.

1. Make a `kustomization.yaml` in `k8s` folder.
2. Add parameters that you need in kustomize file.
3. To build and update file you can use command into k8s folder: `kustomize build`

# Github Actions
1. Catch a token in dockerhub and add in github repository
2. Make a push and see in git actions your CD running.

For Now Update a file and add a Deploy stage in workflow yaml. 



99. To delete cluster `k3d cluster delete {clustername}`.

# About Dockerfile
Little explain about dockerfile for generating go image using multistage build.  

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

[FullCycle - Deploy contínuo com GitOps e ArgoCD](https://www.youtube.com/watch?v=63HGUgQXD1w)
[FullCycle - Reference repo](https://github.com/codeedu/argocd-fullcycle)

[k3d Docs](https://k3d.io/v5.0.1/)
[Kustomize](https://kustomize.io/)
[ArgoCD](https://github.com/argoproj/argo-cd)
[ArgoCD Get Started](https://argo-cd.readthedocs.io/en/stable/)
