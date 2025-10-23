package database

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SensorData struct {
	ID					uuid.UUID	`json:"id" gorm:"type:uuid;primaryKey"`
	Temperature			float32		`json:"temperature"`
	General_Humidity	float32		`json:"general_humidity"`
	Soil_Moisture		float32		`json:"soil_moisture"`
	Luminosity			float32		`json:"luminosity"`
	CreatedAt	        time.Time 	`json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt			time.Time 	`json:"updated_at" gorm:"autoUpdateTime"`
}

func (s *SensorData) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuid.Nil {
		s.ID = uuid.New();
	}
	return
}

func (SensorData) TableName() string {
	return "sensor_data"
}

type DataProvider interface {
	GetAllData() ([]SensorData, error)
}