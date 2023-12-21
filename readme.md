# Learning Minikube, Kubernetes, Docker
---

This repo is a collection of notes and projects I have found across the web as I learn kubernetes with minikube. 

I am looking for small footprint projects which explain to me the ways to get different ways of building kubernetes clusters from the tutorials linked.

I am coming from a PHP webserver background, but also would like to get Golang, python, bash, rust and other languages in the mix.


1st - goHtmlServer
---
https://itnext.io/how-to-experiment-locally-on-kubernetes-with-minikube-and-your-local-dockerfiles-48833fcd90c9

- first run, COULD NOT GET TO THE IP ADDRESS!




2nd - nginx
---
https://medium.com/@mngaonkar/kubernetes-get-started-deploy-a-simple-web-server-9636f4aa8706

kubectl create -f nginx.yaml





3rd - Jeff Geerling - Kubernetes 101
---
https://github.com/geerlingguy/kubernetes-101/tree/master



MiniKube Cheatsheet

Command | Description / Notes
------------ | -------------
minikube start | --driver=virtualbox
eval $(minikube -p minikube docker-env) | uses the local docker in minikube
minikube halt
minikube status
minikube version 
minikube ip | get the ip address
minikube dashboard | open a dashboard in the browser

