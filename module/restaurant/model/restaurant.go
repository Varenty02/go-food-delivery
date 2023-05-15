package restaurantmodel

import (
	"errors"
	"g05-fooddelivery/module/common"
	"strings"
)

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Addr            string `json:"addr" gorm:"column:addr;"`
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Addr            string `json:"addr" gorm:"column:addr;"`
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

type RestaurantUpdate struct {
	Name   *string `json:"name" gorm:"column:name;"`
	Addr   *string `json:"addr" gorm:"column:addr;"`
	Status *int    `json:"status" gorm:"column:status;"`
}

func (Restaurant) TableName() string       { return "restaurants" }
func (RestaurantUpdate) TableName() string { return "restaurants" }
func (RestaurantCreate) TableName() string { return "restaurants" }

var (
	ErrNameIsEmpty = errors.New("name can not empty")
)
