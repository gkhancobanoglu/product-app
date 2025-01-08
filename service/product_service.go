package service

import (
	"errors"
	"product-app/domain"
	"product-app/persistence"
	"product-app/service/model"
)

// IProductService, ürünlerle ilgili servis işlemleri için bir arayüzdür.
type IProductService interface {
	Add(productCreate model.ProductCreate) error
	DeleteById(productId int64) error
	GetById(productId int64) (domain.Product, error)
	UpdatePrice(productId int64, newPrice float32) error
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
}

// ProductService, IProductService arayüzünü uygulayan yapı olup,
// ürünlerin eklenmesi, silinmesi ve alınması gibi işlemleri gerçekleştirir.
type ProductService struct {
	productRepository persistence.IProductRepository
}

// Yeni bir ProductService oluşturur ve gerekli repository'i alır.
func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

// Yeni bir ürün ekler.
// Ürün eklemeden önce doğrulama yapılır.
func (productService *ProductService) Add(productCreate model.ProductCreate) error {
	validateErr := validateProductCreate(productCreate)
	if validateErr != nil {
		// Eğer doğrulama hatası varsa, hata döndürülür.
		return validateErr
	}
	// Ürün veritabanına eklenir.
	return productService.productRepository.AddProduct(domain.Product{
		Name:     productCreate.Name,
		Price:    productCreate.Price,
		Discount: productCreate.Discount,
		Store:    productCreate.Store,
	})
}

// Belirli bir ID'ye sahip ürünü siler.
func (productService *ProductService) DeleteById(productId int64) error {
	return productService.productRepository.DeleteById(productId)
}

// Belirli bir ID'ye sahip ürünü getirir.
func (productService *ProductService) GetById(productId int64) (domain.Product, error) {
	return productService.productRepository.GetById(productId)
}

// Ürünün fiyatını günceller.
func (productService *ProductService) UpdatePrice(productId int64, newPrice float32) error {
	return productService.productRepository.UpdatePrice(productId, newPrice)
}

// Tüm ürünleri getirir.
func (productService *ProductService) GetAllProducts() []domain.Product {
	return productService.productRepository.GetAllProducts()
}

// Belirli bir mağazaya ait tüm ürünleri getirir.
func (productService *ProductService) GetAllProductsByStore(storeName string) []domain.Product {
	return productService.productRepository.GetAllProductsByStore(storeName)
}

// Ürün ekleme işlemi için doğrulama yapılır.
// İndirim oranının %70'ten fazla olmasına izin verilmez.
func validateProductCreate(productCreate model.ProductCreate) error {
	if productCreate.Discount > 70.0 {
		return errors.New("Discount can not be greater than 70")
	}
	return nil
}
