package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dadikovi/race-timer/server/core"
	"github.com/stretchr/testify/assert"
)

func TestPostSegmentsWithValidData(t *testing.T) {
	// given
	clearTable()
	segmentName := "some-new-segment"
	var createdSegment core.SegmentDto

	// when
	response := callCreateSegmentEndpoint(segmentName, &createdSegment)

	// then it returns the newly created element
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, segmentName, createdSegment.Name)
	assert.Equal(t, 1, int(createdSegment.Id))

	// when we get all the segments from the database
	segmentsFromDatabase := getSegmentsFromDatabase()

	// then there will be only our newly created element in it
	assert.Equal(t, len(segmentsFromDatabase), 1)
	assert.Equal(t, segmentName, segmentsFromDatabase[0].Name)
	assert.Equal(t, int(1), segmentsFromDatabase[0].Id)
}

func TestPostSegmentsWithExistingName(t *testing.T) {
	// given
	clearTable()
	segmentName := "some-new-segment"
	var createdSegment core.SegmentDto

	// and an endpoint with the name already exists
	callCreateSegmentEndpoint(segmentName, &createdSegment)

	// when
	response := callCreateSegmentEndpoint(segmentName, &createdSegment)

	// then it returns the newly created element
	assert.Equal(t, http.StatusBadRequest, response.Code)

	// when we get all the segments from the database
	segmentsFromDatabase := getSegmentsFromDatabase()

	// then
	assert.Equal(t, len(segmentsFromDatabase), 1, "One record should be in the database")
}

func TestGetSegmentsWithExistingData(t *testing.T) {
	// given
	clearTable()

	segmentName := "some-new-segment"
	createSegment(segmentName)

	var responseBody []core.SegmentDto

	req, _ := http.NewRequest("GET", "/segments", nil)
	req.Header.Set("Content-Type", "application/json")

	// when we call the endpoint
	response := executeRequest(req)
	json.Unmarshal([]byte(response.Body.String()), &responseBody)

	// then it returns the saved element
	assert.Equal(t, http.StatusOK, response.Code, "Response code should be 200/OK")
	assert.NotNil(t, responseBody)
	assert.Equal(t, segmentName, responseBody[0].Name, "Should return the given segment name")
	assert.Equal(t, 1, int(responseBody[0].Id))
}

func TestGetSegmentsWithEmptyDatabase(t *testing.T) {
	// given
	clearTable()

	var responseBody []core.SegmentDto

	req, _ := http.NewRequest("GET", "/segments", nil)
	req.Header.Set("Content-Type", "application/json")

	// when we call the endpoint
	response := executeRequest(req)
	json.Unmarshal([]byte(response.Body.String()), &responseBody)

	// then it returns the saved element
	assert.Equal(t, http.StatusOK, response.Code, "Response code should be 200/OK")
	assert.Nil(t, responseBody)
}

func callCreateSegmentEndpoint(segmentName string, responseDto interface{}) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("POST", "/segments", bytes.NewBufferString(`{"name": "`+segmentName+`"}`))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	json.Unmarshal([]byte(response.Body.String()), &responseDto)

	return response
}

func createSegment(segmentName string) {
	if _, err := a.DB.Exec(`INSERT INTO segments (id, name) VALUES (DEFAULT, $1)`, segmentName); err != nil {
		log.Panic(err)
	}

}

func getSegmentsFromDatabase() []core.SegmentDto {
	rows, _ := a.DB.Query("SELECT * FROM segments")
	defer rows.Close()
	var result []core.SegmentDto

	for rows.Next() {
		var row = core.SegmentDto{}
		rows.Scan(&row.Id, &row.Name)
		result = append(result, row)
	}

	return result
}
