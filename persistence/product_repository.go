package persistence

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"product-app/domain"
	"product-app/persistence/common"
)

// IProductRepository, ürünlerle ilgili CRUD işlemlerini tanımlayan arayüzdür.
type IProductRepository interface {
	GetAllProducts() []domain.Product                        // Tüm ürünleri getirir.
	GetAllProductsByStore(storeName string) []domain.Product // Belirli bir mağazaya ait ürünleri getirir.
	AddProduct(product domain.Product) error                 // Yeni bir ürün ekler.
	GetById(productId int64) (domain.Product, error)         // Belirli bir ID'ye sahip ürünü getirir.
	DeleteById(productId int64) error                        // Belirli bir ID'ye sahip ürünü siler.
	UpdatePrice(productId int64, newPrice float32) error     // Ürünün fiyatını günceller.
}

// ProductRepository, IProductRepository arayüzünü uygulayan yapıdır.
type ProductRepository struct {
	dbPool *pgxpool.Pool // PostgreSQL bağlantı havuzunu temsil eder.
}

// NewProductRepository, yeni bir ProductRepository örneği oluşturur.
func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{
		dbPool: dbPool,
	}
}

// GetAllProducts, tüm ürünleri veritabanından getirir.
func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	ctx := context.Background()
	productRows, err := productRepository.dbPool.Query(ctx, "Select * from products")

	if err != nil {
		log.Error("Tüm ürünler alınırken hata oluştu: %v", err)
		return []domain.Product{}
	}
	return extractProductsFromRows(productRows)
}

// GetAllProductsByStore, belirli bir mağazaya ait ürünleri getirir.
func (productRepository *ProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	ctx := context.Background()

	getProductsByStoreNameSql := `Select * from products where store = $1`

	productRows, err := productRepository.dbPool.Query(ctx, getProductsByStoreNameSql, storeName)

	if err != nil {
		log.Error("Belirli bir mağazanın ürünleri alınırken hata oluştu: %v", err)
		return []domain.Product{}
	}
	return extractProductsFromRows(productRows)
}

// AddProduct, yeni bir ürünü veritabanına ekler.
func (productRepository *ProductRepository) AddProduct(product domain.Product) error {
	ctx := context.Background()

	insert_sql := `Insert into products (name,price,discount,store) VALUES ($1,$2,$3,$4)`

	addNewProduct, err := productRepository.dbPool.Exec(ctx, insert_sql, product.Name, product.Price, product.Discount, product.Store)

	if err != nil {
		log.Error("Yeni ürün eklenirken hata oluştu", err)
		return err
	}
	log.Info(fmt.Printf("Ürün eklendi: %v", addNewProduct))
	return nil
}

// extractProductsFromRows, ürün bilgilerini pgx.Rows nesnesinden çıkarır.
func extractProductsFromRows(productRows pgx.Rows) []domain.Product {
	var products = []domain.Product{}
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	for productRows.Next() {
		productRows.Scan(&id, &name, &price, &discount, &store)
		products = append(products, domain.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Discount: discount,
			Store:    store,
		})
	}
	return products
}

// GetById, belirli bir ID'ye sahip ürünü veritabanından getirir.
func (productRepository *ProductRepository) GetById(productId int64) (domain.Product, error) {
	ctx := context.Background()

	getByIdSql := `Select * from products where id = $1`

	queryRow := productRepository.dbPool.QueryRow(ctx, getByIdSql, productId)

	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	scanErr := queryRow.Scan(&id, &name, &price, &discount, &store)

	if scanErr != nil && scanErr.Error() == common.NOT_FOUND {
		return domain.Product{}, errors.New(fmt.Sprintf("ID'si %d olan ürün bulunamadı", productId))
	}
	if scanErr != nil {
		return domain.Product{}, errors.New(fmt.Sprintf("ID'si %d olan ürün alınırken hata oluştu", productId))
	}

	return domain.Product{
		Id:       id,
		Name:     name,
		Price:    price,
		Discount: discount,
		Store:    store,
	}, nil
}

// DeleteById, belirli bir ID'ye sahip ürünü veritabanından siler.
func (productRepository *ProductRepository) DeleteById(productId int64) error {
	ctx := context.Background()

	_, getErr := productRepository.GetById(productId)

	if getErr != nil {
		return errors.New("Ürün bulunamadı")
	}

	deleteSql := `Delete from products where id = $1`

	_, err := productRepository.dbPool.Exec(ctx, deleteSql, productId)
	if err != nil {
		return errors.New(fmt.Sprintf("ID'si %d olan ürün silinirken hata oluştu", productId))
	}
	log.Info("Ürün silindi")
	return nil
}

// UpdatePrice, belirli bir ID'ye sahip ürünün fiyatını günceller.
func (productRepository *ProductRepository) UpdatePrice(productId int64, newPrice float32) error {
	ctx := context.Background()

	updateSql := `Update products set price = $1 where id = $2`

	_, err := productRepository.dbPool.Exec(ctx, updateSql, newPrice, productId)

	if err != nil {
		return errors.New(fmt.Sprintf("ID'si %d olan ürünün fiyatı güncellenirken hata oluştu", productId))
	}
	log.Info("Ürün %d fiyatı %v olarak güncellendi", productId, newPrice)
	return nil
}
