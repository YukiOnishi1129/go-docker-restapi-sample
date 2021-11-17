package models

type Todo struct {
	BaseModel
	Title     string `gorm:"size:255" json:"title,omitempty"`
	Comment   string `gorm:"type:text" json:"comment,omitempty"`
	UserId    int `gorm:"not null" json:"user_id"`
	User      User `gorm:"foreignKey:UserId"`
}

type MutationTodoRequest struct {
	Title     string `json:"title,omitempty"`
	Comment   string `json:"comment,omitempty"`
}