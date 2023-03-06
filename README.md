# Restaurant managment

## This project contains three microservices for managing products, ratings, and workers. Each microservice provides CRUD operations through a RESTful API.

# Products Microservice

## The products microservice manages product data. It provides the following endpoints:

- GET /products - retrieves all products
- PUT /products/{id} - updates a product by ID
- POST /products - creates a new product
- DELETE /products/{id} - deletes a product by ID

The microservice is implemented in the products directory. The cmd directory contains the main application code, and the internal directory contains the implementation details.

# Ratings Microservice

## The ratings microservice manages product rating data. It provides the following endpoints:

- GET /ratings - retrieves all ratings
- POST /ratings - creates a new rating
- DELETE /ratings/{id} - deletes a rating by ID

The microservice is implemented in the ratings directory. The cmd directory contains the main application code, and the internal directory contains the implementation details.

# Workers Microservice

## The workers microservice manages worker data. It provides the following endpoints:

- GET /workers - retrieves all workers
- PUT /workers/{id} - updates a worker by ID
- POST /workers - creates a new worker
- DELETE /workers/{id} - deletes a worker by ID

The microservice is implemented in the workers directory. The cmd directory contains the main application code, and the internal directory contains the implementation details.


# Orders Microservice

## The orders microservice manages orders data. It provides the following endpoints:

- GET /orders - retrieves all orders
- POST /orders - creates a new order
- DELETE /orders/{id} - deletes a order by ID

The microservice is implemented in the orders directory. The cmd directory contains the main application code, and the internal directory contains the implementation details.

# Running the Microservices
## To run each microservice, navigate to its directory and run the following commands:

```Make
make run
```
This will build and run the microservice.

# Running Tests

## To run the tests for each microservice, navigate to its directory and run the following command:

```Make
make test
```
This will run the tests for the microservice.

# License

## This project is licensed under the MIT License. See the LICENSE file for details.
