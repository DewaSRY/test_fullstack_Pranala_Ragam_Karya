package database

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type CreateProdductRequest struct {
	Name  string  `json:"name" validate:"required"`
	Price float32 `json:"price" validate:"required"`
	Stock int32   `json:"stock" validate:"required"`
}

type ProductResponse struct {
	ID    int32   `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Stocs int32   `json:"stock"`
}

type ProductServices interface {
	CreateProduct(ctx context.Context, db *gorm.DB, args CreateProdductRequest) (*ProductResponse, error)
	// GetByIDProduct(ctx context.Context, db *gorm.DB, id float32) (*ProductResponse, error)
	// UpdateProduct(ctx context.Context, db *gorm.DB, id float32, args CreateProdductRequest) (*ProductResponse, error)
	// DeleteProduct(ctx context.Context, db *gorm.DB, id float32) (*ProductResponse, error)
	// GetAllProduct(ctx context.Context, db *gorm.DB) []ProductResponse
}

// implementation here

type productServicesImpl struct{}

// CreateProduct implements [ProductServices].
func (s productServicesImpl) CreateProduct(ctx context.Context, db *gorm.DB, args CreateProdductRequest) (*ProductResponse, error) {

	new_product_model := ProductModel{
		Name:  args.Name,
		Price: args.Price,
		Stock: args.Stock,
	}
	if result := db.Create(&new_product_model); result.Error != nil {
		response := ProductResponse{
			Name:  args.Name,
			Price: args.Price,
			Stocs: args.Stock,
			ID:    new_product_model.ID,
		}

		return &response, nil
	}

	return nil, errors.New("failed to create product")
}

// // DeleteProduct implements [ProductServices].
// func (s productServicesImpl) DeleteProduct(ctx context.Context, db *gorm.DB, id float32) (*ProductResponse, error) {
// 	panic("unimplemented")
// }

// // GetAllProduct implements [ProductServices].
// func (s productServicesImpl) GetAllProduct(ctx context.Context, db *gorm.DB) []ProductResponse {
// 	panic("unimplemented")
// }

// // GetByIDProduct implements [ProductServices].
// func (s productServicesImpl) GetByIDProduct(ctx context.Context, db *gorm.DB, id float32) (*ProductResponse, error) {
// 	panic("unimplemented")
// }

// // UpdateProduct implements [ProductServices].
// func (s productServicesImpl) UpdateProduct(ctx context.Context, db *gorm.DB, id float32, args CreateProdductRequest) (*ProductResponse, error) {
// 	panic("unimplemented")
// }

func NewProductServices() ProductServices {
	return productServicesImpl{}
}
