package models

type PostData struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	NameUser  string `json:"name_user"`
	EmailUser string `json:"email_user"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
