package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"product-app/common/app"
	"product-app/common/postgresql"
	"product-app/controller"
	"product-app/persistence"
	"product-app/service"
)

func main() {
	// Uygulama bağlamını başlatıyoruz.
	ctx := context.Background()

	// Echo framework'ü başlatıyoruz.
	e := echo.New()

	// Konfigürasyon yöneticisini oluşturuyoruz.
	configurationManager := app.NewConfigurationManager()

	// PostgreSQL bağlantı havuzunu oluşturuyoruz.
	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgreSqlConfig)

	// Ürün repository'sini (veri erişim katmanı) oluşturuyoruz.
	productRepository := persistence.NewProductRepository(dbPool)

	// Ürün servisini (iş mantığı katmanı) oluşturuyoruz.
	productService := service.NewProductService(productRepository)

	// Ürün kontrolcüsünü (API uç noktalarını yöneten katman) oluşturuyoruz.
	productController := controller.NewProductController(productService)

	// Kontrolcünün API rotalarını Echo'ya kaydediyoruz.
	productController.RegisterRoutes(e)

	// Sunucuyu başlatıyoruz.
	e.Start("localhost:8080")
}
