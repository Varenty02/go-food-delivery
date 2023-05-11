package restaurantstorage

import (
	"context"
	restaurantmodel "g05-fooddelivery/module/restaurant/model"
)

// context hỗ trợ io giúp json đc ,cancel routine được,...
// không viết hàm này tầng business vẫn không ảnh hưởng
func (s *sqlStore) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
