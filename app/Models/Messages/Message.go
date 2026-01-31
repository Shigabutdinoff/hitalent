package Messages

import (
	"hitalent/app/Models"
	"time"
)

// The model being queried.
var model = Message{Model: Model.NewInstance()}

type Message struct {
	Model Model.Model `gorm:"-"`

	Id        int64
	ChatId    int64
	Text      string
	CreatedAt *time.Time
}

// GetModel Get the model instance being queried.
func GetModel() Message {
	return model
}

// TableName Get the table associated with the model.
func (Message) TableName() string {
	return "messages"
}

// Find a model by its primary key.
func Find(id int64) (Message, error) {
	message := GetModel()
	if err := message.Model.GetConnection().First(&message, id).Error; err != nil {
		return message, err
	}

	return message, nil
}
