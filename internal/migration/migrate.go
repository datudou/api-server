package migration

import (
	"fmt"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/retail-ai-test/internal/model"
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
					return tx.AutoMigrate(&model.Recipe{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable("recipes")
				},
			},
			// {
			// 	ID: "inital_data",
			// 	Migrate: func(tx *gorm.DB) error {
			// 		di := NewDataInitial(tx)
			// 		err := di.initTeamData()
			// 		if err != nil {
			// 			return err
			// 		}
			// 		err = di.initPlayerData()
			// 		if err != nil {
			// 			fmt.Printf("Error: %v", err)
			// 			return err
			// 		}
			// 		return nil
			// 	},
			// },
		})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	fmt.Println("Database migration did run successfully")
}
