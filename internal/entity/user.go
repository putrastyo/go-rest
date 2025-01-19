package entity

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(100)" validate:"required"`
	Email     string    `json:"email" gorm:"type:varchar(100);unique" validate:"required,email"`
	Password  string    `json:"password" gorm:"type:varchar(100)" validate:"required,min=8"`
	Phone     string    `json:"phone" gorm:"type:varchar(15)" validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Tasks     []Task    `json:"tasks" gorm:"foreignKey:UserID"`
}
