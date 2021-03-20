package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

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
	groupsFromDatabase := getGroupsFromDatabase()
	racesFromDatabase := getRacesFromDatabase()

	// then there will be only our newly created element in it
	assert.Equal(t, len(groupsFromDatabase), 1, "One record should be in the database")
	assert.Equal(t, int64(segmentId), groupsFromDatabase[0].segmentId)
	assert.Equal(t, int64(expectedGroupId), groupsFromDatabase[0].id)

	assert.Equal(t, len(racesFromDatabase), 1, "One record should be in the database")
	assert.Equal(t, int64(expectedGroupId), racesFromDatabase[0].activeGroupId)
}

type GroupDao struct {
	id        int64
	segmentId int64
	start     time.Time
}

func getGroupsFromDatabase() []GroupDao {

	rows, _ := a.DB.Query("SELECT id, segment_id, start FROM groups")
	defer rows.Close()
	var result []GroupDao

	for rows.Next() {
		var row GroupDao = GroupDao{}
		rows.Scan(&row.id, &row.segmentId, &row.start)
		result = append(result, row)
	}

	return result
}
