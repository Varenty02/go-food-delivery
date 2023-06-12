package reslikebiz

import (
	"context"
	"g05-fooddelivery/common"
	restaurantlikemodel "g05-fooddelivery/module/restaurantlike/model"
	"g05-fooddelivery/pubsub"
	"log"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

//	type IncLikeCountResStore interface {
//		IncreaseLikeCount(ctx context.Context, id int) error
//	}
type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
	//incStore IncLikeCountResStore
	ps pubsub.Pubsub
}

//type userLikeRestaurantBiz struct {
//	store  UserLikeRestaurantStore
//	pubsub pubsub.Pubsub
//}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore, ps pubsub.Pubsub) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{

		store: store,
		//incStore: incStore,
		ps: ps,
	}
}

// func NewUserLikeRestaurantBiz(
//
//		store UserLikeRestaurantStore,
//		pubsub pubsub.Pubsub) *userLikeRestaurantBiz {
//		return &userLikeRestaurantBiz{
//
//			store:  store,
//			pubsub: pubsub,
//		}
//	}
func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	err := biz.store.Create(ctx, data)
	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}
	//send message
	if err := biz.ps.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	//side effect
	//j := asyncjob.NewJob(func(ctx context.Context) error {
	//	return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	//})
	//if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
	//	log.Println(err)
	//}
	//go func() {
	//	defer common.AppRecover()
	//	if err := biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
	//		log.Println(err)
	//	}
	//}()

	return nil
}
