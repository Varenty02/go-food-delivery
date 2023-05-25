package restaurantbiz

import (
	"context"
	"errors"
	"g05-fooddelivery/common"
	restaurantmodel "g05-fooddelivery/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*restaurantmodel.Restaurant, error)
	Delete(context context.Context, id int) error
}
type deleteRestaurantBiz struct {
	store     DeleteRestaurantStore
	requester common.Requester
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore, requester common.Requester) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store, requester: requester}
}
func (biz *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {
	oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})
	//find data chưa chắc db work ok(VD:too many connection)
	if err != nil {
		//common.ErrEntityNotFound(restaurantmodel.EntityName,err)
		return err
	}
	if oldData.Status == 0 {
		//common.ErrEntityDeleted(restaurantmodel.EntityName,nil)
		return errors.New("id empty")
	}
	if oldData.UserId != biz.requester.GetUserId() {
		return common.ErrNoPermission(nil)
	}
	if err := biz.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EnityName, nil)
	}
	return nil
}
