package dtos

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" db:"email" binding:"required,email"`
	Password string `form:"password" db:"password" binding:"required" json:"-"`
	RoleId   int    `json:"roleId" form:"roleId" db:"role_id"`
}

type Profile struct {
	Id int `json:"id"`
}

type Token struct {
	Token string `json:"token"`
}
