package gorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormConnection() (*gorm.DB, error) {
	// TODO: 外部から設定をもらう
	dsn := "host=db user=sudoku password=sudoku dbname=sudoku port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
