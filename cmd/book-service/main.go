package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/heise/myproject/Desktop/firstapp/internal/book/handler"
	"github.com/heise/myproject/Desktop/firstapp/internal/book/repository"
	"github.com/heise/myproject/Desktop/firstapp/internal/book/routes"
	"github.com/heise/myproject/Desktop/firstapp/internal/book/service"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("path-to-file-env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()
	db, err := repository.InitDB(dbUrl)
	if err != nil {
		log.Fatal("Failed to connect to db. Error: ", err)
	}
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)
	routes.RegisterBookRoutes(router, bookHandler)
	router.Run(port)
}
