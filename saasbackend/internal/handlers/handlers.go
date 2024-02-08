package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"saasteamtest/saasbackend/domain"
	"saasteamtest/saasbackend/models"

	"github.com/go-chi/chi"
)

// send a whole YED shipment to this endpoint, return the vendor
func CreateProduct(productService domain.ProductServiceInterface) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return BadDataError{Msg: fmt.Errorf("readAll: %w", err).Error()}
		}

		var product models.Product
		err = json.Unmarshal(body, &product)
		if err != nil {
			return BadDataError{Msg: fmt.Errorf("unmarshal: %w", err).Error()}
		}

		savedProduct, err := productService.Save(product)

		if err != nil {
			return fmt.Errorf("save: %w", err)
		}
		return RespondOK(w, savedProduct)
	}
}

func GetProductById(productService domain.ProductServiceInterface) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		productId := chi.URLParam(r, "product_id")

		product, err := productService.GetProductById(productId)
		if err != nil {
			return NotFoundError{Msg: fmt.Errorf("get product by id: %w", err).Error()}
		}
		return RespondOK(w, product)
	}
}

func GetAllProducts(productService domain.ProductServiceInterface) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		products, err := productService.GetAllProducts()
		if err != nil {
			return fmt.Errorf("get all products: %w", err)
		}

		productsResponse := models.ExternalProductResponse{
			Products: products,
			Count:    int64(len(products)),
		}

		return RespondOK(w, productsResponse)
	}
}

func CalculatePrice(productService domain.ProductServiceInterface) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return BadDataError{Msg: fmt.Errorf("readAll: %w", err).Error()}
		}

		var cartItems []*models.CartItem
		err = json.Unmarshal(body, &cartItems)
		if err != nil {
			return BadDataError{Msg: fmt.Errorf("unmarshal: %w", err).Error()}
		}

		res, err2 := productService.CalculatePrice(cartItems)
		if err2 != nil {
			return fmt.Errorf("compute cost: %w", err2)
		}
		return RespondOK(w, res)
	}
}
