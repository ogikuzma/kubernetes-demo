# Kubernetes application with Kubegres operator

A goal of this application is to deploy Kubernetes cluster with a cluster of PostgreSQL databases managed by Kubegres operator and demonstrate their features. The application consists of two microservices developed in Go programming language with the purpose of enabling consumers to order food from a restaurant.

## Application architecture

The first microservice is consumer-service which is responsible for creating consumers and verify their existence. This microservice consists of 4 endpoints:

- GET /hello - used for verifing pod running status
- GET /crash - it crashes the pod 1 second after it is called with the purspose to show Kubernetes reaction to unexpected failure scenarios
- POST / - it accepts consumer data in JSON format and creates a new consumer
- GET /verify/:consumerId - used for verifing existence of user with id "consumerId" from the endpoint 

The second microservice is order-service which is responsible for accepting orders and updating their status. This microservice consists of 4 endpoints too:

- GET /hello - used for verifing pod running status
- GET /crash - it crashes the pod 1 second after it is called with the purspose to show Kubernetes reaction to unexpected failure scenarios
- POST / - it accepts order data in JSON format and creates a new order

# Kubernetes cluster
Both of these two microservices should be uploaded inside 2 pods wrapped into load balacing services. In front of the Kubernetes cluster there is an Ingress component with the purpose to route traffic to services. PostgreSQL databases are used as a persistence layer and they are managed by the Kubegres operator. This operator was picked because it is easy-to-use and configure. It is responsible for creating a cluster of one master node and two replicas which synhronize theirself with the master node in order to achieve consistent data across all database instances. Kubegres operator is also responsible for handling database instances failure. The complete Kubernetes architecture could be seen in the picture below:

![diagram](https://user-images.githubusercontent.com/57645292/155567194-424c15a8-c634-4b06-b488-4785a37f4858.png)

## Installation
1. Start a Kubernetes cluster with the minikube tool. Docker could also be used as a driver
```
minikube start --nodes=3 --driver=hyperv
```

2. Enable Ingress controller
```
minikube addons enable ingress
```
3. Get service URLs and use Ingress URL and port for accessing the application
```
minikube service list
```
4. Open Kubernetes dashboard in order to have a graphic view on what is going on in the cluster
```
minikube dashboard --url
```
5. Go to deploy.sh script and change [URL_TO_YOUR_IMAGE] to define the URL where you would like to deploy consumer and order images
6. Go to kubernetes/consumer/consumer-deployment.yaml and again change the [URL_TO_YOUR_IMAGE] with your URL
7. Do the same thing for kubernetes/order/order-deployment.yaml 
8. Run restart.sh script in order to start the whole cluster
9. That's it

## Features

1. Scale number of pods for consumer and order microservices.
```
kubectl scale –-replicas=[number-of-replicas] –f [yml-file]
```
2. Install HorizontalPodAutoscaler component (see official Kubernetes website for installation [2]) in order to enable dynamic pod scaling depending on resource status. Then, run this command to keep all of the pods on 50% CPU usage.
```
kubectl autoscale deployment [deployment] --cpu-percent=50 --min=1 --max=10
```
3. Send HTTP request to /crash endpoing and watch Kubernetes dashboard in order to observe Kubernetes behaviour to pod failure. As long as there is still one pod active, application will continue working.
4. Exec one of the replicas in order to see if data you write to the master instance, with the POST / requests, is propaged to the replicas. Kubegres operator creates PersistentVolume and PersistentVolumeClaims in order to keep the data if one of instance fail or you stop the whhole cluster.
```
kubectl exec consumer-db-2-0 -it -- bash

psql –U postgres
\c
SELECT * FROM consumers;
```
5. Change image for one of the microservices and observe how Kubernetes gradually does transition to the new version.
```
kubectl set image deployments/consumer-app consumer-app=[new-image]
```

# Useful links
1. https://kubernetes.io/docs/home/
2. https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/
3. https://www.kubegres.io/
