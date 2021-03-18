package core

import (
	"database/sql"
	"encoding/json"
	"log"
)

type Participant struct {
	startNumber int64
	group       Group
	raceTimeMs  int64
}

type participantDto struct {
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

func (p *Participant) ToJson() ([]byte, error) {
	participantDto := participantDto{p.startNumber, p.group.id, p.raceTimeMs}
	j, err := json.Marshal(participantDto)
	if err != nil {
		return nil, err
	}
	return j, nil
}
