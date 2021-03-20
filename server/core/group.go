package core

import (
	"database/sql"
	"log"
	"time"
)

type Group struct {
	id            int
	parentSegment Segment
}

type GroupDto struct {
	Id        int `json:"id"`
	SegmentId int `json:"segmentId"`
}

func MakeGroupForSegment(segment Segment) Group {
	return Group{0, segment}
}

func fetchGroupById(db *sql.DB, id int) (Group, error) {
	var group = Group{}
	var parentSegmentId int
	if err := db.QueryRow("SELECT id, segment_id FROM groups WHERE id = $1", id).Scan(&group.id, &parentSegmentId); err != nil {
		return group, err
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

	if err != nil {
		log.Print("Could not save group ", g, err)
		return g, err
	}

	return g, nil
}

func (g Group) StartGroup(db *sql.DB) error {
	res, err := db.Exec(
		"UPDATE groups SET start = $1 WHERE id = $2",
		time.Now(), g.id)

	rows, _ := res.RowsAffected()

	log.Print("Updated rows: ", rows)

	return err
}

func (g *Group) Dto() GroupDto {
	return GroupDto{g.id, g.parentSegment.id}
}
