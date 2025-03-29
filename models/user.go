package models

type ManagerOwner struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"fullname"`
	RoleName string `json:"role_name"`
}
