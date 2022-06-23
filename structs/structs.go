package structs

type Translate struct {
	Text string `json:"text"`
	To   string `json:"to"`
}
type TranslationRes struct {
	Translation []Translate `json:"translations"`
}

// type ChatLog struct {
// 	Name string `gorm:"name"`
// 	Text string `gorm:"text"`
// 	Time string `gorm:"time"`
// }

type User struct {
	//gorm.Model
	Username string `form:"username" binding:"required" gorm:"unique;not null"`
	Password string `form:"password" binding:"required"`
	Birthday string `form:"birthday" binding:"required"`
}

type APIKey struct {
	Key      string `json:"key"`
	Location string `json:"location"`
	Endpoint string `json:"endpoint"`
	URI      string `json:"uri"`
}

type Session struct {
	UID interface{} //`gorm:"primary_key"`
}

type JsonMessage struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type JsonReturn struct {
	Name string `json:"name"`
	Text string `json:"text"`
}
