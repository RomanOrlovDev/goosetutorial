package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
	"log"
)

func init() {
	goose.AddMigration(upAddSexColumnUserTable, downAddSexColumnUserTable)
}

func upAddSexColumnUserTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
CREATE TYPE sex as ENUM('m', 'w');
ALTER TABLE users ADD COLUMN sex sex;
`)
	if err != nil {
		log.Fatalln("error occurred while upping a migration:", err)
	}
	return nil
}

func downAddSexColumnUserTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
DROP TABLE IF EXISTS sex;
DROP TYPE IF EXISTS sex;
`)
	if err != nil {
		log.Fatal("error occurred while downing a migration:", err)
	}
	return nil
}
