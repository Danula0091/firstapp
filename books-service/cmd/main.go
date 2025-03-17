package main

import (
	"diary-api/pkg/books"
	"diary-api/pkg/common/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()
	handler := db.Init(dbUrl)

	books.RegisterRoutes(router, handler)

	router.Run(port)
}
