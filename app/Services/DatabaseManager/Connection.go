package DatabaseManager

import (
	"fmt"
	. "hitalent/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connection Get a database connection instance./**
func Connection(connection string) (*gorm.DB, error) {
	if connection == "" {
		connectionName, isTrue := Config("database.default", "pgsql").(string)
		if isTrue {
			connection = connectionName
		}
	}
	driverName := Config("database.connections."+connection+".driver", nil)

	switch driverName {
	case "pgsql":
		return postgreSQLSchemaDriver(connection)
	case nil:
		return nil, fmt.Errorf("unsupported database connection: %v", connection)
	default:
		return nil, fmt.Errorf("unsupported database driver: %v", driverName)
	}
}

func postgreSQLSchemaDriver(connection string) (*gorm.DB, error) {
	host := Config("database.connections."+connection+".host", "127.0.0.1")
	port := Config("database.connections."+connection+".port", "5432")
	database := Config("database.connections."+connection+".database", "hitalent")
	username := Config("database.connections."+connection+".username", "root")
	password := Config("database.connections."+connection+".password", "")
	searchPath := Config("database.connections."+connection+".search_path", "public")
	sslmode := Config("database.connections."+connection+".sslmode", "prefer")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=%s",
		host,
		port,
		username,
		password,
		database,
		searchPath,
		sslmode,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
