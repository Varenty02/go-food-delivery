package common

import "time"

type SQLModel struct {
	Id       int        `json:"id" gorm:"column:id";`
	Status   int        `json:"status" gorm:"column:status";`
	CreateAt *time.Time `json:"createAt" gorm:"column:create_at"";`
	UpdateAt *time.Time `json:"updateAt" gorm:"column:update_at";`
}
