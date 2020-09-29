package dto

type Feedback struct{
	Id int `json:"id" db:"id"`
	DateCreate TimeJson `json:"date_create" db:"date_create"`
	Content *string `json:"content" db:"content"`
}

