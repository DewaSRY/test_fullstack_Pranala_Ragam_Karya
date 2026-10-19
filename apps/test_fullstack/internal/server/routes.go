package server

import (
	"DewaSRY/test_fullstack/internal/database"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)

	// Product CRUD routes
	products := r.Group("/products")
	{
		products.POST("", s.CreateProduct)
		products.GET("", s.GetAllProducts)
		products.GET("/:id", s.GetProductByID)
		products.PUT("/:id", s.UpdateProduct)
		products.DELETE("/:id", s.DeleteProduct)
	}

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

// Parse and validate a product ID from the URL.
func parseProductID(c *gin.Context) (int32, bool) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid product ID",
		})
		return 0, false
	}

	return int32(id), true
}

// Return 404 for missing records and 500 for unexpected errors.
func respondProductError(c *gin.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "product not found",
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "internal server error",
	})
}

// CREATE PRODUCT
func (s *Server) CreateProduct(c *gin.Context) {
	ctx := c.Request.Context()

	var req database.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	product, err := s.productServices.CreateProduct(
		ctx,
		s.db.GetInstance(),
		req,
	)
	if err != nil {
		respondProductError(c, err)
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GET ALL PRODUCTS
func (s *Server) GetAllProducts(c *gin.Context) {
	ctx := c.Request.Context()

	products, err := s.productServices.GetAllProduct(
		ctx,
		s.db.GetInstance(),
	)
	if err != nil {
		respondProductError(c, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

// GET PRODUCT BY ID
func (s *Server) GetProductByID(c *gin.Context) {
	id, ok := parseProductID(c)
	if !ok {
		return
	}

	product, err := s.productServices.GetByIDProduct(
		c.Request.Context(),
		s.db.GetInstance(),
		id,
	)
	if err != nil {
		respondProductError(c, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

// UPDATE PRODUCT
func (s *Server) UpdateProduct(c *gin.Context) {
	id, ok := parseProductID(c)
	if !ok {
		return
	}

	var req database.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	product, err := s.productServices.UpdateProduct(
		c.Request.Context(),
		s.db.GetInstance(),
		id,
		req,
	)
	if err != nil {
		respondProductError(c, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

// DELETE PRODUCT
func (s *Server) DeleteProduct(c *gin.Context) {
	id, ok := parseProductID(c)
	if !ok {
		return
	}

	product, err := s.productServices.DeleteProduct(
		c.Request.Context(),
		s.db.GetInstance(),
		id,
	)
	if err != nil {
		respondProductError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product deleted successfully",
		"product": product,
	})
}
