package model

import (
	"time"

	"github.com/google/uuid"
)

const TableNameUser = "user"

type UserData struct {
	ID        string    `gorm:"column:id;type:varchar;primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now();type:timestamp(3);index" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:now();type:timestamp(3);index" json:"updated_at"`
	Username  string    `gorm:"column:username;not null;type:varchar" json:"username"`
	Fullname  string    `gorm:"column:fullname;not null;type:varchar" json:"fullname"`
}

func (*UserData) TableName() string {
	return TableNameUser
}

func NewFlag() UserData {
	return UserData{
		ID: uuid.New().String(),
	}
}
