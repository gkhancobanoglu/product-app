package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"product-app/controller/request"
	"product-app/controller/response"
	"product-app/service"
	"strconv"
)

// ProductController, ürünlerle ilgili işlemleri yöneten bir kontrolcü yapısıdır.
type ProductController struct {
	productService service.IProductService
}

// NewProductController, yeni bir ProductController nesnesi oluşturur ve döndürür.
func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

// RegisterRoutes, ürünle ilgili API uç noktalarını Echo framework'e kaydeder.
func (productController *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/products/:id", productController.GetProductById)       // Belirli bir ürünü ID ile getirir.
	e.GET("/api/v1/products", productController.GetAllProducts)           // Tüm ürünleri listeler.
	e.POST("/api/v1/products", productController.AddProduct)              // Yeni bir ürün ekler.
	e.PUT("/api/v1/products/:id", productController.UpdatePrice)          // Belirli bir ürünün fiyatını günceller.
	e.DELETE("/api/v1/products/:id", productController.DeleteProductById) // Belirli bir ürünü siler.
}

// GetProductById, ID'ye göre bir ürünü getirir.
func (productController *ProductController) GetProductById(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.Atoi(param) // ID'yi string'den int'e çevirir.

	// Ürünü servis katmanından alır.
	product, err := productController.productService.GetById(int64(productId))
	if err != nil {
		// Eğer ürün bulunamazsa, 404 döner.
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}
	// Ürün bulunduysa, 200 ile ürün detaylarını döner.
	return c.JSON(http.StatusOK, response.ToResponse(product))
}

// GetAllProducts, tüm ürünleri veya belirli bir mağazaya ait ürünleri getirir.
func (productController *ProductController) GetAllProducts(c echo.Context) error {
	store := c.QueryParam("store") // Mağaza sorgu parametresini alır.
	if len(store) == 0 {
		// Mağaza belirtilmemişse tüm ürünleri getirir.
		allProducts := productController.productService.GetAllProducts()
		return c.JSON(http.StatusOK, response.ToResponseList(allProducts))
	}
	// Belirli bir mağazanın ürünlerini getirir.
	productsWithGivenStore := productController.productService.GetAllProductsByStore(store)
	return c.JSON(http.StatusOK, response.ToResponseList(productsWithGivenStore))
}

// AddProduct, yeni bir ürün ekler.
func (productController *ProductController) AddProduct(c echo.Context) error {
	var addProductRequest request.AddProductRequest
	bindErr := c.Bind(&addProductRequest) // Gelen isteği modele bağlar.
	if bindErr != nil {
		// Eğer bağlama sırasında hata olursa, 400 döner.
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: bindErr.Error(),
		})
	}
	// Ürünü servis katmanına ekler.
	err := productController.productService.Add(addProductRequest.ToModel())

	if err != nil {
		// Eğer ekleme sırasında doğrulama hatası oluşursa, 422 döner.
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}
	// Başarılı ekleme durumunda 201 döner.
	return c.NoContent(http.StatusCreated)
}

// UpdatePrice, bir ürünün fiyatını günceller.
func (productController *ProductController) UpdatePrice(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.Atoi(param) // ID'yi string'den int'e çevirir.

	newPrice := c.QueryParam("newPrice") // Yeni fiyat sorgu parametresini alır.
	if len(newPrice) == 0 {
		// Eğer yeni fiyat belirtilmemişse, 400 döner.
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: "Parameter newPrice is required!",
		})
	}
	// Yeni fiyatı float'a dönüştürür.
	convertedPrice, err := strconv.ParseFloat(newPrice, 32)
	if err != nil {
		// Dönüştürme sırasında hata olursa, 400 döner.
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: "NewPrice Format Disrupted!",
		})
	}
	// Ürün fiyatını servis katmanında günceller.
	productController.productService.UpdatePrice(int64(productId), float32(convertedPrice))
	return c.NoContent(http.StatusOK) // Başarılı güncelleme durumunda 200 döner.
}

// DeleteProductById, ID'ye göre bir ürünü siler.
func (productController *ProductController) DeleteProductById(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.Atoi(param) // ID'yi string'den int'e çevirir.

	// Ürünü servis katmanında siler.
	err := productController.productService.DeleteById(int64(productId))
	if err != nil {
		// Eğer ürün bulunamazsa, 404 döner.
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}
	return c.NoContent(http.StatusOK) // Başarılı silme durumunda 200 döner.
}
