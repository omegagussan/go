//+build after

package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"saasteamtest/saasbackend/data"
	"saasteamtest/saasbackend/domain"

	"github.com/go-chi/chi"
	"github.com/google/go-cmp/cmp"
)

var createProductId = ""

func SecondSetUp() {
	productHandler := data.NewMemoryStore()
	productService = domain.NewProductService(productHandler)
}

func MapJSONBodyIsEqualStringWithoutId(t *testing.T, responsString string, myPB string) {
	testBytes := []byte(myPB)
	responseBytes := []byte(responsString)

	testMap := map[string]interface{}{}
	err := json.Unmarshal(testBytes, &testMap)
	if err != nil {
		t.Error(err)
	}
	responseMap := map[string]interface{}{}
	err = json.Unmarshal(responseBytes, &responseMap)

	productId, ok := responseMap["product_id"].(string)
	if !ok {
		t.Error("product_id is not a string")
	}
	createProductId = productId	
	delete(responseMap, "product_id")

	if err != nil {
		t.Error(err)
	}
	// compare the json ojbects without respect to order a second time with a different function
	if !cmp.Equal(testMap, responseMap) {
		diffstring := cmp.Diff(testMap, responseMap)
		t.Errorf("\n...cmp Diff string  = %v", diffstring)
		t.Errorf("\n...cmp Equal false expected = %v\n...obtained = %v", testMap, responseMap)
	}
}

func TestCreateProductMemoryStore(t *testing.T) {
	//this is a hack relying on the ordering of tests. 
	SecondSetUp()

	newProduct := make(map[string]interface{})
	newProduct["product_name"] = "volleyball"
	newProduct["product_price"] = 750
	newProduct["product_type"] = "sporting_good"
	newProduct["product_discount_price"] = 525
	newProduct["coupon_code"] = "sport30"

	body, err := json.Marshal(newProduct)
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/products", bytes.NewReader(body))

	r := chi.NewRouter()
	r.Method("POST", "/products", BaseHandler(CreateProduct(productService)))
	r.ServeHTTP(rec, req)

	expectedResult, err := json.Marshal(newProduct)
	if err != nil {
		t.Error(err)
	}

	MapJSONBodyIsEqualStringWithoutId(t, rec.Body.String(), string(expectedResult))
}

func TestGetProductByIdMemoryStore(t *testing.T) {

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/" + createProductId, nil)

	r := chi.NewRouter()
	r.Method("GET", "/products/{product_id}", BaseHandler(GetProductById(productService)))
	r.ServeHTTP(rec, req)

	myProduct := make(map[string]interface{})
	myProduct["product_id"] = createProductId
	myProduct["product_name"] = "volleyball"
	myProduct["product_price"] = 750
	myProduct["product_type"] = "sporting_good"

	expectedResult, err := json.Marshal(myProduct)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestGetAllProductsMemoryStore(t *testing.T) {

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)

	r := chi.NewRouter()
	r.Method("GET", "/products", BaseHandler(GetAllProducts(productService)))
	r.ServeHTTP(rec, req)

	myProduct := make(map[string]interface{})
	myProduct["product_id"] = createProductId
	myProduct["product_name"] = "volleyball"
	myProduct["product_price"] = 750
	myProduct["product_type"] = "sporting_good"

	//TODO: one could test many products here

	var productSlice []map[string]interface{}
	productSlice = append(productSlice, myProduct)

	myProductResult := make(map[string]interface{})
	myProductResult["count"] = 1
	myProductResult["products"] = productSlice

	expectedResult, err := json.Marshal(myProductResult)
	if err != nil {
		t.Error(err)
	}
	

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}