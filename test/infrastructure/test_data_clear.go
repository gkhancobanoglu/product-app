package infrastructure

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

func TruncateTestData(ctx context.Context, dbPool *pgxpool.Pool) {
	// 'products' tablosunu sıfırlamak için truncate işlemi gerçekleştirilir.
	_, truncateResultErr := dbPool.Exec(ctx, "TRUNCATE products RESTART IDENTITY")
	if truncateResultErr != nil {
		// Hata oluşursa loglanır.
		log.Error(truncateResultErr)
	} else {
		// İşlem başarılıysa bilgi logu yazdırılır.
		log.Info("Products table truncated")
	}
}
