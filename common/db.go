package common

import (
	"os"

	schemas "github.com/Caicrs/Payfood-backend/graph/model"
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
	db.AutoMigrate(&schemas.User{}, &schemas.Meetup{})
	return db, nil
}
