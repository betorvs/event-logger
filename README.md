# event-logger

![Go Test](https://github.com/betorvs/event-logger/workflows/Go%20Test/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/betorvs/event-logger/badge.svg?branch=main)](https://coveralls.io/github/betorvs/event-logger?branch=main)

This is a simple event to log web service to register any kind of events like deployments, alerts, code analysis and whatever you believe that needs to be registered. 

## Environment variables

```sh
export PORT=9090

export APP_NAME=event-logger

export LOG_LEVEL=INFO
```

## Dependency Management
The project is using [Go Modules](https://blog.golang.org/using-go-modules) for dependency management
Module: github.com/betorvs/event-logger

## Test and coverage

Run the tests

```sh 
TESTRUN=true go test ./... -coverprofile=cover.out

go tool cover -html=cover.out
```

Install [golangci-lint](https://github.com/golangci/golangci-lint#install) and run lint:

```sh
golangci-lint run
```

## Docker Build

```sh
docker build .
```

## Send one test

```sh
#!/usr/bin/env bash

curl -v -X POST -H "Content-Type: application/json" -d '{
    "name": "test1",
    "kind": "test",
    "labels": {
        "testa": "testa1"
    },
    "message": "test1 event logging",
    "source": "curl command",
    "status": 2
}' \
  http://localhost:9090/event-logger/v1/event
```

## References

### Golang Spell
The project was initialized using [Golang Spell](https://github.com/danilovalente/golangspell).

### Architectural Model
The Architectural Model adopted to structure the application is based on The Clean Architecture.
Further details can be found here: [The Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html) and in the Clean Architecture Book.
