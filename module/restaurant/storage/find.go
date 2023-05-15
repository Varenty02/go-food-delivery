package restaurantstorage

import (
	"context"
	restaurantmodel "g05-fooddelivery/module/restaurant/model"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*restaurantmodel.Restaurant, error) {
	var data = restaurantmodel.Restaurant{}
	if err := s.db.
		Where(condition).
		First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
