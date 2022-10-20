# Structure

## Structure of the project

-   app: contains the app
-   build: contains the Dockerfile in order to build the project
-   config: contains the structure of the configuration of the project
-   docs: contains all the docs related to the project
-   .env: contains all the environment variables of the project
-   .gitignore
-   doc.go: contains all the meta for the swagger
-   main.go: contains the initialisation of the application

## Structure of the app

-   constants: contains all the constants
-   contracts: contains all the DTO's (DATA TRANSFER OBJECTS)
-   controllers: contains the mapping between gin & our business layer
-   database: contains all the models and the database connection
-   gateways: contains all the gateways to communicate with external systems
-   middlewares: contains all the middlewares used by the router to connect to the controllers
-   repositories: contains all the repositories to connect the business layer to the data layer
-   routers: contains all the routers (REST, ...) to communicate with the controllers
-   usesases: contains all the usecases (business layer)

## Layers

### Constants

This layer contains all the constants for the application, this includes all the error messages sent back to the callers.

### Contracts

This layer contains all the DTO's of the application, a base DTO is always provided if the returned type is a resource of our application.

### Controllers

This layer contains the mapping between a router request (GIN) and our business layer (Usecase). This will read from the request the different parameters, queries, bodies, headers, ... in order to verify if the request
should/can be translated to our business layer. This is commonly used to verify the integrity of the structure.

### Database

This layer contains all the database related structures, the models are 1-1 mappings between the SQL (underlying layer) and the GO structures. We use GORM as our ORM.
Important note : If no database connection string is given a in-memory version of sqlite is used.

### Middlewares

This layer contains all the middlewares that are applied before a requests enters the Controllers layer.

Examples would be the AuthMiddleware that verifies that a bearer_token has been given and that this one contains a still non-expired user. The Sub is than extracted and passed to the request.

### Repositories

This layer contains all the interactions between your business logic and your data layer, every repository should at least have those functions :

-   GetById
-   GetByPage
-   Add
-   Modify
-   Delete

### Gateways

This layer contains all the interactions between your business logic and external services.

### Routers

This layer contains the definition of the routes for the REST API.

### Usecases

This layer contains all the business logic of your application.

## Minimal Layers Structure

```
├── app
│   ├── constants
│   │   ├── ...
│   │   ├── constants.go
│   │   └── errors.go
│   ├── contracts
│   │   ├── ...
│   │   └── DTO.go
│   ├── controllers
│   │   ├── ...
│   │   └── controller.go
│   ├── database
│   │   ├── models
│   │   │   ├── ...
│   │   │   └── model.go
│   │   └── database.go
│   ├── middlewares
│   │   ├── ...
│   │   └── auth.go
│   ├── repositories
│   │   ├── ...
│   │   └── repository.go
│   ├── gateways
│   │   ├── ...
│   │   └── gateway.go
│   ├── routers
│   │   └── router.go
│   └── usecases
│       ├── ...
│       └── usecase.go
├── build
│   └── Dockerfile
├── config
│   └── config.go
├── docs
│   ├── 000_Structure.md
│   └── 001_Installation.md
│   └── ...
├── .env
├── .gitignore
├── doc.go
├── doc.go
├── go.mod
├── main.go
└── README.md
```