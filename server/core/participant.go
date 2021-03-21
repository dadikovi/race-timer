package core

import (
	"database/sql"
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

func (p Participant) Save(db *sql.DB) (Participant, error) {
	err := db.QueryRow(
		"INSERT INTO participants(start_number, group_id, race_time) VALUES($1, $2, $3) RETURNING start_number",
		p.startNumber, p.group.id, p.raceTimeMs).Scan(&p.startNumber)

	return p, err
}

func (p Participant) Finish(db *sql.DB) (Participant, error) {
	err := db.QueryRow(
		`UPDATE participants p SET p.race_time = $1
		WHERE p.start_number = $2
		RETURNING p.race_time`,
		p.group.start, p.startNumber).Scan(&p.raceTimeMs)

	return p, err
}

func (p *Participant) Dto() ParticipantDto {
	return ParticipantDto{p.startNumber, p.group.id, p.raceTimeMs}
}
