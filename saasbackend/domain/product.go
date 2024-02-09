package domain

import (
	"fmt"
	"saasteamtest/saasbackend/models"
)

type ProductHandler interface {
	Create(models.Product) (*models.Product, error)
	ReadOne(string) (*models.Product, error)
	Read() ([]*models.Product, error)
}

type ProductServiceInterface interface {
	CalculatePrice([]*models.CartItem) (models.CartCostResponse, error)
	GetAllProducts() ([]*models.ExternalProduct, error)
	GetProductById(string) (*models.ExternalProduct, error)
	GetProductByIdInternal(string) (*models.Product, error)
	Save(models.Product) (*models.Product, error)
}

type ProductService struct {
	productHandler ProductHandler
}

func NewProductService(p1 ProductHandler) ProductServiceInterface {
	return ProductService{
		productHandler: p1,
	}
}

func (ps ProductService) CalculatePrice(cartItems []*models.CartItem) (models.CartCostResponse, error) {
	// Calculate the cost of the products
	var cost int64 = 0
	var count int64 = 0
	for _, item := range cartItems {
		p, err := ps.GetProductByIdInternal(item.ProductId)
		if err != nil {
			continue
		}
		var effectiveCost int64 = 0
		if item.CouponCode != nil && p.CouponCode == *item.CouponCode {
			effectiveCost = p.ProductDiscountPrice
		} else {
			effectiveCost = p.ProductPrice
		}

		cost = cost + (effectiveCost * item.Quantity)
		count = count + item.Quantity
	}
	r := models.CartCostResponse{
		TotalObjects: count,
		TotalCost: cost,
	}
	return r, nil
}

func Map[T, V any](ts []T, fn func(T) V) []V {
    result := make([]V, len(ts))
    for i, t := range ts {
        result[i] = fn(t)
    }
    return result
}

func (ps ProductService) GetAllProducts() ([]*models.ExternalProduct, error) {
	myProducts, err := ps.productHandler.Read()
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}
	externalProducts := Map(myProducts, func(p *models.Product) *models.ExternalProduct { 
		return &models.ExternalProduct{
			ProductId:            p.ProductId,
			ProductName:          p.ProductName,
			ProductPrice:         p.ProductPrice,
			ProductType:          p.ProductType,
		} 
	})
	return externalProducts, nil
}

func (ps ProductService) GetProductById(productId string) (*models.ExternalProduct, error) {
	myProduct, err := ps.productHandler.ReadOne(productId)
	if err != nil {
		return nil, fmt.Errorf("read one: %w", err)
	}
	externalProduct := models.ExternalProduct{
		ProductId:            myProduct.ProductId,
		ProductName:          myProduct.ProductName,
		ProductPrice:         myProduct.ProductPrice,
		ProductType:          myProduct.ProductType,
	}
	return &externalProduct, nil
}

func (ps ProductService) GetProductByIdInternal(productId string) (*models.Product, error) {
	myProduct, err := ps.productHandler.ReadOne(productId)
	if err != nil {
		return nil, fmt.Errorf("read one: %w", err)
	}
	return myProduct, nil
}

func (ps ProductService) Save(product models.Product) (*models.Product, error) {
	myProduct := models.Product{
		ProductName:          product.ProductName,
		ProductPrice:         product.ProductPrice,
		ProductDiscountPrice: product.ProductDiscountPrice,
		CouponCode:           product.CouponCode,
		ProductType:          product.ProductType,
	}
	savedProduct, err := ps.productHandler.Create(myProduct)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	return savedProduct, nil
}
