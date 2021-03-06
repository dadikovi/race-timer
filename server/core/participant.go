package core

import (
	"database/sql"
	"time"
)

type Participant struct {
	startNumber int
	group       Group
	raceTimeMs  int
}

type ParticipantDto struct {
	StartNumber int `json:"startNumber"`
	GroupId     int `json:"groupId"`
	RaceTimeMs  int `json:"raceTimeMs"`
}

func MakeParticipantForGroup(startNumber int, group Group) Participant {
	return Participant{startNumber, group, -1}
}

func FetchParticipantByStartNumber(db *sql.DB, startNumber int) (Participant, error) {
	participant := Participant{}
	var groupId int

	timer := startDbTimer("FetchParticipantByStartNumber")
	if err := db.QueryRow(
		"SELECT start_number, group_id, race_time FROM participants WHERE start_number = $1",
		startNumber).Scan(&participant.startNumber, &groupId, &participant.raceTimeMs); err != nil {
		return participant, err
	}
	timer.ObserveDuration()

	if group, err := fetchGroupById(db, groupId); err != nil {
		return participant, err
	} else {
		participant.group = group
	}

	return participant, nil
}

func (p Participant) Save(db *sql.DB) (Participant, error) {
	timer := startDbTimer("saveParticipant")
	err := db.QueryRow(
		"INSERT INTO participants(start_number, group_id, race_time) VALUES($1, $2, $3) RETURNING start_number",
		p.startNumber, p.group.id, p.raceTimeMs).Scan(&p.startNumber)
	timer.ObserveDuration()

	return p, err
}

func (p Participant) Finish(db *sql.DB) (Participant, error) {
	timer := startDbTimer("finishParticipant")
	err := db.QueryRow(
		`UPDATE participants SET race_time = $1
		WHERE start_number = $2
		RETURNING race_time`,
		time.Now().UTC().Sub(p.group.start).Milliseconds(), p.startNumber).Scan(&p.raceTimeMs)
	timer.ObserveDuration()

	return p, err
}

func (p *Participant) Dto() ParticipantDto {
	return ParticipantDto{p.startNumber, p.group.id, p.raceTimeMs}
}
