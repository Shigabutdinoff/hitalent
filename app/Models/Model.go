package Model

import (
	"hitalent/app/Services/DatabaseManager"

	"gorm.io/gorm"
)

// The connection name for the model.
var connection, _ = DatabaseManager.Connection("")

type Model struct {
	connection *gorm.DB
}

// NewInstance Create a new instance of the given model.
func NewInstance() Model {
	return Model{
		connection: connection,
	}
}

// GetConnection Get the database connection for the model.
func (model Model) GetConnection() *gorm.DB {
	return model.connection
}
