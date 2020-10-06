package dto

type User struct {
	Id         string `json:"id" db:"id"`
	Email      string `json:"email" db:"email"`
	Name       string `json:"name" db:"name"`
	GivenName  string `json:"given_name" db:"given_name"`
	FamilyName string `json:"family_name" db:"family_name"`
	Locale     string `json:"locale" db:"locale"`
	Picture    string `json:"picture" db:"picture"`
}
