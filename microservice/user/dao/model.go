package dao

import (
	"workbench/model"
)

type UserModel struct {
	Id           uint32 `gorm:"column:id;not null" binding:"required"`
	Name         string `gorm:"column:name;" binding:"required"`
	RealName     string `gorm:"column:real_name;" binding:"required"`
	Email        string `gorm:"column:email;" binding:"required"`
	Avatar       string `gorm:"column:avatar;" binding:"required"`
	Tel          string `gorm:"column:tel;"`
	Role         uint32 `gorm:"column:role;" binding:"required"`
	TeamId       uint32 `gorm:"column:team_id;" binding:"required"`
	GroupId      uint32 `gorm:"column:group_id;" binding:"required"`
	EmailService uint32 `gorm:"column:email_service;" binding:"required"`
	Message      uint32 `gorm:"column:message;" binding:"required"`
}

func (UserModel) TableName() string {
	return "users"
}

// Create ...
func (u *UserModel) Create() error {
	return model.DB.Self.Create(u).Error
}

// Save ...
func (u *UserModel) Save() error {
	return model.DB.Self.Save(u).Error
}
