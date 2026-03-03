# Connect to mysql server in docker-compose
```shell
docker exec -it ecommerce_db bash
mysql -u root -p
# pwd 😘
use ecom
```

# If goose migration fails from FK dependencies, drop all tables
## From mysql:
```sql
SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS cart_items, carts, order_items, orders, product_images, products, categories, users;
SET FOREIGN_KEY_CHECKS = 1;
```
## From root dir:
```shell
goose reset
goose up
```