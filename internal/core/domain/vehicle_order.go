package domain

import "time"

type VehicleOrder struct {
	ID        int64     `gorm:"column:id;primaryKey;not null"`
	UserID    int64     `gorm:"column:user_id;not null"`
	HandlerID int64     `gorm:"column:handler_id;not null"`
	Status    Status    `gorm:"column:status;not null"`
	Note      string    `gorm:"column:note"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type CreateVehicleOrderRequest struct {
	UserID    int64  `json:"user_id"`
	HandlerID int64  `json:"handler_id"`
	Note      string `json:"note"`
}

type UpdateVehicleOrderRequest struct {
	ID        int64  `param:"id"`
	HandlerID int64  `json:"handler_id"`
	Note      string `json:"note"`
	Status    string `json:"status"`
}

type VehicleOrderModel struct {
	ID              int64                            `json:"id"`
	User            UserModel                        `json:"user"`
	Handler         UserModel                        `json:"handler"`
	Status          Status                           `json:"status"`
	Note            string                           `json:"note"`
	HistoryStatuses []VehicleOrderStatusHistoryModel `json:"history_statuses,omitempty"`
	Items           []VehicleOrderItemModel          `json:"items,omitempty"`
	CreatedAt       time.Time                        `json:"created_at"`
	UpdatedAt       time.Time                        `json:"updated_at"`
}

type FilterRequest struct {
	Status    string `query:"status"`
	SearchKey string `query:"search_key"`
	UserID    int64  `query:"user_id"`
	HandlerID int64  `query:"handler_id"`
}
