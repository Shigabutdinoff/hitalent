package Chat

import (
	"hitalent/app/Models"
	"hitalent/app/Models/Message"
	"time"
)

// The model being queried.
var model = Chat{Model: Model.NewInstance()}

type Chat struct {
	Model    Model.Model       `gorm:"-"`
	Messages []Message.Message `gorm:"foreignKey:ChatId"`

	Id        int64
	Title     string
	CreatedAt *time.Time
}

// GetModel Get the model instance being queried.
func GetModel() Chat {
	return model
}

// TableName Get the table associated with the model.
func (Chat) TableName() string {
	return "chats"
}

// Find a model by its primary key.
func Find(id int64) (Chat, error) {
	chat := GetModel()
	if err := chat.Model.GetConnection().First(&chat, id).Error; err != nil {
		return chat, err
	}

	return chat, nil
}
