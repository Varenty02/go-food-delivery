package restaurantmodel

type Restaurant struct {
	Id     int    `json:"id" gorm:"column:id;"`
	Name   string `json:"name" gorm:"column:name;"`
	Addr   string `json:"addr" gorm:"column:addr;"`
	Status int    `json:"status" gorm:"column:status;"`
}

// Thực tế khi tạo thông tin thường ít hơn update (VD:đăng ký user không cần điền het thong tin)
type RestaurantCreate struct {
	Id     int    `json:"id" gorm:"column:id;"`
	Name   string `json:"name" gorm:"column:name;"`
	Addr   string `json:"addr" gorm:"column:addr;"`
	Status int    `json:"status" gorm:"column:status;"`
}

// dành cho update để biết được giá trị có sự thay đổi không(gorm không nhận 0 "" false
type RestaurantUpdate struct {
	Name   *string `json:"name" gorm:"column:name;"`
	Addr   *string `json:"addr" gorm:"column:addr;"`
	Status *int    `json:"status" gorm:"column:status;"`
}

func (Restaurant) TableName() string       { return "restaurants" }
func (RestaurantUpdate) TableName() string { return "restaurants" }
func (RestaurantCreate) TableName() string { return "restaurants" }
