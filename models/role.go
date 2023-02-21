package models

type Role struct {
	Id          int
	Title       string
	Description string
	Status      int
	AddTime     int
	Access      []*Access `json:"-" gorm:"many2many:role_access;"` //角色权限之间的多对多
}

func (Role) TableName() string {
	return "role"
}
