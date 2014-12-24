package eccore

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// NewDB will generate a new sql.DB pool from the
// flags.
func NewDB() (*sql.DB, error) {
	connectString := fmt.Sprintf("postgres://%s:%s@%s/%s", psqlUser, psqlPass, psqlHost, psqlDb)
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
