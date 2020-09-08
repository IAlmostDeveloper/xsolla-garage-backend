package dto

type Task struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Title string `json:"title"`
	Category string `json:"category"`
	TextContent string `json:"text_content"`
	DateCreate string `json:"date_create"`
	DateClose string `json:"date_close"`
	DateTarget string `json:"date_target"`
	IsCompleted bool `json:"is_completed"`
}
