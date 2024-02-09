package data

import (
	"errors"
	"saasteamtest/saasbackend/models"
	"github.com/google/uuid"
)

type MemoryStore struct {}

func NewMemoryStore() *MemoryStore {
	memoryStore := MemoryStore{}
	return &memoryStore
}

//variable at package level
var product_store = make(map[string]models.Product)

func (h *MemoryStore) Create(obj models.Product) (*models.Product, error) {
	obj.ProductId = uuid.New().String()
	product_store[obj.ProductId] = obj
	return &obj, nil
}

func (h *MemoryStore) ReadOne(id string) (*models.Product, error) {
	val, ok := product_store[id]
	if !ok {
		return nil, errors.New("no such product found")
	}
	return &val, nil
}

func (h *MemoryStore) Read() ([]*models.Product, error) {
	var items []*models.Product

	for _, product := range product_store {
		items = append(items, &product)
	}
	return items, nil
}
