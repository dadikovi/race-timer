package core

import (
	"database/sql"
	"encoding/json"
	"log"
)

type Group struct {
	id            int64
	parentSegment Segment
}

type groupDto struct {
	Id        int64 `json:"id"`
	SegmentId int64 `json:"segmentId"`
}

func MakeGroupForSegment(segment Segment) Group {
	return Group{0, segment}
}

func (g *Group) Save(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO groups(segment_id) VALUES($1) RETURNING id",
		g.parentSegment.id).Scan(&g.id)

	if err != nil {
		log.Fatal("Could not save group ", g, err)
		return err
	}

	return nil
}

func (g *Group) ToJson() ([]byte, error) {
	groupDto := groupDto{g.id, g.parentSegment.id}
	j, err := json.Marshal(groupDto)
	if err != nil {
		return nil, err
	}
	return j, nil
}
