package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dadikovi/race-timer/server/core"
	"github.com/stretchr/testify/assert"
)

func TestResultsEndpoint(t *testing.T) {
	// given
	segmentName := "any-name"
	var createdSegment core.SegmentDto
	var createdGroup core.GroupDto
	var participant core.ParticipantDto
	var results core.RaceResultsDto

	// and
	clearTable()
	callCreateSegmentEndpoint(segmentName, &createdSegment)
	callCreateGroupEndpoint(int(createdSegment.Id), &createdGroup)

	// and we register 3 participants in DESC order
	callRegisterParticipantEndpoint(createdGroup.Id, 3, &participant)
	callRegisterParticipantEndpoint(createdGroup.Id, 2, &participant)
	callRegisterParticipantEndpoint(createdGroup.Id, 1, &participant)
	callStartActiveGroupEndpoint()

	// and run forest run
	callParticipantFinishedEndpoint(1, &participant)
	time.Sleep(time.Second)
	callParticipantFinishedEndpoint(2, &participant)
	time.Sleep(time.Second)
	callParticipantFinishedEndpoint(3, &participant)

	// when
	response := callResultsEndpoint(&results)
	log.Print(response.Body.String())

	// then
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, len(results.ActiveGroup), 3)
	assert.Equal(t, len(results.Segments), 1)
	assert.Equal(t, results.Segments[0].SegmentName, segmentName)
	assert.Equal(t, len(results.Segments[0].List), 3)
	assert.Equal(t, results.Segments[0].List[0].StartNumber, 1)
	assert.Equal(t, results.Segments[0].List[1].StartNumber, 2)
	assert.Equal(t, results.Segments[0].List[2].StartNumber, 3)
}

func callResultsEndpoint(responseDto interface{}) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", `/race/results`, bytes.NewBufferString(`{}`))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	json.Unmarshal([]byte(response.Body.String()), &responseDto)

	return response
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
