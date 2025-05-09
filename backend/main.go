package main

import (
	"library-app/handlers"
	"library-app/migrations"
	"library-app/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=db user=postgres password=postgres dbname=411221174_apriantinuban port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	err = migrations.MigrateEnums(db)
	if err != nil {
		log.Fatal("failed to migrate enums:", err)
	}

	models.Migrate(db)

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	handlers.RegisterRoutes(r)

	r.Run(":8080")
}
