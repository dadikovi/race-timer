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

func TestPostParticipantsWithValidData(t *testing.T) {
	// given
	clearTable()
	createSegment("any-name")
	segmentId := 1

	// and
	startNumber := 1

	var responseBody map[string]interface{}

	// and group created (we expect with id 1)
	groupReq, _ := http.NewRequest("POST", "/groups", bytes.NewBufferString(`{"segmentId": `+fmt.Sprint(segmentId)+`}`))
	groupReq.Header.Set("Content-Type", "application/json")
	executeRequest(groupReq)
	groupId := 1

	req, _ := http.NewRequest("POST", "/participants", bytes.NewBufferString(`{"groupId": `+fmt.Sprint(groupId)+`, "startNumber": `+fmt.Sprint(startNumber)+`}`))
	req.Header.Set("Content-Type", "application/json")

	// when we call the endpoint
	response := executeRequest(req)
	json.Unmarshal([]byte(response.Body.String()), &responseBody)

	// then it returns the newly created element
	log.Print(response.Body.String())

	assert.Equal(t, http.StatusOK, response.Code, "Response code should be 200/OK")
	assert.NotNil(t, responseBody)
	assert.Equal(t, groupId, int(responseBody["groupId"].(float64)), "Should return the given group id")
	assert.Equal(t, startNumber, int(responseBody["startNumber"].(float64)))
	assert.Equal(t, -1, int(responseBody["raceTimeMs"].(float64))) // -1 raceTimeMs means that there is not a raceTime yet.

	// when we get all the participants from the database
	participantsFromDatabase := getParticipants()

	// then there will be only our newly created element in it
	assert.Equal(t, len(participantsFromDatabase), 1, "One record should be in the database")
	assert.Equal(t, int64(groupId), participantsFromDatabase[0]["group_id"])
	assert.Equal(t, int64(startNumber), participantsFromDatabase[0]["start_number"])
	assert.Equal(t, int64(-1), participantsFromDatabase[0]["race_time"])
}

func createGroup(segmentId int) {
	if _, err := a.DB.Exec(`INSERT INTO groups(id, segment_id) VALUES(DEFAULT, $1) RETURNING id`, segmentId); err != nil {
		log.Panic(err)
	}
}

func getParticipants() []RAWROW {
	rows, _ := a.DB.Query("SELECT group_id, start_number, race_time FROM participants")
	defer rows.Close()
	var result []RAWROW

	for rows.Next() {
		var (
			groupId     int64
			startNumber int64
			raceTime    int64
		)
		rows.Scan(&groupId, &startNumber, &raceTime)

		row := make(RAWROW)
		row["group_id"] = groupId
		row["start_number"] = startNumber
		row["race_time"] = raceTime
		result = append(result, row)
	}

	return result
}
