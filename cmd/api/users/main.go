package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/NFAbricio/example-api/config"
	"github.com/NFAbricio/example-api/users/store"
)

func main() {
	envs, err := config.LoadEnvs()
	if err != nil {
		panic("failed to load envs")
	}

	dburl := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		envs.DatabaseHost,
		envs.DatabaseName,
		envs.DatabasePort,
		envs.DatabaseUser,
		envs.DatabasePassword,
	)
	
	db, err := gorm.Open("postgres", dburl)
	
	repository := store.NewRepository(db)
	
	router := gin.Default()

	router.Group("/api/v1")

	router.POST("/user")
	router.GET("/user/:id")
	router.PUT("/user/:id")
	router.DELETE("/user/:id")
	router.POST("/user/login")

	router.Run()
}