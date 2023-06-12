package subscriber

import (
	"context"
	"g05-fooddelivery/component/appctx"
	restaurantstorage "g05-fooddelivery/module/restaurant/storage"
	"g05-fooddelivery/pubsub"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	//GetUserId() int
}

//func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
//	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)
//	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
//	go func() {
//		defer common.AppRecover()
//		for {
//			msg := <-c
//			likeData := msg.Data().(HasRestaurantId)
//			_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
//		}
//	}()
//}

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}

}

//func PushNotification(appCtx appctx.AppContext, ctx context.Context) {
//	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)
//	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
//	go func() {
//		defer common.AppRecover()
//		for {
//			msg := <-c
//			likeData := msg.Data().(HasRestaurantId)
//			_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
//		}
//	}()
//}
