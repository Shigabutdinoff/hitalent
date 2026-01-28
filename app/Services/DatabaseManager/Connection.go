package DatabaseManager

import (
	"hitalent/config"
	"hitalent/helpers"
)

// Connection Get a database connection instance./**
func Connection(connection string) string {
	if connection == "" {
		connectionName, isTrue := helpers.DataGet(config.Database, "default", "pgsql").(string)
		if isTrue {
			connection = connectionName
		}
	}
	driverName := helpers.DataGet(config.Database, "connections."+connection+".driver", nil)

	switch driverName {
	case "pgsql":
		return driverName.(string)
	default:
		return "No no no Mr. Fish"
	}
}
