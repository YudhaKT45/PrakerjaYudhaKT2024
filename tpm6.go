// Yudha Kusuma Triatmaja
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	engine := gin.New()
	dsn := "host=localhost user=postgres password=Zfxonfire45 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	productRepo := &ProductRepo{DB: db}

	productRepo.Migrate()

	productHdl := &ProductHdl{Repository: productRepo}
	productGroup := engine.Group("/products")
	{
		// Get all user
		productGroup.GET("", productHdl.GetGorm)
		// Create user
		productGroup.POST("", productHdl.CreateGorm)
		// Update user
		productGroup.PUT("/:id", productHdl.UpdateGorm)
		// Update user
		productGroup.DELETE("/:id", productHdl.DeleteGorm)
	}

	err = engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
