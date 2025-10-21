package database

import (
	"time"
)

type SensorData struct {
	ID					uint	`json:"id" gorm:"primaryKey"`
	Temperature			float32	`json:"temperature"`
	General_Humidity	float32	`json:"general_humidity"`
	Soil_Moisture		float32	`json:"soil_moisture"`
	Luminosity			float32	`json:"luminosity"`
	CreatedAt	        time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt			time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (SensorData) TableName() string {
	return "sensor_data"
}

type DataProvider interface {
	GetAllData() ([]SensorData, error)
}