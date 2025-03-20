package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./configs/envs/.env")
	viper.ReadInConfig()
	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)
	router := gin.Default()
	handler := repo.Init(dbUrl)
	controllers.RegisterRoutes(router, handler)
	router.Run(port)

}
