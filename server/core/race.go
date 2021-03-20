package core

import (
	"database/sql"
	"log"
)

type Race struct {
	activeGroup Group
}

func GetRaceInstance(db *sql.DB) (Race, error) {
	var instance = Race{}
	var activeGroupId int

	if err := db.QueryRow("SELECT active_group_id FROM races").Scan(&activeGroupId); err != nil {
		if err := db.QueryRow(
			"INSERT INTO races(active_group_id) VALUES(NULL) RETURNING active_group_id").Err(); err != nil {
			log.Panic(err)
			return instance, err
		}

		return instance, nil
	}

	activeGroup, err := fetchGroupById(db, activeGroupId)

	if err != nil {
		return instance, err
	}

	instance.activeGroup = activeGroup
	return instance, nil
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
