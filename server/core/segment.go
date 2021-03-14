package core

import (
	"database/sql"
	"log"
)

// TODO: normally we shouldn't export members of this domain object.
type Segment struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (s *Segment) Save(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO segments(name) VALUES($1) RETURNING id",
		s.Name).Scan(&s.Id)

	if err != nil {
		log.Fatal("Could not save segment ", s, err)
		return err
	}

	return nil
}
