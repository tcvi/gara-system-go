package domain

import "time"

type Item struct {
	ID          int64     `gorm:"column:id;primaryKey;not null"`
	CategoryID  int64     `gorm:"column:category_id;not null"`
	Name        string    `gorm:"column:name;not null"`
	Description string    `gorm:"column:description;not null"`
	Price       int64     `gorm:"column:price;not null"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

type ItemModel struct {
	ID          int64         `json:"id"`
	Category    CategoryModel `json:"category"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Price       int64         `json:"price"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type FilterItemRequest struct {
	CategoryID int64  `query:"category_id"`
	SearchKey  string `query:"search_key"`
}

type CreateItemReq struct {
	CategoryID  int64  `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}

type UpdateItemReq struct {
	ID          int64  `param:"id"`
	CategoryID  int64  `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}
