package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateChatsTable, downCreateChatsTable)
}

func upCreateChatsTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.ExecContext(ctx, `
		create table chats (
			id bigserial primary key,
			title varchar(255) not null,
			created_at timestamp (0)
		);
	`)
	return err
}

func downCreateChatsTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.ExecContext(ctx, `drop table chats;`)
	return err
}
