package restaurantstorage

import (
	"context"
	"g05-fooddelivery/common"
	restaurantmodel "g05-fooddelivery/module/restaurant/model"
)

func (s *sqlStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
