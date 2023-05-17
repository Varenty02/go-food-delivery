package restaurantbiz

import (
	"context"
	"errors"
	restaurantmodel "g05-fooddelivery/module/restaurant/model"
)

type FindRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*restaurantmodel.Restaurant, error)
}
type findRestaurantBiz struct {
	store FindRestaurantStore
}

func NewFindRestaurantBiz(store FindRestaurantStore) *findRestaurantBiz {
	return &findRestaurantBiz{store: store}
}
func (biz *findRestaurantBiz) FindDataWithCondition(context context.Context, condition map[string]interface{}, moreKey ...string) (*restaurantmodel.Restaurant, error) {
	data, err := biz.store.FindDataWithCondition(context, condition)

	if err != nil {
		return nil, err
	}
	if data.Status == 0 {
		return nil, errors.New("id empty")
	}
	return data, nil
}
