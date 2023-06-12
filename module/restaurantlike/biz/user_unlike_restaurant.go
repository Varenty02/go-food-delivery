package reslikebiz

import (
	"context"
	"g05-fooddelivery/common"
	restaurantlikemodel "g05-fooddelivery/module/restaurantlike/model"
	"g05-fooddelivery/pubsub"
	"log"
)

type UserUnlikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

//	type DecLikeCountResStore interface {
//		DecreateLikeCount(ctx context.Context, id int) error
//	}
type userUnlikeRestaurantBiz struct {
	store UserUnlikeRestaurantStore
	//decStore DecLikeCountResStore
	ps pubsub.Pubsub
}

func NewUserUnlikeRestaurantBiz(
	store UserUnlikeRestaurantStore,
	// decStore DecLikeCountResStore,
	ps pubsub.Pubsub,
) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{
		store: store,
		//decStore: decStore,
		ps: ps,
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
	if err := biz.ps.Publish(ctx, common.TopicUserUnLikeRestaurant, pubsub.NewMessage(&restaurantlikemodel.Like{RestaurantId: restaurantId})); err != nil {
		log.Println(err)
	}
	//sideeffect
	//j := asyncjob.NewJob(func(ctx context.Context) error {
	//	return biz.decStore.DecreateLikeCount(ctx, userId)
	//})
	//if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
	//	log.Println(err)
	//}

	//go func() {
	//	defer common.AppRecover()
	//	if err := biz.decStore.DecreateLikeCount(ctx, userId); err != nil {
	//		log.Println(err)
	//	}
	//}()
	return nil
}
