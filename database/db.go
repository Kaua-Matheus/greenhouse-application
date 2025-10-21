package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

type Data struct {
	ID			uint	`json:"id" gorm:"primaryKey"`
	Temperature	float32	`json:"temperature"`
}

func NewConnection() (*Database, error) {
	// Postgres Host
	dsn := ""

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err);
	}

	err = db.AutoMigrate(&Data{}); if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err);
	}

	return &Database{DB: db}, nil;
}

func (d *Database) GetAllData() ([]Data, error) {
	var alldata []Data;
	result := d.DB.Find(&alldata);
	return alldata, result.Error;
}