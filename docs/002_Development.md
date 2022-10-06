# Development

## Setup

### Environment variables

Verify that you have the .env file in your project, you should normally find this in the keybase of Nightborn for the given project.

## Running

```
go run .
```

## Adding a resource

Firstly you should always verify if your resource should be added in the database, if it's the case you should start by adding the Model in the database folder.

When you've added this resource, you should go in the database.go file and add to the autoMigrate your new Model.

Once added, you should verify if you need to create a new repository, if it's the case create the new repository and add the default functions

-   GetById
-   GetByPage
-   Add
-   Modify
-   Delete

Once those functions written, you can start by adding the new Repository in the repository container (repository.go)

And now comes the fun part, you can create the Usecase, this one is the one that should be unit-tested. First create the function you had to create (no need to create the basic ones).

-   In the usecase you should validate that incoming data is correctly formatted and contains the correct data

And directly write the unit test that verifies that your code runs correctly.

Once this is done verify that all tests still run with this command

```
go test ./...
```

Once this is done and all tests work, you can go in the controller part and create the corresponding functions, in this part you translate the GIN request into the inhouse one.

When all of this is done, you add it to the router and you should be done.

If you require the user to be authenticated, add the

```
    middlewares.Authorize(...) function to your router, you have to pass in the function that has to be called afterwards and/or the role that the user should have to call your function
```

when running the debugger and/or running the go run . command, you should see your route available.