package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithValidData(t *testing.T) {
	// given
	clearTable()
	segmentName := "some-new-segment"
	var responseBody map[string]interface{}

	req, _ := http.NewRequest("POST", "/segments", bytes.NewBufferString(`{"name": "`+segmentName+`"}"`))

	// when we call the endpoint
	response := executeRequest(req)
	json.Unmarshal([]byte(response.Body.String()), &responseBody)

	// then it returns the newly created element
	assert.Equal(t, response.Code, http.StatusOK, "Response code should be 200/OK")
	assert.NotNil(t, responseBody)
	assert.Equal(t, responseBody["name"], segmentName, "Should return the given segment name")
	assert.Equal(t, responseBody["id"], 1)

	// when we get all the segments from the database
	segmentsFromDatabase := getSegments()

	// then there will be only our newly created element in it
	assert.Equal(t, len(segmentsFromDatabase), 1, "One record should be in the database")
	assert.Equal(t, segmentsFromDatabase[0]["name"], segmentName)
	assert.Equal(t, segmentsFromDatabase[0]["id"], 1)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

type M map[string]interface{}

func getSegments() []M {
	rows, _ := a.DB.Query("SELECT * FROM segments")
	defer rows.Close()
	var result []M

	for rows.Next() {
		var (
			id   int64
			name string
		)
		rows.Scan(&id, &name)

		var row M
		row["name"] = name
		row["id"] = id
		result = append(result, row)
	}

	return result
}
