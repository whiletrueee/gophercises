// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package pg_db

import (
	"database/sql"
)

type UrlShortener struct {
	ID          string
	OriginalUrl string
	Shortkey    string
	HitCount    sql.NullInt32
	CreatedAt   sql.NullTime
}
