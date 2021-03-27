package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dadikovi/race-timer/server/core"
	"github.com/stretchr/testify/assert"
)

func TestPostGroupsWithValidData(t *testing.T) {
	// given
	clearTable()
	createSegment("any-name")
	segmentId := 1
	expectedGroupId := 1

	// when we call the endpoint
	var createdGroup core.GroupDto
	response := callCreateGroupEndpoint(segmentId, &createdGroup)

	// then it returns the newly created element
	log.Print(response.Body.String())

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, segmentId, createdGroup.SegmentId)
	assert.Equal(t, expectedGroupId, createdGroup.Id)

	// when we get all the segments from the database
	groupsFromDatabase := getGroupsFromDatabase()
	racesFromDatabase := getRacesFromDatabase()

	// then there will be only our newly created element in it
	assert.Equal(t, len(groupsFromDatabase), 1)
	assert.Equal(t, segmentId, groupsFromDatabase[0].segmentId)
	assert.Equal(t, expectedGroupId, groupsFromDatabase[0].id)

	assert.Equal(t, len(racesFromDatabase), 1)
	assert.Equal(t, expectedGroupId, racesFromDatabase[0].activeGroupId)
}

func TestStartActiveGroup(t *testing.T) {

	// given
	var createdSegment core.SegmentDto
	var createdGroup core.GroupDto

	// and
	clearTable()
	callCreateSegmentEndpoint("any-name", &createdSegment)
	callCreateGroupEndpoint(int(createdSegment.Id), &createdGroup)

	// when active group start endpoint is called
	response := callStartActiveGroupEndpoint()
	log.Print(response.Body.String())

	// then
	assert.Equal(t, http.StatusOK, response.Code)

	// when we get all the groups from the database
	groupsFromDatabase := getGroupsFromDatabase()

	// then there will be only our newly created element in it
	assert.Equal(t, len(groupsFromDatabase), 1)
	assert.True(t, time.Now().UTC().Sub(groupsFromDatabase[0].start) < 10000000) // 10ms difference
}

type GroupDao struct {
	id        int
	segmentId int
	start     time.Time
}

func callStartActiveGroupEndpoint() *httptest.ResponseRecorder {
	req, _ := http.NewRequest("POST", "/groups/active", bytes.NewBufferString(``))
	req.Header.Set("Content-Type", "application/json")
	return executeRequest(req)
}

func callCreateGroupEndpoint(segmentId int, responseDto interface{}) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("POST", "/groups", bytes.NewBufferString(`{"segmentId": `+fmt.Sprint(segmentId)+`}`))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	json.Unmarshal([]byte(response.Body.String()), &responseDto)

	return response
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
