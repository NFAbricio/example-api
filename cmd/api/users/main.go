package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/NFAbricio/example-api/config"
	"github.com/NFAbricio/example-api/internal/controller"
	"github.com/NFAbricio/example-api/internal/payments/stripe"
	"github.com/NFAbricio/example-api/users"
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
		envs.DatabasePort,
		envs.DatabaseUser,
		envs.DatabaseName,
		envs.DatabasePassword,
	)
	
	db, err := gorm.Open("postgres", dburl)
	if err != nil {
		log.Fatalf("gorm open error: %v", err)
	}

	sqlDB := db.DB()
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("ping erro: %v", err)
	}
	log.Println("Connected to database")

	paymentService := stripe.NewPaymentService(envs.StripKey)
	
	reposiotry := store.NewRepository(db)
	svc := users.NewService(reposiotry, paymentService)
	controller := controller.NewUserController(svc)
	
	router := gin.Default()

	router.Group("/api/v1")

	router.POST("/user", controller.Create)

	router.Run()
}