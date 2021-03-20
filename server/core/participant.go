package core

import (
	"database/sql"
	"log"
)

type Participant struct {
	startNumber int64
	group       Group
	raceTimeMs  int64
}

type ParticipantDto struct {
	StartNumber int64 `json:"startNumber"`
	GroupId     int64 `json:"groupId"`
	RaceTimeMs  int64 `json:"raceTimeMs"`
}

func MakeParticipantForGroup(startNumber int64, group Group) Participant {
	return Participant{startNumber, group, -1}
}

func (p Participant) Save(db *sql.DB) (Participant, error) {
	err := db.QueryRow(
		"INSERT INTO participants(start_number, group_id, race_time) VALUES($1, $2, $3) RETURNING start_number",
		p.startNumber, p.group.id, p.raceTimeMs).Scan(&p.startNumber)

	if err != nil {
		log.Print("Could not save participant ", p, err)
		return p, err
	}

	return p, nil
}

func (p *Participant) Dto() ParticipantDto {
	return ParticipantDto{p.startNumber, p.group.id, p.raceTimeMs}
}
