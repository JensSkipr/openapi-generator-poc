# Welcome

## Introduction

You need :

```
NPX installed
GO installed
NPM installed
NODE installed
```

To run the example :

In 1 terminal do :
```
go run .
```

In another do :

```
chmod +x ./scripts/e2e.sh
./scripts/e2e.sh
```

This will start the go back-end and run tests on it via a testID, 
you'll have a new database with an unique guid available to verify the state of the database


To work on the E2E : 

```
cd e2e
npm i && npm run generate-api
```

When you're done coding you can go back to the root of the project and run the 

```
./scripts/e2e.sh
```