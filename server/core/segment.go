package core

import (
	"database/sql"
	"encoding/json"
	"log"
)

// TODO: normally we shouldn't export members of this domain object.
type Segment struct {
	id   int64
	name string
}

type segmentDto struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func MakeSegment(jsonRepresentation string) (Segment, error) {
	var segment Segment
	var dto segmentDto
	if err := json.Unmarshal([]byte(jsonRepresentation), &dto); err != nil {
		return segment, err
	}

	segment = Segment{dto.Id, dto.Name}
	return segment, nil
}

func (s *Segment) ToJson() ([]byte, error) {
	segmentDto := segmentDto{s.id, s.name}
	j, err := json.Marshal(segmentDto)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (s *Segment) Save(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO segments(name) VALUES($1) RETURNING id",
		s.name).Scan(&s.id)

	if err != nil {
		log.Fatal("Could not save segment ", s, err)
		return err
	}

	return nil
}
