package data

import (
	"saasteamtest/saasbackend/models"
)

type ProductHandle interface{
	Create(obj models.Product) (*models.Product, error)
	ReadOne(id string) (*models.Product, error)
	Read() ([]*models.Product, error)
}