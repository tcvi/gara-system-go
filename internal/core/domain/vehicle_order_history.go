package domain

import "time"

type VehicleOrderStatusHistory struct {
	ID             int64     `gorm:"column:id;primaryKey;not null"`
	VehicleOrderID int64     `gorm:"column:vehicle_order_id;not null"`
	Status         Status    `gorm:"column:status;not null"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

type VehicleOrderStatusHistoryModel struct {
	ID             int64     `json:"id"`
	VehicleOrderID int64     `json:"vehicle_order_id"`
	Status         Status    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
