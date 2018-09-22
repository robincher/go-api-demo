# GO API Demo

Demo REST API built with Golang

Before you start , make sure you have installed and configure your Go binary correctly.

## Tech Stack
* POSTMAN or Restlet for Testing
* Go 1.11
* Go Libraries
    - mux : Request Router and Dispatcher
    - mgo : MongoDB Driver
    - toml : Parse mongodb server configuration files

## Walkthrough

### Starting the REST API Server
```
go run main.go
```

### Building an executabler
```
go build -o ./bin/output.exe .

./bin/output.exe (or double click on the exe file)
```

Execute your HTTP request @ http://localhost:8000 

## Reference
+ [Codementor Tutorial](https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo)
+ [Hackernoon Tutorial](https://hackernoon.com/build-restful-api-in-go-and-mongodb-5e7f2ec4be94)
+ [Importing in GO](https://scene-si.org/2018/01/25/go-tips-and-tricks-almost-everything-about-imports/)
