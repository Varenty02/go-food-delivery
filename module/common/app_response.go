package common

// common dùng để chuẩn hóa
type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

//filter đối soát thứ client gửi lên có phải thứ mà lưu vào server hay không
//omitempty nếu có thi lay neu khong co thì bo qua

func NewSuccessResponse(data, paging, filter interface{}) *successRes {
	return &successRes{Data: data, Paging: paging, Filter: filter}
}
func SimpleSuccessResponse(data interface{}) *successRes {
	return NewSuccessResponse(data, nil, nil)
}
