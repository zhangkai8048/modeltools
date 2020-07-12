package models

type User struct {
	Id int64 `json:"id"` // 主键
	UserId int64 `json:"user_id"` // 用户信息id(user_info 的id)
	UserName string `json:"user_name"` // 用户名
	Name string `json:"name"` // 姓名
	Mobile string `json:"mobile"` // 手机号
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
	Version int `json:"version"` // 版本号
	IsDelete int `json:"is_delete"` // 是否删除,0:否 1:是
}