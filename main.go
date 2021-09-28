package main

import (
	"fmt"
	"go-gin-sql-rest-api/internal/app"
	. "go-gin-sql-rest-api/internal/model"
)

func main() {
	err := ConnectDB()

	if err != nil {
		fmt.Println("Status:", &err)
	}

	r := app.SetUpRouter()

	r.Run("localhost:8080")

}
