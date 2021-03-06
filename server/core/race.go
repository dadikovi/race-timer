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
	LastRefresh time.Time
	ActiveGroup ActiveGroupResultsDto `json:"activeGroup"`
	Segments    []SegmentResultsDto   `json:"segments"`
}

type ActiveGroupResultsDto struct {
	Group        GroupDto         `json:"group"`
	Participants []ParticipantDto `json:"participants"`
}

type SegmentResultsDto struct {
	SegmentName string           `json:"segmentName"`
	List        []ParticipantDto `json:"participants"`
}

var RACE_RESULTS_CACHE_EVICTION_TIMEOUT = 5 * time.Second

func GetRaceInstance(db *sql.DB) (Race, error) {
	var instance = Race{}

	timer := startDbTimer("getRaceInstance")
	if err := db.QueryRow("SELECT active_group_id FROM races").Scan(&instance.activeGroup.id); err != nil {
		log.Print("Creating initial Race object")
		if err := db.QueryRow(
			"INSERT INTO races(active_group_id) VALUES(NULL) RETURNING active_group_id").Err(); err != nil {
			log.Panic(err)
			return instance, err
		}

		return instance, nil
	}
	timer.ObserveDuration()

	err := instance.refreshGroupData(db)
	return instance, err
}

func (r *Race) refreshGroupData(db *sql.DB) error {
	activeGroup, err := fetchGroupById(db, r.activeGroup.id)
	race := *r
	race.activeGroup = activeGroup
	*r = race
	return err
}

func (r Race) GetActiveGroup() Group {
	return r.activeGroup
}

func (r Race) SetActiveGroup(db *sql.DB, group Group) (Race, error) {
	timer := startDbTimer("setActiveGroup")
	if _, err := db.Exec(`UPDATE races SET active_group_id = $1`, group.id); err != nil {
		return r, err
	}
	timer.ObserveDuration()

	r.activeGroup = group
	return r, nil
}

func (r *Race) Results(db *sql.DB) (RaceResultsDto, error) {

	if r.GetActiveGroup().Dto().Id < 1 {
		return r.results, nil
	}

	err := r.refreshResultsIfNeeded(db)
	return r.results, err
}

func (r *Race) refreshResultsIfNeeded(db *sql.DB) error {
	if time.Now().UTC().Sub(r.results.LastRefresh) < RACE_RESULTS_CACHE_EVICTION_TIMEOUT {
		return nil
	}

	if err := r.refreshGroupData(db); err != nil {
		return err
	}

	if err := r.refreshActiveGroupStats(db); err != nil {
		return err
	}

	if err := r.refreshSegmentsGroupStats(db); err != nil {
		return err
	}

	race := *r
	race.results.LastRefresh = time.Now()
	*r = race

	return nil
}

func (r *Race) refreshSegmentsGroupStats(db *sql.DB) error {
	results := *&r.results
	results.Segments = make([]SegmentResultsDto, 0)

	rows, err := db.Query(`
	SELECT s.name, p.start_number, p.race_time
	FROM segments AS s
	JOIN groups AS g ON g.segment_id = s.id
	JOIN participants AS p ON p.group_id = g.id
	ORDER BY s.name ASC, p.race_time ASC
	`)

	if err != nil {
		return err
	}
	defer rows.Close()

	var lastSegment string
	var currentSegment SegmentResultsDto
	for rows.Next() {
		row := ParticipantDto{}
		rows.Scan(&currentSegment.SegmentName, &row.StartNumber, &row.RaceTimeMs)

		if currentSegment.SegmentName != lastSegment {
			if lastSegment != "" {
				results.Segments = append(results.Segments, currentSegment)
			}
			lastSegment = currentSegment.SegmentName
			currentSegment = SegmentResultsDto{SegmentName: lastSegment}
		}
		currentSegment.List = append(currentSegment.List, row)
	}
	results.Segments = append(results.Segments, currentSegment)

	*&r.results = results
	return nil
}

func (r *Race) refreshActiveGroupStats(db *sql.DB) error {
	results := *&r.results
	results.ActiveGroup.Group = r.activeGroup.Dto()
	results.ActiveGroup.Participants = make([]ParticipantDto, 0)

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
		results.ActiveGroup.Participants = append(results.ActiveGroup.Participants, row)
	}

	*&r.results = results
	return nil
}
