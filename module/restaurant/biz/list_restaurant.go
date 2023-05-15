package restaurantbiz

import (
	"context"
	"g05-fooddelivery/module/common"
	restaurantmodel "g05-fooddelivery/module/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]restaurantmodel.Restaurant, error)
}
type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}
func (biz *listRestaurantBiz) ListDataWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]restaurantmodel.Restaurant, error) {
	data, err := biz.store.ListDataWithCondition(context, filter, paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
