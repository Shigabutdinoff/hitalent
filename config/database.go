package config

import . "hitalent/helpers"

var Database = map[string]any{

	/*
	   |--------------------------------------------------------------------------
	   | Default Database Connection Name
	   |--------------------------------------------------------------------------
	   |
	   | Here you may specify which of the database connections below you wish
	   | to use as your default connection for database operations. This is
	   | the connection which will be utilized unless another connection
	   | is explicitly specified when you execute a query / statement.
	   |
	*/

	"default": Env("DB_CONNECTION", "pgsql"),

	/*
	   |--------------------------------------------------------------------------
	   | Database Connections
	   |--------------------------------------------------------------------------
	   |
	   | Below are all of the database connections defined for your application.
	   | An example configuration is provided for each database system which
	   | is supported by GORM. You're free to add / remove connections.
	   |
	*/

	"connections": map[string]any{

		"pgsql": map[string]any{
			"driver":         "pgsql",
			"url":            Env("DB_URL", nil),
			"host":           Env("DB_HOST", "127.0.0.1"),
			"port":           Env("DB_PORT", "5432"),
			"database":       Env("DB_DATABASE", "hitalent"),
			"username":       Env("DB_USERNAME", "root"),
			"password":       Env("DB_PASSWORD", ""),
			"charset":        Env("DB_CHARSET", "utf8"),
			"prefix":         "",
			"prefix_indexes": true,
			"search_path":    "public",
			"sslmode":        Env("DB_SSLMODE", "prefer"),
		},
	},
}
