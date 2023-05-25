package reslikebiz

import (
	"context"
	"g05-fooddelivery/common"
	restaurantlikemodel "g05-fooddelivery/module/restaurantlike/model"
	"log"
	"time"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}
type IncLikeCountResStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}
type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	incStore IncLikeCountResStore
}

//type userLikeRestaurantBiz struct {
//	store  UserLikeRestaurantStore
//	pubsub pubsub.Pubsub
//}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore, incStore IncLikeCountResStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{

		store:    store,
		incStore: incStore,
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
	//biz.pubsub.Publish(ctx, common2.TopicUserLikeRestaurant, pubsub.NewMessage(data))
	go func() {
		defer common.AppRecover()
		time.Sleep(time.Second * 3)
		if err := biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
