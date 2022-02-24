kubectl delete all --all
kubectl delete kubegres consumer-db
kubectl delete kubegres order-db
kubectl delete -f https://raw.githubusercontent.com/reactive-tech/kubegres/v1.15/kubegres.yaml

cd kubernetes

kubectl apply -f ingress.yml
kubectl apply -f https://raw.githubusercontent.com/reactive-tech/kubegres/v1.15/kubegres.yaml

cd consumer
kubectl apply -f consumer-db-configmap.yml
kubectl apply -f consumer-db-secret.yml
kubectl apply -f consumer-db.yml
kubectl apply -f consumer-deployment.yml

cd ../order
kubectl apply -f order-db-configmap.yml
kubectl apply -f order-db-secret.yml
kubectl apply -f order-db.yml
kubectl apply -f order-deployment.yml
