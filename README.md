# Microservices application written in Go
This is a simple application written it Go to help understand how to build microservices. It uses HTTP for communication between client and the  API Gateway, and gRPC for communication between services.
![alt text](/diagram.drawio.png)


## Create DB
```shell
psql postgres
CREATE DATABASE auth_svc;
CREATE DATABASE order_svc;
CREATE DATABASE product_svc;
\l
\q
```