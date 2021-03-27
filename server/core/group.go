package core

import (
	"database/sql"
	"time"
)

type Group struct {
	id            int
	start         time.Time
	parentSegment Segment
}

type GroupDto struct {
	Id        int       `json:"id"`
	Start     time.Time `json:"start"`
	SegmentId int       `json:"segmentId"`
}

func MakeGroupForSegment(segment Segment) Group {
	return Group{0, time.Time{}, segment}
}

func fetchGroupById(db *sql.DB, id int) (Group, error) {
	var group = Group{}
	var start sql.NullTime
	var parentSegmentId int

	if err := db.QueryRow("SELECT id, start, segment_id FROM groups WHERE id = $1", id).Scan(&group.id, &start, &parentSegmentId); err != nil {
		return group, err
	}

	if start.Valid {
		group.start = start.Time
	}

	segment, err := FetchSegmentById(db, parentSegmentId)

	if err != nil {
		return group, err
	}

	group.parentSegment = segment

	return group, nil
}

func (g Group) Save(db *sql.DB) (Group, error) {
	err := db.QueryRow(
		"INSERT INTO groups(id, segment_id) VALUES(DEFAULT, $1) RETURNING id",
		g.parentSegment.id).Scan(&g.id)

	return g, err
}

func (g Group) StartGroup(db *sql.DB) (Group, error) {
	err := db.QueryRow(
		"UPDATE groups SET start = $1 WHERE id = $2 RETURNING start",
		time.Now().UTC(), g.id).Scan(&g.start)

	return g, err
}

func (g *Group) Dto() GroupDto {
	return GroupDto{g.id, g.start, g.parentSegment.id}
}
