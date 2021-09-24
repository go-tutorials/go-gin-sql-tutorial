# go-sql-rest-api

## How to run
#### Clone the repository
```shell
git clone https://github.com/go-tutorials/go-gin-sql-rest-api.git
cd go-gin-sql-rest-api
```

#### To run the application
```shell
go run main.go
```

## API Design
### Common HTTP methods
- GET: retrieve a representation of the resource
- POST: create a new resource
- PUT: update the resource
- PATCH: perform a partial update of a resource
- DELETE: delete a resource

## API design for health check
To check if the service is available.
#### *Request:* GET /health
#### *Response:*
```json
{
    "status": "UP",
    "details": {
        "sql": {
            "status": "UP"
        }
    }
}
```


## API design for users
#### *Resource:* users

### Get all users
#### *Request:* GET /users
#### *Response:*
```json
[
    {
        "id": "spiderman",
        "username": "peter.parker",
        "email": "peter.parker@gmail.com",
        "phone": "0987654321",
        "dateOfBirth": "1962-08-25T16:59:59.999Z"
    },
    {
        "id": "wolverine",
        "username": "james.howlett",
        "email": "james.howlett@gmail.com",
        "phone": "0987654321",
        "dateOfBirth": "1974-11-16T16:59:59.999Z"
    }
]
```

### Get one user by id
#### *Request:* GET /users/:id
```shell
GET /users/wolverine
```
#### *Response:*
```json
{
    "id": "wolverine",
    "username": "james.howlett",
    "email": "james.howlett@gmail.com",
    "phone": "0987654321",
    "dateOfBirth": "1974-11-16T16:59:59.999Z"
}
```

### Create a new user
#### *Request:* POST /users 
```json
{
    "id": "wolverine",
    "username": "james.howlett",
    "email": "james.howlett@gmail.com",
    "phone": "0987654321",
    "dateOfBirth": "1974-11-16T16:59:59.999Z"
}
```
#### *Response:* 1: success, 0: duplicate key, -1: error
```json
1
```

### Update one user by id
#### *Request:* PUT /users/:id
```shell
PUT /users/wolverine
```
```json
{
    "username": "james.howlett",
    "email": "james.howlett@gmail.com",
    "phone": "0987654321",
    "dateOfBirth": "1974-11-16T16:59:59.999Z"
}
```
#### *Response:* 1: success, 0: not found, -1: error
```json
1
```

### Delete a new user by id
#### *Request:* DELETE /users/:id
```shell
DELETE /users/wolverine
```
#### *Response:* 1: success, 0: not found, -1: error
```json
1
```

## Common libraries
- [core-go/health](https://github.com/core-go/health): include HealthHandler, HealthChecker, SqlHealthChecker
- [core-go/config](https://github.com/core-go/config): to load the config file, and merge with other environments (SIT, UAT, ENV)
- [core-go/log](https://github.com/core-go/log): log and log middleware

### core-go/config
To load the config from "config.yml", in "configs" folder
```go
server:
  name: go-postgresql-gin-rest-api
  port: 8080

sql:
  driver: postgres
  data_source_name: postgres://postgres:admin@123456@localhost:5432/demogo?sslmode=disable
  # postgres://username:password@host:5432/database?sslmode=disable
log:
  level: info
  duration: duration
  map:
    time: "@timestamp"
    msg: message

middleware:
  log: true
  skips: /health
  request: body
  response: response
  size: size
```

### core-go/log *&* core-go/middleware
```go
package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/core-go/config"
	sv "github.com/core-go/service"
	"github.com/gin-gonic/gin"
	"net/http"

	"go-service/internal/app"
)

func main() {
	var conf app.Root
	er1 := config.Load(&conf, "configs/config")
	if er1 != nil {
		panic(er1)
	}

	g := gin.New()

	g.Use(gin.Logger())

	g.Use(gin.Recovery())

	g.Use(ginBodyLogMiddleware())

	er2 := app.Route(g , context.Background(), conf)
	if er2 != nil {
		panic(er2)
	}

	fmt.Println(sv.ServerInfo(conf.Server))
	if er3 := http.ListenAndServe(sv.Addr(conf.Server.Port), g); er3 != nil {
		fmt.Println(er3.Error())
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func ginBodyLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString("\n"), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		fmt.Println("Response body: " + blw.body.String())
	}
}
```
To configure to ignore the health check, use "skips":
```yaml
middleware:
  skips: /health
```
