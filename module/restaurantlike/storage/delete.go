package restaurantlikestorage

import (
	"context"
	common2 "g05-fooddelivery/common"
	restaurantlikemodel "g05-fooddelivery/module/restaurantlike/model"
)

func (s *sqlStore) Delete(ctx context.Context, userId, restaurantId int) error {
	db := s.db
	if err := db.Table(restaurantlikemodel.Like{}.TableName()).
		Where("user_id=? and restaurant_id=?", userId, restaurantId).
		Delete(nil).
		Error; err != nil {
		return common2.ErrDB(err)
	}
	return nil
}
