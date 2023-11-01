# Upera interview (go-task)

## Ports:

- **mongodb**:  `27017`
- **redis**:  `6379`
- **product**:  `5000`
- **revision**:  `5001`

Please ensure that these ports are available.

## Getting Started:

1. Clone the repository: `git clone https://github.com/Bright98/upera-interview.git`
2. Change to the project directory: `cd upera-interview`

project contains the `docker-compose.yaml` file, which is used to start the services.
To run the application, execute the following commands:

- Start MongoDB: `docker compose up -d mongodb`
- Start Redis: `docker compose up -d redis`
- Build and start the product service: `docker compose up --build -d product`
- Build and start the revision service: `docker compose up --build -d revision`

After building, the names of the containers will begin with `upera-interview-<container>`

## Architecture and Folder Structure:

This project service consists of two microservices: `product` and `revision`.

Each microservice contains various files and directories such as **repository**, **domain**, **received**, **Dockerfile**, ...

Below, I explain each of these folders:

### Repository

The repository folder is the most internal part of the service and is responsible for handling database and cache connections.
In this project, `MongoDB` is used as the main database, and `Redis` is used for caching and message handling.

- For MongoDB, the mongo-driver package was used. [Link to mongo-driver package](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo)
- For Redis, the redis/v9 package was used. [Link to redis/v9 package](https://pkg.go.dev/github.com/redis/go-redis/v9)

### Domain

The domain folder follows the principles of **Domain-Driven Design** and constitutes the second layer in our architecture.
It contains the core functions and interacts with the "repository" folder, utilizing its functions.

**Models** are created in this folder.

### Received

The received folder is the most external layer, responsible for handling APIs and incoming messages.
It can utilize functions from the "domain" to send requests and receive responses.

**REST** is used for API handling, and **Redis** is employed for messaging.
- For REST, the Gin framework is used. [Link to Gin framework](https://pkg.go.dev/github.com/gin-gonic/gin)

### Initial and Main Files

The initial file contains an `init` function that handles the following steps:
1. Loading environments variables
2. Checking if directories are correctly connected
3. Establishing database and cache connections
4. Starting the Gin server on the service port

The main file handles the Gin methods and routes and subscribes to messages on separate threads, with each **thread** handling a specific channel.
### Dockerfile

Each service has its Dockerfile for creating its Docker image.

### .env

This file is used during the development process, but after deployment, environment variables are sourced from the `docker-compose.yaml` file.

## Models

#### 1. Products
```go
type Products struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Color         string `json:"color"`
	Price         int64  `json:"price"`
	ImageUrl      string `json:"image_url"`
	CreatedAt     int64  `json:"created_at"`
	LastUpdatedAt int64  `json:"last_updated_at"`
	Status        string `json:"status"` // for handling status of product: active, removed
}
```
#### 2. Product attributes
```go
type ProductAttributes struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Price       int64  `json:"price"`
	ImageUrl    string `json:"image_url"`
}
```
#### 3. Revision
```go
type Revisions struct {
	ID                string    `json:"id"`
	ProductID         string    `json:"Product_id"`
	RevisionNo        int       `json:"revision_no"`
	UpdatedAttributes []string  `json:"updated_attributes"`
	PreviousProduct   *Products `json:"previous_product"`
	NewProduct        *Products `json:"new_product"`
	UpdatedAt         int64     `json:"updated_at"`
}
```
#### 4. Revision for send
```go
type Revisions struct {
	ProductID         string    `json:"Product_id"`
	UpdatedAttributes []string  `json:"updated_attributes"`
	PreviousProduct   *Products `json:"previous_product"`
	NewProduct        *Products `json:"new_product"`
}
```
#### 5. Errors
```go
type Errors struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

```
#### 6. API response

```json
{
  "data": {},
  "error": {},
  "message": ""
}
```

#### 7. routes
```text
/api/<service-name>/<entity>/id/<:entity-id>
```

- In the models, the `status` field is used to manage the status of each record. Two status types are considered:
  - **active**: for newly created records
  - **removed**: for removal operations (although this part, as you didn't mention, not implemented)
- Unix timestamps are used for `updated_at` and `created_at` attributes.
- For unique ID, the Google uuid used [Link to Google uuid package](https://pkg.go.dev/github.com/google/uuid)

## Service Descriptions

### Product Service

Provides 4 routes to fulfill the required functionalities:
- insert product:
  - method: `POST`
  - get product model from request
  - generate unique id
  - set created_at and last_updated_at to the current unix timestamp
  - set status = active
  - create revision model
  - send the revision model to the revision service via messaging
  - return inserted product id in the response

- update product:
  - method: `PUT`
  - get product attribute from request
  - get last updated product
  - calculate the differences between the last updated attributes and the new attributes
  - set last_updated_at to current unix timestamp
  - update product in database
  - create revision model
  - send the revision model to the revision service via messaging

- get details of specific product (last version):
  - method: `GET`
  - get `product-id` from route to get the details.
  - return the product json in response

- get all products (with pagination):
  - method: `GET`
  - get `skip` and `limit` from query params
  - return a list of products in response

### Revision Service

- get all revision of one product:
  - method: `GET`
  - extract `product-id` from the route to retrieve product revisions
  - get `skip` and `limit` from query parameters
  - return a list of revisions in response

- get specific revision of one product
  - method: `GET`
  - extract `product-id` from route.
  - extract `revision_no` from route.
  - find specific revision by product ID and revision number.
  - return revision.new_product in response.

## API Documentation
> I attached the **postman** collection on this GitHub repository.