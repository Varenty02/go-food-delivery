package reslikebiz

import (
	"context"
	restaurantlikemodel "g05-fooddelivery/module/restaurantlike/model"
)

type UserUnlikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}
type userUnlikeRestaurantBiz struct {
	store UserUnlikeRestaurantStore
}

func NewUserUnlikeRestaurantBiz(
	store UserUnlikeRestaurantStore,
) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{
		store: store,
	}
}
func (biz *userUnlikeRestaurantBiz) UnlikeRestaurant(
	ctx context.Context,
	userId,
	restaurantId int,
) error {
	err := biz.store.Delete(ctx, userId, restaurantId)
	if err != nil {
		return restaurantlikemodel.ErrCannotUnlikeRestaurant(err)
	}
	return nil
}
