echo -e "deploying consumer-service..."
cd consumer-service
docker build -t [URL_TO_YOUR_IMAGE] .
docker push [URL_TO_YOUR_IMAGE]

echo -e "deploying order-service..."
cd ../order-service
docker build -t [URL_TO_YOUR_IMAGE] .
docker push [URL_TO_YOUR_IMAGE]