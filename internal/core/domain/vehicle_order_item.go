package domain

import "time"

type VehicleOrderItem struct {
	ID            int64     `gorm:"column:id;primaryKey;not null"`
	ItemID        int64     `gorm:"column:item_id;not null"`
	VehicleOderID int64     `gorm:"column:vehicle_order_id;not null"`
	Note          string    `gorm:"column:note"`
	Price         int64     `gorm:"column:price;not null"`
	Quantity      int64     `gorm:"column:quantity;not null"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

type VehicleOrderItemModel struct {
	ID        int64     `json:"id"`
	Item      ItemModel `json:"item"`
	Note      string    `json:"note"`
	Price     int64     `json:"price"`
	Quantity  int64     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VehicleOrderItemReq struct {
	ItemID   int64  `json:"item_id"`
	Note     string `json:"note"`
	Price    int64  `json:"price"`
	Quantity int64  `json:"quantity"`
}

type UpdateVehicleOrderItem struct {
	ID int64 `json:"id"`
	VehicleOrderItemReq
}

type CreateVehicleOrderItemRequest struct {
	VehicleOrderID int64                 `param:"id"`
	Items          []VehicleOrderItemReq `json:"items"`
}

type UpdateVehicleOrderItemRequest struct {
	VehicleOrderID int64                    `param:"id"`
	Items          []UpdateVehicleOrderItem `json:"items"`
}
