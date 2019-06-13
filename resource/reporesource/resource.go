package reporesource

import (
	"database/sql"
)

type RepoResource struct {
	Db *sql.DB
}
