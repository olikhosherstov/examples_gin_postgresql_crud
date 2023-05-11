package main

import (
	"github.com/olikhosherstov/examples_gin_postgresql_crud/pkg/books"
	"github.com/olikhosherstov/examples_gin_postgresql_crud/pkg/common/db"

	"github.com/gin-gonic/gin"
)

func main() {
	dbUrl := "postgres://goapimedium:wsadqe123@localhost:5432/go_medium_api"

	r := gin.Default()
	h := db.Init(dbUrl)

	books.RegisterRoutes(r, h)

	r.Run()
}
