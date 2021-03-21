package main

type RaceDao struct {
	activeGroupId int
}

func getRacesFromDatabase() []RaceDao {
	rows, _ := a.DB.Query("SELECT active_group_id FROM races")
	defer rows.Close()
	var result []RaceDao

	for rows.Next() {
		var row = RaceDao{}
		rows.Scan(&row.activeGroupId)
		result = append(result, row)
	}

	return result
}
