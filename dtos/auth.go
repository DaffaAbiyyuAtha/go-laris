package dtos

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" db:"email" valid:"email,required"`
	Password string `json:"-" form:"password" db:"password" valid:"required,nefield=Email"`
	RoleId   int    `json:"roleId" form:"roleId" db:"role_id"`
}

type Token struct {
	Token string `json:"token"`
}
