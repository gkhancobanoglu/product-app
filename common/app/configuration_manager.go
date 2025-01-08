package app

import "product-app/common/postgresql"

// ConfigurationManager, uygulama ayarlarını yöneten bir yapı tanımıdır.
type ConfigurationManager struct {
	PostgreSqlConfig postgresql.Config // PostgreSQL bağlantı ayarlarını tutar.
}

// NewConfigurationManager, yeni bir ConfigurationManager nesnesi oluşturur ve döndürür.
func NewConfigurationManager() *ConfigurationManager {
	postgreSqlConfig := getPostgreSqlConfig() // PostgreSQL ayarlarını alır.
	return &ConfigurationManager{
		PostgreSqlConfig: postgreSqlConfig, // ConfigurationManager içinde PostgreSQL ayarlarını saklar.
	}
}

// getPostgreSqlConfig, PostgreSQL bağlantı ayarlarını döndürür.
func getPostgreSqlConfig() postgresql.Config {
	return postgresql.Config{
		Host:                  "localhost",  // Veritabanı sunucusunun adresi.
		Port:                  "6432",       // Veritabanı bağlantı portu.
		UserName:              "postgres",   // Veritabanı kullanıcı adı.
		Password:              "postgres",   // Veritabanı kullanıcı şifresi.
		DbName:                "productapp", // Bağlanılacak veritabanı adı.
		MaxConnections:        "10",         // Maksimum bağlantı sayısı.
		MaxConnectionIdleTime: "30s",        // Maksimum bağlantı boşta kalma süresi.
	}
}
