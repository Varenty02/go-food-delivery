package restaurantlikestorage

import (
	"context"
	"fmt"
	common2 "g05-fooddelivery/common"
	restaurantlikemodel "g05-fooddelivery/module/restaurantlike/model"
	"log"
	"time"
)

const timeLayout = "2006-01-02T15:04:05.999999"

func (s *sqlStore) GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)
	type sqlData struct {
		RestaurantId int `gorm:"column:restaurant_id;"`
		LikeCount    int `gorm:"column:count;"`
	}
	var listLike []sqlData
	if err := s.db.Table(restaurantlikemodel.Like{}.TableName()).
		Select("restaurant_id,count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").Find(&listLike).Error; err != nil {
		return nil, common2.ErrDB(err)
	}
	for _, item := range listLike {
		result[item.RestaurantId] = item.LikeCount
	}
	return result, nil
}
func (s *sqlStore) GetUsersLikeRestaurant(ctx context.Context,
	condition map[string]interface{},
	filter *restaurantlikemodel.Filter,
	paging *common2.Paging,
	moreKey ...string,
) ([]common2.SimpleUser, error) {
	var result []restaurantlikemodel.Like
	db := s.db
	db = db.Table(restaurantlikemodel.Like{}.TableName()).Where(condition)
	if v := filter; v != nil {
		if v.RestaurantId > 0 {
			db = db.Where("restaurant_id=?", v.RestaurantId)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common2.ErrDB(err)
	}
	db = db.Preload("User")
	if v := paging.FakeCursor; v != "" {
		log.Println(paging.FakeCursor)
		timeCreated, err := time.Parse(timeLayout, v)
		if err != nil {
			return nil, common2.ErrDB(err)
		}
		db = db.Where("created_at<?", timeCreated.Format("2006-01-02 15:04:05"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}
	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common2.ErrDB(err)
	}
	users := make([]common2.SimpleUser, len(result))
	log.Println(result[0].User)
	for i, item := range result {
		result[i].User.CreatedAt = item.CreatedAt
		result[i].User.UpdatedAt = nil
		users[i] = *result[i].User
		if i == len(result)-1 {
			cursorStr := fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))
			paging.NextCursor = cursorStr
		}
	}
	return users, nil
}
