package restaurantlikemodel

import (
	"fmt"
	common2 "g05-fooddelivery/common"
	"time"
)

type Like struct {
	RestaurantId int                 `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int                 `json:"user_id" gorm:"user_id"`
	CreatedAt    *time.Time          `json:"created_at" gorm:"column:created_at;"`
	User         *common2.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (Like) TableName() string { return "restaurant_likes" }
func (l *Like) GetRestaurantId() int {
	return l.RestaurantId
}
func ErrCannotLikeRestaurant(err error) *common2.AppError {
	return common2.NewCustomError(err, fmt.Sprintf("Cannot like this restaurant"), fmt.Sprintf("ErrCannotLikeRestaurant"))
}
func ErrCannotUnlikeRestaurant(err error) *common2.AppError {
	return common2.NewCustomError(err, fmt.Sprintf("Cannot unlike this restaurant"), fmt.Sprintf("ErrCannotUnlikeRestaurant"))
}
