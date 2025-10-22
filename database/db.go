package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

func NewConnection() (*gorm.DB, error) {
	err := godotenv.Load(); if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %w", err);
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_SSLMODE"),
	);

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err);
	}

	err = db.AutoMigrate(&SensorData{}); if err != nil {
		return nil, fmt.Errorf("Failed to migrate database: %w", err);
	}

	return db, nil;
}

func GetAllData(db *gorm.DB) ([]SensorData, error) {
	var alldata []SensorData;
	result := db.Find(&alldata);
	return alldata, result.Error;
}

func AddData(db *gorm.DB, data SensorData) {
	db.Create(&data);

	fmt.Println("Data added successfully");
}

func UpdateData(db *gorm.DB, id uint, data SensorData) error {

	result := db.Model(&SensorData{}).Where("id = ?", id).Updates(data);
	if result.Error !=  nil {
		return fmt.Errorf("Failed to update specific field: %w", result.Error)
	}

	fmt.Println("Data updated successfully");
	return nil;
}

func DeleteData() {
	
}