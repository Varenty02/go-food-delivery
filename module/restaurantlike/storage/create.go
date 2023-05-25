package restaurantlikestorage

import (
	"context"
	common2 "g05-fooddelivery/common"
	restaurantlikemodel "g05-fooddelivery/module/restaurantlike/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantlikemodel.Like) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common2.ErrDB(err)
	}
	return nil
}
