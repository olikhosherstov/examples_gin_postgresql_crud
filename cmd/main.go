package main

import (
	"github.com/olikhosherstov/examples_gin_postgresql_crud/pkg/books"
	"github.com/olikhosherstov/examples_gin_postgresql_crud/pkg/common/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("../pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	books.RegisterRoutes(r, h)

	r.Run(port)
}
