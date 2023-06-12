package restaurantmodel

import (
	"errors"
	common2 "g05-fooddelivery/common"
	"strings"
)

const EnityName = "restaurant"

type Restaurant struct {
	common2.SQLModel
	Name string         `json:"name" gorm:"column:name;"`
	Addr string         `json:"addr" gorm:"column:addr;"`
	Logo *common2.Image `json:"logo" gorm:"column:logo;"`
	//ĐỂ trống gorm thì luôn luôn mapping ở lúc create
	UserId    int                 `json:"-" gorm:"column:user_id;"`
	User      *common2.SimpleUser `json:"user" gorm:"preload:false;"`
	LikeCount int                 `json:"like_count" gorm:"like_count"`
}

type RestaurantCreate struct {
	common2.SQLModel
	Name   string         `json:"name" gorm:"column:name;"`
	Addr   string         `json:"addr" gorm:"column:addr;"`
	UserId int            `json:"-" gorm:"column:user_id;"`
	Logo   *common2.Image `json:"logo" gorm:"column:logo;"`
	//Cover  *common2.Images `json:"cover" gorm:"column:cover;"`
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

type RestaurantUpdate struct {
	Name   *string        `json:"name" gorm:"column:name;"`
	Addr   *string        `json:"addr" gorm:"column:addr;"`
	Status *int           `json:"status" gorm:"column:status;"`
	Logo   *common2.Image `json:"logo" gorm:"column:logo;"`
}

func (Restaurant) TableName() string       { return "restaurants" }
func (RestaurantUpdate) TableName() string { return "restaurants" }
func (RestaurantCreate) TableName() string { return "restaurants" }

var (
	ErrNameIsEmpty = errors.New("name can not empty")
)

//func (r *Restaurant) Mask(isAdminOrOwner bool) {
//	r.GenUID(common.DbTypeRestaurant)
//}
//func (r *RestaurantCreate) Mask(isAdminOrOwner bool) {
//	r.GenUID(common.DbTypeRestaurant)
//}
