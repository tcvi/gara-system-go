package domain

import "time"

type User struct {
	ID          int64     `gorm:"column:id;primaryKey;not null" dynamodbav:"id"`
	PhoneNumber string    `gorm:"column:phone_number;not null" dynamodbav:"phone_number"`
	UserName    string    `gorm:"column:user_name;not null" dynamodbav:"user_name"`
	Password    string    `gorm:"column:password" dynamodbav:"password"`
	CreatedAt   time.Time `gorm:"column:created_at" dynamodbav:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" dynamodbav:"updated_at"`
	IsActive    bool      `gorm:"column:is_active" dynamodbav:"is_active"`
	ActiveCode  string    `gorm:"column:active_code" dynamodbav:"active_code"`
	ExpCode     time.Time `gorm:"column:exp_code" dynamodbav:"exp_code"`
}

type VerifyRequest struct {
	Code     string `json:"code"`
	UserName string `json:"user_name"`
}

type RegisterUserReq struct {
	UserName string `json:"user_name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type ResendCodeRequest struct {
	UserName string `json:"user_name"`
}
