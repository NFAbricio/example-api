package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		envs.DatabaseUser,
		envs.DatabasePassword,
		envs.DatabaseHost,
		envs.DatabasePort,
		envs.DatabaseName,
	)

	migrations(dburl)
	
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

	router.Run(":45457")
}

func migrations(url string) {

	m, err := migrate.New("file://migrations", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	if err := m.Up(); err != nil {
		if fmt.Sprintf("%s", err) != "no change" {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}
	}


}