package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dadikovi/race-timer/server/core"
	"github.com/stretchr/testify/assert"
)

func TestPostParticipantsWithValidData(t *testing.T) {
	// given
	var createdSegment core.SegmentDto
	var createdGroup core.GroupDto
	var registeredParticipant core.ParticipantDto
	startNumber := 1

	// and
	clearTable()
	callCreateSegmentEndpoint("any-name", &createdSegment)
	callCreateGroupEndpoint(int(createdSegment.Id), &createdGroup)

	// when
	response := callRegisterParticipantEndpoint(createdGroup.Id, startNumber, &registeredParticipant)

	// then
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, createdGroup.Id, registeredParticipant.GroupId)
	assert.Equal(t, startNumber, registeredParticipant.StartNumber)
	assert.Equal(t, -1, registeredParticipant.RaceTimeMs) // -1 raceTimeMs means that there is not a raceTime yet.

	// when we get all the participants from the database
	participantsFromDatabase := getParticipants()

	// then there will be only our newly created element in it
	assert.Equal(t, len(participantsFromDatabase), 1)
	assert.Equal(t, createdGroup.Id, participantsFromDatabase[0].groupId)
	assert.Equal(t, startNumber, participantsFromDatabase[0].startNumber)
	assert.Equal(t, -1, participantsFromDatabase[0].raceTime)
}

func createGroup(segmentId int) {
	if _, err := a.DB.Exec(`INSERT INTO groups(id, segment_id) VALUES(DEFAULT, $1) RETURNING id`, segmentId); err != nil {
		log.Panic(err)
	}
}

type ParticipantDao struct {
	groupId     int
	startNumber int
	raceTime    int
}

func callRegisterParticipantEndpoint(groupId int, startNumber int, responseDto interface{}) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("POST", "/participants", bytes.NewBufferString(`{"groupId": `+fmt.Sprint(groupId)+`, "startNumber": `+fmt.Sprint(startNumber)+`}`))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	json.Unmarshal([]byte(response.Body.String()), &responseDto)

	return response
}

func getParticipants() []ParticipantDao {
	rows, _ := a.DB.Query("SELECT group_id, start_number, race_time FROM participants")
	defer rows.Close()
	var result []ParticipantDao

	for rows.Next() {
		var row = ParticipantDao{}
		rows.Scan(&row.groupId, &row.startNumber, &row.raceTime)
		result = append(result, row)
	}

	return result
}
