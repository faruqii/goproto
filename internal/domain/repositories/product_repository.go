package repositories

import (
	"github.com/faruqii/goproto/internal/domain/entities"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *entities.Product) (err error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) CreateProduct(product *entities.Product) (err error) {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}

	return nil
}
