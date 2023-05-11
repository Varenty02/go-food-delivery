package restaurantbiz

import (
	"context"
	"errors"
	restaurantmodel "g05-fooddelivery/module/restaurant/model"
)

// tạo ra interface các struct khác implement lấy các hàm
type CreateRestaurantStore interface {
	CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error
}
type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}
func (biz *createRestaurantBiz) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	//goi xuong tang storage
	if data.Name == "" {
		return errors.New("Name cannot be empty")
	}
	if err := biz.store.CreateRestaurant(context, data); err != nil {
		return err
	}
	return nil
}
