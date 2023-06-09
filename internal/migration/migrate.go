package migration

import (
	"fmt"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/retail-ai-test/internal/pkg/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	m := gormigrate.New(
		db,
		gormigrate.DefaultOptions,
		[]*gormigrate.Migration{
			{
				ID: "Initial",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.Recipe{}, &model.User{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable(model.Recipe{}.TableName(), "users")
				},
			},
		})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	fmt.Println("Database migration did run successfully")
}
