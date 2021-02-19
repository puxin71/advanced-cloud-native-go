# github reference
[github repo -Gin](https://github.com/JacobSNGoodwin/memrizr)

# Using CURL to ad-hoc test the local server
* Note that after we deploy multiple replica of this service to the ks8 cluster and 
* after configuring the ingress/loadbalancer for the service pod set, we can drop 
* the port from the URL

# Launch HTTP web server
PORT=3000 ./Gin-Web

## ping
curl -v http://localhost:3000/v1/ping

## get all books
curl -v http://localhost:3000/v1/books

## add two new books and each has its own ISBN number
curl -v \
 -X POST \
 -H "Content-Type:application/json" \
 -d '{"title":"GOOGLE INC","isbn":"2345UUx90"}' \
 http://localhost:3000/v1/books

curl -v \
 -X POST \
 -H "Content-Type:application/json" \
 -d '{"title":"TOY STORY","isbn":"2345UUx91"}' \
 http://localhost:3000/v1/books

## get a book by ISBN
curl -v \
 -X GET \
 http://localhost:3000/v1/books/2345UUx90

## get a slice of books by a list of ISBNs
curl -v 'http://localhost:3000/v1/books?isbns=2345UUx90&isbns=2345UUx91'

# Personal notes
[git](https://alvinalexander.com/git/)

## remove files from the previous commit
git rm -r -f file

## push to github repo
```
git remote add origin https://github.com/puxin71/advanced-cloud-native-go.git
git branch -M main
git push -u origin main
```

# build the smallest go docker image
[blog to build small go service image](https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324)
[github repo](https://github.com/chemidy/smallest-secured-golang-docker-image)

## remove the temp build docker image and other unused resources
* Note that the multi-step build can leave dangling images which are shown as the 'none' tagged images
`docker system prune` or `docker image prune`
`docker build -t gin-web:1.0.1 --rm .` and `docker image prune`

## use docker-compose to build, run, and stop the dockers
`docker-compose build && docker image prune`

# Kubernetes

## ks8 cluster
kubectl cluster-info
kubectl apply -f kubernetes/
kubectl get deployment
kubectl get pods -o wide
kubectl scale deployment gin-web replicas=1
kubectl scale deployment gin-web replicas=0
kubectl delete deployment gin-web

## delete deployment, delete all pods and remove all docker containers
kubectl delete deployment gin-web

## install Nginx ingress ks8 controller for our service

* Run ks8 nginx deployment to set up the Nginx Ingress Controller. Now, we can create Ingress resources in our Kubernetes cluster and route external requests to our services.
* [ks8 ingress with Nginx example](https://matthewpalmer.net/kubernetes-app-developer/articles/kubernetes-ingress-guide-nginx-example.html)
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/cloud/deploy.yaml
kubectl -n ingress-nginx get pods
```

* Using Nginx ingress, 
$ curl -v  -X POST  -H "Content-Type:application/json"  -d '{"title":"GOOGLE INC","isbn":"2345UUx90"}'  http://localhost/v1/books
$ curl -kL http://localhost/v1/books
[{"title":"GOOGLE INC","author":"","isbn":"2345UUx90"}]

# ks8 deployed dockers
$ docker ps
CONTAINER ID   IMAGE                                 COMMAND                  CREATED          STATUS          PORTS                               NAMES
e241bff4365a   bfbbf260995e                          "/go/bin/Gin-Web"        8 minutes ago    Up 8 minutes                                        k8s_gin-web_gin-web-7bcbb77d79-xtxmp_ingress-nginx_842aaa62-5293-4849-a44e-169271817a3d_0 
2f9dcb0c5181   bfbbf260995e                          "/go/bin/Gin-Web"        8 minutes ago    Up 8 minutes                                        k8s_gin-web_gin-web-7bcbb77d79-tvt6z_ingress-nginx_6f90e3d7-b4c2-4e78-944b-3e8ef6597c11_0 
cf9c089a22a7   k8s.gcr.io/ingress-nginx/controller   "/usr/bin/dumb-init …"   19 minutes ago   Up 19 minutes                                       k8s_controller_ingress-nginx-controller-7fc74cf778-vbz9k_ingress-nginx_dc259637-6dfe-4b08-

# delete dockers 
kubectl -n ingress-nginx get deployment
kubectl -n ingress-nginx delete deployment gin-web ingress-nginx-controller
kubectl -n ingress-nginx delete service gin-web ingress-nginx-controller ingress-nginx-controller-admission
kubectl -n ingress-nginx delete ingress gin-web

# manage ks8 resources using custom namespace
Kubernetes Namespace would be the perfect options for you. You can easily create namespace resource.

kubectl create -f custom-namespace.yaml
```
$  apiVersion: v1
    kind: Namespace
    metadata:
      name:custom-namespace
```
Now you can deploy all of the other resources(Deployment,ReplicaSet,Services etc) in that custom namespaces.

If you want to delete all of these resources, you just need to delete custom namespace. by deleting custom namespace, all of the other resources would be deleted. Without it, ReplicaSet might create new pods when existing pods are deleted.
```
kubectl delete all --all -n {namespace}
kubectl delete all --all -n ingress-nginx
kubectl get all -n ingress-nginx
 No resources found in ingress-nginx namespace.
```