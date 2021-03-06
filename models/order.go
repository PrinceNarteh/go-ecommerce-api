package models

import "time"

type Order struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	ProductID uint    `json:"productId"`
	Product   Product `gorm:"foreignkey:ProductID"`
	UserID    uint    `json:"userId"`
	User      User    `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
}
