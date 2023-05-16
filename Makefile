run-all:
	cd api-gateway && start /b make server
	cd auth-svc && start /b make server
	cd order-svc && start /b make server
	cd product-svc && start /b  make server
view-ports:
	netstat -ano | findstr :3000
	netstat -ano | findstr :50051
	netstat -ano | findstr :50052
	netstat -ano | findstr :50053
#replace process id number with results from make ports command
stop:
	taskkill /PID 13988 /F
	taskkill /PID 9160 /F
	taskkill /PID 8868 /F
	taskkill /PID 6864 /F
docker-build:
	cd api-gateway && docker build -t amalmadhu06/ecom-api-gateway .
	cd auth-svc && docker build -t amalmadhu06/ecom-auth-svc .
	cd order-svc && docker build -t amalmadhu06/ecom-order-svc .
	cd product-svc && docker build -t amalmadhu06/ecom-product-svc .

docker-push:
	cd api-gateway && docker push amalmadhu06/ecom-api-gateway
	cd auth-svc && docker push amalmadhu06/ecom-auth-svc
	cd order-svc && docker push amalmadhu06/ecom-order-svc
	cd product-svc && docker push amalmadhu06/ecom-product-svc
run:
	sudo docker compose up
