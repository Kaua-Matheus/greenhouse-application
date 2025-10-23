package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/google/uuid"
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

	fmt.Printf("Connecting in the database with: %s\n", dsn);

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err);
	}

	fmt.Println("Connected to database successfully!")
    fmt.Println("Starting AutoMigrate...")

	err = db.AutoMigrate(&SensorData{}); if err != nil {
		return nil, fmt.Errorf("Failed to migrate database: %w", err);
	}

	fmt.Println("AutoMigrate completed successfully!")

	if db.Migrator().HasTable(&SensorData{}) {
        fmt.Println("Table 'sensor_data' exists!")
    } else {
        fmt.Println("WARNING: Table 'sensor_data' was not created!")
    }

	return db, nil;
}

func GetAllData(db *gorm.DB) ([]SensorData, error) {

	var alldata []SensorData;
	result := db.Find(&alldata);
	return alldata, result.Error;
}

func AddData(db *gorm.DB, data SensorData) error {

	result := db.Create(&data); if result.Error != nil {
		return fmt.Errorf("Failed to add data: %w", result.Error)
	}

	fmt.Println("Data added successfully");
	return nil;
}

func UpdateData(db *gorm.DB, id uuid.UUID, data SensorData) error {

	result := db.Model(&SensorData{}).Where("id = ?", id).Updates(data);
	if result.Error !=  nil {
		return fmt.Errorf("Failed to update specific field: %w", result.Error)
	}

	fmt.Println("Data updated successfully");
	return nil;
}

func DeleteData(db *gorm.DB, id uuid.UUID) error {
	
	result := db.Delete(&SensorData{}, id);
	if result == nil {
		return fmt.Errorf("Failed to delete the data: delete returned nil result");
	}
	if result.Error != nil {
		return fmt.Errorf("Failed to delete the data: %w", result.Error);
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("No records found with id %s", id.String());
	}

	fmt.Printf("Data deleted successfully, %d row(s) affected", result.RowsAffected);
	return nil;
}