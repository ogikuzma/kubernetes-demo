kubectl delete all --all
kubectl delete kubegres consumer-db

cd kubernetes

kubectl apply -f ingress.yml

cd consumer
kubectl apply -f consumer-db-configmap.yml
kubectl apply -f consumer-db-secret.yml
kubectl apply -f consumer-db.yml
kubectl apply -f consumer-deployment.yml