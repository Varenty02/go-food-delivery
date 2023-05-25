package restaurantbiz

import (
	"context"
	"g05-fooddelivery/common"
	restaurantmodel "g05-fooddelivery/module/restaurant/model"
)

type ListRestaurantRepo interface {
	ListRestaurant(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]restaurantmodel.Restaurant, error)
}
type listRestaurantBiz struct {
	repo ListRestaurantRepo
}

func NewListRestaurantBiz(repo ListRestaurantRepo) *listRestaurantBiz {
	return &listRestaurantBiz{repo: repo}
}
func (biz *listRestaurantBiz) ListDataWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]restaurantmodel.Restaurant, error) {
	data, err := biz.repo.ListRestaurant(context, filter, paging, "User")
	if err != nil {
		return nil, err
	}
	return data, nil
}
