package DatabaseManager

import (
	"fmt"
	. "hitalent/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connection Get a database connection instance./**
func Connection(connection string) (*gorm.DB, error) {
	connectionName := GetConnectionName(connection)
	driverName := GetDriverNameByConnectionName(connectionName)
	switch driverName {
	case "pgsql", "postgres":
		return postgreSQLSchemaDriver(connectionName)
	case "":
		return nil, fmt.Errorf("unsupported database connection: %v", connectionName)
	default:
		return nil, fmt.Errorf("unsupported database driver: %v", driverName)
	}
}

func GetConnectionName(connection string) string {
	connectionName := &connection
	if connection == "" {
		*connectionName = Config("database.default", "pgsql").(string)
	}

	return *connectionName
}

func GetDriverNameByConnectionName(connectionName string) string {
	return Config("database.connections."+connectionName+".driver", "").(string)
}

func GetDsn(connection string) string {
	host := Config("database.connections."+connection+".host", "127.0.0.1")
	port := Config("database.connections."+connection+".port", "5432")
	database := Config("database.connections."+connection+".database", "hitalent")
	username := Config("database.connections."+connection+".username", "root")
	password := Config("database.connections."+connection+".password", "")
	searchPath := Config("database.connections."+connection+".search_path", "public")
	sslmode := Config("database.connections."+connection+".sslmode", "prefer")

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=%s",
		host,
		port,
		username,
		password,
		database,
		searchPath,
		sslmode,
	)
}

func postgreSQLSchemaDriver(connection string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(GetDsn(connection)), &gorm.Config{})
}
