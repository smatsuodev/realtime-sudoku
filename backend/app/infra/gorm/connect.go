package gorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormConnection() (*gorm.DB, error) {
	// TODO: 外部から設定をもらう
	dsn := "host=db user=sudoku password=sudoku dbname=sudoku port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// ここでやるべきか?
	if err := migrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
