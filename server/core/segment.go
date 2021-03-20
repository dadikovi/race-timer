package core

import (
	"database/sql"
	"errors"
	"log"
	"strings"
)

type Segment struct {
	id   int64
	name string
}

type SegmentDto struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

var ALREADY_EXISTS_ERROR_CODE = "ALREADY_EXISTS"

func MakeSegment(name string) Segment {
	return Segment{0, name}
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

func (s Segment) Save(db *sql.DB) (Segment, error) {
	err := db.QueryRow(
		"INSERT INTO segments(id, name) VALUES(DEFAULT, $1) RETURNING id",
		s.name).Scan(&s.id)

	if err != nil {
		if strings.Contains(err.Error(), "segments_name_key") {
			return s, errors.New(ALREADY_EXISTS_ERROR_CODE)
		}

		log.Print("Could not save segment ", s, err)
		return s, err
	}

	return s, nil
}

func (s *Segment) Dto() SegmentDto {
	return SegmentDto{s.id, s.name}
}
