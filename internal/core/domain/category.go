package domain

import "time"

type Category struct {
	ID        int64     `gorm:"column:id;primaryKey;not null"`
	Name      string    `gorm:"column:name;not null"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type CategoryModel struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCategoryReq struct {
	Name string `json:"name"`
}

type UpdateCategoryReq struct {
	ID   int64  `param:"id"`
	Name string `json:"name"`
}
