package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStartActiveGroup(t *testing.T) {
	// given
	clearTable()
	createSegment("any-name")
	segmentId := 1

	// and group created (we expect with id 1)
	groupReq, _ := http.NewRequest("POST", "/groups", bytes.NewBufferString(`{"segmentId": `+fmt.Sprint(segmentId)+`}`))
	groupReq.Header.Set("Content-Type", "application/json")
	executeRequest(groupReq)

	// when active group start endpoint is called
	req, _ := http.NewRequest("PUT", "/groups/active", bytes.NewBufferString(``))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)

	// then
	log.Print(response.Body.String())

	assert.Equal(t, http.StatusOK, response.Code, "Response code should be 200/OK")

	// when we get all the groups from the database
	groupsFromDatabase := getGroupsFromDatabase()

	// then there will be only our newly created element in it
	assert.Equal(t, len(groupsFromDatabase), 1, "One record should be in the database")
	assert.True(t, time.Now().Sub(groupsFromDatabase[0].start) < 1000)
}

type RaceDao struct {
	activeGroupId int64
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
