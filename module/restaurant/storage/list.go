package restaurantstorage

import (
	"context"
	"g05-fooddelivery/module/common"
	restaurantmodel "g05-fooddelivery/module/restaurant/model"
)

func (s *sqlStore) ListDataWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]restaurantmodel.Restaurant, error) {
	var data = []restaurantmodel.Restaurant{}
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("status in (1)")
	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("Owner_id=?", f.OwnerId)
		}
		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.OwnerId)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	offset := (paging.Page - 1) * paging.Limit
	if err := db.Offset(offset).Limit(paging.Limit).Order("id desc").Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return data, nil
}
