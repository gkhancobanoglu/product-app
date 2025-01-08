package service

import (
	"errors"
	"product-app/domain"
	"product-app/persistence"
)

type FakeProductRepository struct {
	products []domain.Product
}

func NewFakeProductRepository(initialProducts []domain.Product) persistence.IProductRepository {
	return &FakeProductRepository{
		products: initialProducts,
	}
}

func (fakeRepository *FakeProductRepository) GetAllProducts() []domain.Product {
	return fakeRepository.products
}

func (fakeRepository *FakeProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	// Belirtilen mağazaya ait ürünleri döndüren fonksiyon
	var filteredProducts []domain.Product
	for _, product := range fakeRepository.products {
		if product.Store == storeName {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

func (fakeRepository *FakeProductRepository) AddProduct(product domain.Product) error {
	// Yeni bir ürünü ürün listesine ekler
	fakeRepository.products = append(fakeRepository.products, domain.Product{
		Id:       int64(len(fakeRepository.products)) + 1,
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	})
	return nil
}

func (fakeRepository *FakeProductRepository) GetById(productId int64) (domain.Product, error) {
	// Belirtilen ID'ye sahip ürünü döndürür
	for _, product := range fakeRepository.products {
		if product.Id == productId {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("Ürün bulunamadı")
}

func (fakeRepository *FakeProductRepository) DeleteById(productId int64) error {
	// Belirtilen ID'ye sahip ürünü siler
	for index, product := range fakeRepository.products {
		if product.Id == productId {
			// Ürünü listeden kaldır
			fakeRepository.products = append(fakeRepository.products[:index], fakeRepository.products[index+1:]...)
			return nil
		}
	}
	return errors.New("Ürün bulunamadı")
}

func (fakeRepository *FakeProductRepository) UpdatePrice(productId int64, newPrice float32) error {
	// Belirtilen ID'ye sahip ürünün fiyatını günceller
	for i, product := range fakeRepository.products {
		if product.Id == productId {
			fakeRepository.products[i].Price = newPrice
			return nil
		}
	}
	return errors.New("Ürün bulunamadı")
}
