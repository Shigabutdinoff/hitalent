package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateMessagesTable, downCreateMessagesTable)
}

func upCreateMessagesTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.ExecContext(ctx, `
		create table messages (
			id bigserial primary key,
			chat_id bigint not null constraint messages_chat_id_foreign references chats on delete cascade,
			text varchar(255) not null,
			created_at timestamp (0)
		);
	`)
	return err
}

func downCreateMessagesTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.ExecContext(ctx, `drop table messages;`)
	return err
}
