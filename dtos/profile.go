package dtos

type Profile struct {
	Id         int     `json:"id"`
	Picture    *string `json:"picture"  db:"picture"`
	FullName   string  `json:"fullName"  valid:"type(string),required" form:"fullname" db:"fullname"`
	Province   *string `json:"province" db:"province" form:"province"`
	City       *string `json:"city" db:"city" form:"city"`
	PostalCode *string `json:"postalCode" db:"postal_code" form:"postalCode"`
	Country    *string `json:"country" db:"country" form:"country"`
	Mobile     *string `json:"mobile" db:"mobile" form:"mobile"`
	Address    *string `json:"address" db:"address" form:"address"`
	UserId     int     `json:"userId" form:"userId" db:"user_id"`
}

type JoinRegist struct {
	Id       int     `json:"id"`
	Email    *string `json:"email" form:"email" db:"email" binding:"required,email"`
	Password string  `form:"password" db:"password" binding:"required" json:"-"`
	RoleId   int     `json:"roleId" form:"roleId" db:"role_id"`
	Results  Profile
}
