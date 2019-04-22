package entity

// 用户实体
type User struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"createTime"`
}
