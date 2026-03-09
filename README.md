# Prerequisites

## Requirements
- Docker

# Run locally
## Environment variables
Create a `.env` file in the project root:
```
DBUSER=root
DBPASS=
DBURL=mysql:3306
DBNAME=ecom

JWT_SECRET=
```

## Start the application
```bash
docker compose up --build
```

# Run application on AWS

## Uppload image to docker hub 
```shell
docker buildx build \
  --platform linux/amd64 \
  -t <docker username>/e-commerce:latest \
  --push .
```

## Copy files
Copy `.env` and `docker-compose.prod.yaml` file from the local project to the aws instances root directory.

## Pull to aws
```shell
docker compose -f docker-compose.prod.yaml pull
docker compose -f docker-compose.prod.yaml up -d
```

# Connect to mysql server in docker-compose
```shell
docker exec -it <mysql container name> bash
mysql -u root -p
# pwd 😘
use ecom
```
# Note on goose
If goose migration fails from FK dependencies, reset data in tables. Dont do in prod tho o.o
## From mysql:
```sql
SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS cart_items, carts, order_items, orders, product_images, products, categories, users;
SET FOREIGN_KEY_CHECKS = 1;
```
## From root dir:
```shell
docker exec -it <gin container name> -dir migrations reset
docker exec -it <gin container name> -dir migrations up
```