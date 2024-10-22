package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id          string `gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name        string
	Description string
	Price       float32
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.Id = uuid.NewString()
	return nil
}
