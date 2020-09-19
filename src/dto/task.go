package dto

type Task struct {
	Id          int     `json:"id" db:"id"`
	UserId      int     `json:"user_id" db:"user_id"`
	Title       *string `json:"title" db:"title"`
	TextContent *string `json:"text_content" db:"text_content"`
	DateCreate  *string `json:"date_create" db:"date_create"`
	DateClose   *string `json:"date_close" db:"date_close"`
	DateTarget  *string `json:"date_target" db:"date_target"`
	IsCompleted bool    `json:"is_completed" db:"is_completed"`
	IsImportant bool    `json:"is_important" db:"is_important"`
	IsUrgent    bool    `json:"is_urgent" db:"is_urgent"`
	Tags        []*Tag  `json:"tags"`
}
