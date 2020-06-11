## Sample Golang API
Sample API code for Go.

## File structure
```commandline
./
├── app
│  ├── controller
│  │  ├── receiver.go
│  │  └── sender.go
│  ├── go.mod
│  ├── go.sum
│  ├── infrastructure
│  │  └── database.go
│  ├── interface
│  │  └── user.go
│  └── main.go
├── docker
│  ├── app
│  │  └── Dockerfile-app
│  └── db
│     └── Dockerfile-db
└── docker-compose.yml
```

## Source Summary
### Base Language
* Golang v.1.14.3
    * [Document](https://golang.org/doc/)

### Database
* Redis v6.0.3
    * [Document](https://redis.io/documentation)

### Hot Reload
* [cosmtrek/air](ttps://github.com/cosmtrek/air)

### ORM
* [gomodule/redigo](https://github.com/gomodule/redigo)