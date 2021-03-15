package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostGroupsWithValidData(t *testing.T) {
	// given
	clearTable()
	createSegment("any-name")
	segmentId := 1
	expectedGroupId := 1

	var responseBody map[string]interface{}

	req, _ := http.NewRequest("POST", "/groups", bytes.NewBufferString(`{"segmentId": `+fmt.Sprint(segmentId)+`}`))
	req.Header.Set("Content-Type", "application/json")

	// when we call the endpoint
	response := executeRequest(req)
	json.Unmarshal([]byte(response.Body.String()), &responseBody)

	// then it returns the newly created element
	log.Print(response.Body.String())

	assert.Equal(t, http.StatusOK, response.Code, "Response code should be 200/OK")
	assert.NotNil(t, responseBody)
	assert.Equal(t, segmentId, int(responseBody["segmentId"].(float64)), "Should return the given segment name")
	assert.Equal(t, expectedGroupId, int(responseBody["id"].(float64)))

	// when we get all the segments from the database
	groupsFromDatabase := getGroups()
	racesFromDatabase := getRaces()

	// then there will be only our newly created element in it
	assert.Equal(t, len(groupsFromDatabase), 1, "One record should be in the database")
	assert.Equal(t, int64(segmentId), groupsFromDatabase[0]["segment_id"])
	assert.Equal(t, int64(expectedGroupId), groupsFromDatabase[0]["id"])

	assert.Equal(t, len(racesFromDatabase), 1, "One record should be in the database")
	assert.Equal(t, int64(expectedGroupId), racesFromDatabase[0]["active_group_id"])
}

func getGroups() []RAWROW {
	rows, _ := a.DB.Query("SELECT id, segment_id FROM groups")
	defer rows.Close()
	var result []RAWROW

	for rows.Next() {
		var (
			id         int64
			segment_id int64
		)
		rows.Scan(&id, &segment_id)

		log.Print("FOUND ROW, YO", id, segment_id)

		row := make(RAWROW)
		row["segment_id"] = segment_id
		row["id"] = id
		result = append(result, row)
	}

	return result
}

func getRaces() []RAWROW {
	rows, _ := a.DB.Query("SELECT active_group_id FROM races")
	defer rows.Close()
	var result []RAWROW

	for rows.Next() {
		var (
			active_group_id int64
		)
		rows.Scan(&active_group_id)

		row := make(RAWROW)
		row["active_group_id"] = active_group_id
		result = append(result, row)
	}

	return result
}
