package initialize

import (
	db "github.com/retail-ai-test/internal/database"
	"go.uber.org/zap"
)

func InitDB() error {
	if err := db.SetupConn(); err != nil {
		zap.S().Errorf("Failed to setup database connection: %v", err)
		return err
	}
	return nil
}
