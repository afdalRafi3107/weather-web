package main

import (
	"backend/handlers"
	"backend/utils"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	utils.InitConfig()

	port := viper.GetString("PORT")

	if port == ""{
		port="8000"
	}

	r:= gin.Default()

	cfg := cors.DefaultConfig()
	cfg.AllowOrigins = []string{"http://localhost:5173"}
	cfg.AllowMethods= []string{"GET", "POST","OPTIONS"}
	cfg.AllowHeaders= []string{"Content-Type", "Authorization"}
	r.Use(cors.New(cfg))

	api := r.Group("/api")
	{
		api.GET("/health", handlers.GEtHealt)
		api.GET("/weather", handlers.GetWeather)
	}

	fmt.Println("Server is Running on http://localhost:" + port)
	if err := r.Run(":" + port); err !=nil{
		panic(err)
	}
}