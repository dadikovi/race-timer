package core

import "database/sql"

type Race struct {
	activeGroup Group
}

func GetRaceInstance(db *sql.DB) (Race, error) {
	var instance = Race{}
	var activeGroupId int64

	if err := db.QueryRow("SELECT active_group_id FROM races").Scan(&activeGroupId); err != nil {
		var defaultActiveGroupId int64 // Only for checking the save result
		if err := db.QueryRow(
			"INSERT INTO races(active_group_id) VALUES($1) RETURNING active_group_id",
			0).Scan(&defaultActiveGroupId); err != nil {
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
