run:
	cd api-gateway && start /b make server
	cd auth-svc && start /b make server
	cd order-svc && start /b make server
	cd product-svc && start /b  make server
