package restaurantstorage

import (
	"context"
	restaurantmodel "g05-fooddelivery/module/restaurant/model"
)

func (s *sqlStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
