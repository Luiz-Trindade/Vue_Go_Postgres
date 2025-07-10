package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Price float64 `json:"price"`
}

var db *gorm.DB

func main() {
	dsn := "host=postgres user=postgres password=postgres dbname=crud port=5432 sslmode=disable"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&Product{})

	// Rotas
	r := gin.Default()
	r.GET("/products", listProducts)
	r.POST("/products", createProduct)
	/*
	r.PUT("/products/:id", updateProduct)
	r.DELETE("/products/:id", deleteProduct)
	*/

	r.Run(":8080") // Servidor na porta 8080
}

// Handler: Listar produtos
func listProducts(c *gin.Context) {
	var products []Product
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

// Handler: Criar produto
func createProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&product)
	c.JSON(http.StatusOK, product)
}