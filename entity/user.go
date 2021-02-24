package entity

// User 用户
type User struct {
	*Entity
	Name     string `json:"name" gorm:"type:string;size:50"`
	Username string `json:"username" gorm:"type:string;size:50"`
	Password string `json:"password" gorm:"type:string;size:200"`
}

// TableName 表名
func (User) TableName() string {
	return "sys_user"
}
