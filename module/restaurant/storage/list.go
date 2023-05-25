package restaurantstorage

import (
	"context"
	common2 "g05-fooddelivery/common"
	restaurantmodel "g05-fooddelivery/module/restaurant/model"
	"log"
	"strconv"
)

func (s *sqlStore) ListDataWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common2.Paging,
	moreKey ...string,
) ([]restaurantmodel.Restaurant, error) {
	var data = []restaurantmodel.Restaurant{}
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())

	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("user_id=?", f.OwnerId)
		}
		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common2.ErrDB(err)
	}
	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}
	log.Println(paging.FakeCursor)
	if v := paging.FakeCursor; v != "" {
		if cursorId, err := strconv.Atoi(v); err != nil {
			return nil, err
		} else {
			db = db.Where("id<?", cursorId)
		}

	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}
	//offset := (paging.Page - 1) * paging.Limit
	//db = db.Offset(offset)
	log.Println(paging.Limit)
	if err := db.Limit(paging.Limit).Order("id desc").Find(&data).Error; err != nil {
		return nil, common2.ErrDB(err)
	}
	if len(data) > 0 {
		last := data[len(data)-1]
		paging.NextCursor = strconv.Itoa(last.Id)
	}
	return data, nil
}
