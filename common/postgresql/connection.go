package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

// GetConnectionPool, PostgreSQL veritabanına bağlantı havuzu oluşturur.
func GetConnectionPool(context context.Context, config Config) *pgxpool.Pool {
	// Sağlanan yapılandırma parametreleriyle bağlantı dizesini oluşturur.
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable statement_cache_mode=describe pool_max_conns=%s pool_max_conn_idle_time=%s",
		config.Host,
		config.Port,
		config.UserName,
		config.Password,
		config.DbName,
		config.MaxConnections,
		config.MaxConnectionIdleTime)

	// Bağlantı dizesini pgxpool için yapılandırmaya çevirir.
	connConfig, parseConfigErr := pgxpool.ParseConfig(connString)
	if parseConfigErr != nil {
		panic(parseConfigErr) // Hata durumunda programı durdurur.
	}

	// Bağlantı havuzunu yapılandırmaya göre oluşturur.
	conn, err := pgxpool.ConnectConfig(context, connConfig)
	if err != nil {
		log.Error("Veritabanına bağlanılamıyor: %v\n", err)
		panic(err) // Hata durumunda programı durdurur.
	}

	return conn // Başarılı bağlantı havuzunu döner.
}
