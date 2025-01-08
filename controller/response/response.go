package response

import "product-app/domain"

// ErrorResponse struct, hata mesajlarının dönmesi için kullanılır.
type ErrorResponse struct {
	ErrorDescription string `json:"errorDescription"` // Hata açıklamasını tutar
}

// ProductResponse struct, ürün verilerini dışa aktarmak için kullanılır.
type ProductResponse struct {
	Name     string  `json:"name"`     // Ürünün ismi
	Price    float32 `json:"price"`    // Ürünün fiyatı
	Discount float32 `json:"discount"` // Ürüne uygulanmış indirim oranı
	Store    string  `json:"store"`    // Ürünün satıldığı mağaza adı
}

// ToResponse fonksiyonu, domain.Product tipindeki bir ürünü ProductResponse'a dönüştürür.
func ToResponse(product domain.Product) ProductResponse {
	return ProductResponse{
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	}
}

// ToResponseList fonksiyonu, domain.Product listesinde bulunan her bir ürünü ProductResponse listesine dönüştürür.
func ToResponseList(products []domain.Product) []ProductResponse {
	var productResponseList = []ProductResponse{} // Boş bir response listesi oluşturulur
	for _, product := range products {            // Tüm ürünler üzerinde dönülür
		productResponseList = append(productResponseList, ToResponse(product)) // Ürün her seferinde response'a dönüştürülüp listeye eklenir
	}
	return productResponseList // Dönüştürülmüş ürün listesi geri döndürülür
}
