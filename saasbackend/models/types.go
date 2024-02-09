package models

type ProductsResponse struct {
	Count    int64      `json:"count"`
	Products []*Product `json:"products"`
}

type ExternalProductResponse struct {
	Count    int64      		`json:"count"`
	Products []*ExternalProduct `json:"products"`
}

type CartCostResponse struct {
	TotalObjects    int64		`json:"total_objects"`
	TotalCost 		int64		`json:"total_cost"`
}

type Product struct {
	ProductId            string `json:"product_id"`
	ProductName          string `json:"product_name"`
	ProductPrice         int64  `json:"product_price"`
	ProductDiscountPrice int64  `json:"product_discount_price"`
	CouponCode           string `json:"coupon_code"`
	ProductType          string `json:"product_type"`
}

type CartItem struct {
	ProductId            string 	`json:"product_id"`
	Quantity         	 int64  	`json:"quantity"`
	CouponCode           *string 	`json:"coupon_code,omitempty"` //optional
}

type Cart struct {
	CartItems []*CartItem `json:"cart"`	
}

type ExternalProduct struct {
	ProductId            string `json:"product_id"`
	ProductName          string `json:"product_name"`
	ProductPrice         int64  `json:"product_price"`
	ProductType          string `json:"product_type"`
}
