package core

import (
	"database/sql"
	"log"
	"time"
)

type Race struct {
	activeGroup Group
	results     RaceResultsDto
}

type RaceResultsDto struct {
	lastRefresh time.Time
	activeGroup []ParticipantDto
	segments    []SegmentResultsDto
}

type SegmentResultsDto struct {
	segmentName string
	list        []ParticipantDto
}

var CACHE_EVICTION_TIMEOUT = 5 * time.Second

func GetRaceInstance(db *sql.DB) (Race, error) {
	var instance = Race{}

	if err := db.QueryRow("SELECT active_group_id FROM races").Scan(&instance.activeGroup.id); err != nil {
		if err := db.QueryRow(
			"INSERT INTO races(active_group_id) VALUES(NULL) RETURNING active_group_id").Err(); err != nil {
			log.Panic(err)
			return instance, err
		}

		return instance, nil
	}

	activeGroup, err := fetchGroupById(db, instance.activeGroup.id)
	instance.activeGroup = activeGroup
	return instance, err
}

func (r Race) GetActiveGroup() Group {
	return r.activeGroup
}

func (r Race) SetActiveGroup(db *sql.DB, group Group) (Race, error) {
	if _, err := db.Exec(`UPDATE races SET active_group_id = $1`, group.id); err != nil {
		return r, err
	}

	r.activeGroup = group
	return r, nil
}

func (r Race) Results(db *sql.DB) (RaceResultsDto, error) {
	err := r.refreshResultsIfNeeded(db)
	return r.results, err
}

func (r Race) refreshResultsIfNeeded(db *sql.DB) error {
	if time.Now().UTC().Sub(r.results.lastRefresh) < CACHE_EVICTION_TIMEOUT {
		return nil
	}

	r.results.segments = nil

	if err := r.refreshActiveGroupStats(db); err != nil {
		return err
	}

	if err := r.refreshSegmentsGroupStats(db); err != nil {
		return err
	}

	r.results.lastRefresh = time.Now()

	return nil
}

func (r Race) refreshSegmentsGroupStats(db *sql.DB) error {
	r.results.activeGroup = nil

	rows, err := db.Query(`
	SELECT s.name, p.start_number, p.race_time
	FROM segments AS s
	JOIN groups AS g ON g.segment_id = s.id
	JOIN participants AS p ON p.group_id = g.id
	ORDER BY s.name ASC, p.race_time DESC
	`, r.activeGroup.id)

	if err != nil {
		return err
	}
	defer rows.Close()

	var lastSegment string
	var currentSegment SegmentResultsDto
	for rows.Next() {
		row := ParticipantDto{}
		rows.Scan(&currentSegment, &row.StartNumber, &row.RaceTimeMs)

		if currentSegment.segmentName != lastSegment {
			lastSegment = currentSegment.segmentName
			r.results.segments = append(r.results.segments, currentSegment)
			currentSegment = SegmentResultsDto{}
		}
		currentSegment.list = append(currentSegment.list, row)
	}
	r.results.segments = append(r.results.segments, currentSegment)

	return nil
}

func (r Race) refreshActiveGroupStats(db *sql.DB) error {
	r.results.segments = nil

	rows, err := db.Query(`
		SELECT start_number, race_time
		FROM participants
		WHERE group_id = $1
	`, r.activeGroup.id)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		row := ParticipantDto{}
		rows.Scan(&row.StartNumber, &row.RaceTimeMs)
		r.results.activeGroup = append(r.results.activeGroup, row)
	}

	return nil
}
