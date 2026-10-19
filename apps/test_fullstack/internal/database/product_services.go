package database

import (
	"context"

	"gorm.io/gorm"
)

type CreateProductRequest struct {
	Name  string  `json:"name" validate:"required"`
	Price float32 `json:"price" validate:"gt=0"`
	Stock int32   `json:"stock" validate:"gt=0"`
}

type ProductResponse struct {
	ID    int32   `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Stock int32   `json:"stock"`
}

type ProductServices interface {
	CreateProduct(ctx context.Context, db *gorm.DB, args CreateProductRequest) (*ProductResponse, error)
	GetByIDProduct(ctx context.Context, db *gorm.DB, id int32) (*ProductResponse, error)
	UpdateProduct(ctx context.Context, db *gorm.DB, id int32, args CreateProductRequest) (*ProductResponse, error)
	DeleteProduct(ctx context.Context, db *gorm.DB, id int32) (*ProductResponse, error)
	GetAllProduct(ctx context.Context, db *gorm.DB) ([]ProductResponse, error)
}

type productServicesImpl struct{}

func NewProductServices() ProductServices {
	return productServicesImpl{}
}

// Helper to convert ProductModel into ProductResponse.
func toProductResponse(product ProductModel) *ProductResponse {
	return &ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}
}

// CreateProduct creates a new product.
func (s productServicesImpl) CreateProduct(
	ctx context.Context,
	db *gorm.DB,
	args CreateProductRequest,
) (*ProductResponse, error) {

	product := ProductModel{
		Name:  args.Name,
		Price: args.Price,
		Stock: args.Stock,
	}

	result := db.WithContext(ctx).Create(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return toProductResponse(product), nil
}

// GetByIDProduct retrieves a product by ID.
func (s productServicesImpl) GetByIDProduct(
	ctx context.Context,
	db *gorm.DB,
	id int32,
) (*ProductResponse, error) {

	var product ProductModel

	result := db.WithContext(ctx).First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return toProductResponse(product), nil
}

// UpdateProduct updates an existing product.
func (s productServicesImpl) UpdateProduct(
	ctx context.Context,
	db *gorm.DB,
	id int32,
	args CreateProductRequest,
) (*ProductResponse, error) {

	var product ProductModel

	result := db.WithContext(ctx).First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}

	product.Name = args.Name
	product.Price = args.Price
	product.Stock = args.Stock

	result = db.WithContext(ctx).Save(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return toProductResponse(product), nil
}

// DeleteProduct deletes a product by ID.
func (s productServicesImpl) DeleteProduct(
	ctx context.Context,
	db *gorm.DB,
	id int32,
) (*ProductResponse, error) {

	var product ProductModel

	result := db.WithContext(ctx).First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}

	result = db.WithContext(ctx).Delete(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return toProductResponse(product), nil
}

// GetAllProduct retrieves all products.
func (s productServicesImpl) GetAllProduct(
	ctx context.Context,
	db *gorm.DB,
) ([]ProductResponse, error) {

	var products []ProductModel

	result := db.WithContext(ctx).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	responses := make([]ProductResponse, 0, len(products))

	for _, product := range products {
		responses = append(responses, ProductResponse{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
			Stock: product.Stock,
		})
	}

	return responses, nil
}
