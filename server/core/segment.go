package core

import (
	"database/sql"
	"encoding/json"
	"log"
)

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

func FetchSegmentById(db *sql.DB, id int64) (Segment, error) {
	var segment = Segment{}

	if err := db.QueryRow("SELECT id, name FROM segments WHERE id = $1", id).Scan(&segment.id, &segment.name); err != nil {
		return segment, err
	}
	return segment, nil
}

func FetchAll(db *sql.DB) ([]Segment, error) {
	var result []Segment
	var rows, err = db.Query("SELECT id, name FROM segments")

	if err != nil {
		log.Fatal("Could not get segments", err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var id int64
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal("Could not get segments", err)
			return nil, err
		}
		result = append(result, Segment{id, name})
	}
	// get any error encountered during iteration

	if err := rows.Err(); err != nil {
		log.Fatal("Could not get segments", err)
		return nil, err
	}

	return result, nil
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
		"INSERT INTO segments(id, name) VALUES(DEFAULT, $1) RETURNING id",
		s.name).Scan(&s.id)

	if err != nil {
		log.Print("Could not save segment ", s, err)
		return err
	}

	return nil
}
