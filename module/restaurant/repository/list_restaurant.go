package restaurantrepo

import (
	"context"
	"g05-fooddelivery/common"
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

//	type LikeRestaurantStore interface {
//		GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
//	}
type listRestaurantRepo struct {
	store ListRestaurantStore
}

func NewListRestaurantRepo(store ListRestaurantStore) *listRestaurantRepo {
	return &listRestaurantRepo{store: store}
}
func (biz *listRestaurantRepo) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]restaurantmodel.Restaurant, error) {
	data, err := biz.store.ListDataWithCondition(context, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EnityName, err)
	}
	//ids := make([]int, len(data))
	//for i := range ids {
	//	ids[i] = data[i].Id
	//}
	//
	//likeMap, err := biz.likeStore.GetRestaurantLikes(context, ids)
	//if err != nil {
	//	log.Println(err)
	//	return data, nil
	//}
	//for i, item := range data {
	//	data[i].LikedCount = likeMap[item.Id]
	//}

	return data, nil
}
