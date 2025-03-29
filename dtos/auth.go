package dtos

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" db:"email" valid:"email,required"`
	Password string `json:"-" form:"password" db:"password" valid:"required,nefield=Email"`
	RoleId   int    `json:"roleId" form:"roleId" db:"role_id"`
}

type UserRegist struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" db:"email" valid:"email,required"`
	Password string `json:"-" form:"password" db:"password" valid:"required,minstringlength(8)"`
	RoleId   int    `json:"roleId" form:"roleId" db:"role_id"`
	FullName string `json:"fullName" form:"fullName" db:"fullname" valid:"required"`
}

type Token struct {
	Token string `json:"token"`
}
