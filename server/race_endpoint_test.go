package main

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"github.com/dadikovi/race-timer/server/core"
	"github.com/stretchr/testify/assert"
)

func TestStartActiveGroup(t *testing.T) {

	// given
	var createdSegment core.SegmentDto
	var createdGroup core.GroupDto

	// and
	clearTable()
	callCreateSegmentEndpoint("any-name", &createdSegment)
	callCreateGroupEndpoint(int(createdSegment.Id), &createdGroup)

	// when active group start endpoint is called
	req, _ := http.NewRequest("PUT", "/groups/active", bytes.NewBufferString(``))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)

	// then
	assert.Equal(t, http.StatusOK, response.Code)

	// when we get all the groups from the database
	groupsFromDatabase := getGroupsFromDatabase()

	// then there will be only our newly created element in it
	assert.Equal(t, len(groupsFromDatabase), 1)
	assert.True(t, time.Now().Sub(groupsFromDatabase[0].start) < 1000)
}

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
