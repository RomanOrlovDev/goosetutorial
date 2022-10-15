package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
	"log"
)

func init() {
	goose.AddMigration(upInitUserTable, downInitUserTable)
}

func upInitUserTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
CREATE TABLE IF NOT EXISTS users (
    "id" SERIAL PRIMARY KEY NOT NULL,
	"fullname" VARCHAR(64) not null
)
`)
	if err != nil {
		log.Fatalln("error occurred while upping a migration:", err)
	}
	return nil
}

func downInitUserTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
DROP TABLE IF EXISTS users;
`)
	if err != nil {
		log.Fatal("error occurred while downing a migration:", err)
	}
	return nil
}
