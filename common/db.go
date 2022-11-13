package common

import (
	"os"

	customTypes "github.com/Caicrs/Payfood-backend/graph/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDb() (*gorm.DB, error) {
	godotenv.Load()
	myDb := os.Getenv("POSTGRES_URL")
	var err error
	db, err := gorm.Open(postgres.Open(myDb), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&customTypes.Marketplaces{})
	db.AutoMigrate(&customTypes.Products{})
	db.AutoMigrate(&customTypes.Tables{})
	db.AutoMigrate(&customTypes.Clients{})
	db.AutoMigrate(&customTypes.Orders{})
	db.Migrator().CreateConstraint(&customTypes.Marketplaces{}, "id")
	db.Migrator().CreateConstraint(&customTypes.Products{}, "id")
	db.Migrator().CreateConstraint(&customTypes.Tables{}, "id")
	db.Migrator().CreateConstraint(&customTypes.Clients{}, "id")
	db.Migrator().CreateConstraint(&customTypes.Orders{}, "id")
	return db, nil
}
