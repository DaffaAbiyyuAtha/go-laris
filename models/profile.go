package models

type Profile struct {
	Id         int     `json:"id"`
	Picture    *string `json:"picture"`
	FullName   string  `json:"fullName"`
	Province   *string `json:"province"`
	City       *string `json:"city"`
	PostalCode *string `json:"postalCode"`
	Gender     *int    `json:"gender"`
	Country    *string `json:"country"`
	Mobile     *int    `json:"mobile"`
	Address    *string `json:"address"`
	UserId     int     `json:"userId"`
}
