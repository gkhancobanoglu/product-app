package request

import "product-app/service/model"

// AddProductRequest, bir ürün ekleme isteği için kullanılan yapıdır.
// JSON formatında gönderilen veriler bu yapıya map edilir.
type AddProductRequest struct {
	Name     string  `json:"name"`     // Ürünün adı
	Price    float32 `json:"price"`    // Ürünün fiyatı
	Discount float32 `json:"discount"` // Ürünün indirim oranı
	Store    string  `json:"store"`    // Ürünün mağazası
}

// ToModel, AddProductRequest yapısını service katmanında kullanılan
// ProductCreate modeline dönüştürmek için kullanılan fonksiyondur.
func (addProductRequest AddProductRequest) ToModel() model.ProductCreate {
	return model.ProductCreate{
		Name:     addProductRequest.Name,
		Price:    addProductRequest.Price,
		Discount: addProductRequest.Discount,
		Store:    addProductRequest.Store,
	}
}
