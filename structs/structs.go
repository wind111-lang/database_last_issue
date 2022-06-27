package structs

type Translate struct {
	Text string `json:"text"`
	To   string `json:"to"`
}
type TranslationRes struct {
	Translation []Translate `json:"translations"`
}

type ChatLog struct {
	Username string `gorm:"username"`
	Text     string `gorm:"text"`
}

type User struct {
	//gorm.Model
	Id       int    `gorm:"primary_key"`
	Username string `form:"username" binding:"required" gorm:"unique;not null"`
	Password string `form:"password" binding:"required"`
	Birthday string `form:"birthday" binding:"required"`
}

// type LoggedInUser struct {
// 	Username string `json:"username"`
// }

type Session struct {
	UID interface{}
}

type JsonMessage struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

type JsonReturn struct {
	Name string `json:"name"`
	Text string `json:"text"`
}
