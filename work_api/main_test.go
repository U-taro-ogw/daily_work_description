package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/U-taro-ogw/daily_work_description/work_api/models"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/work_records", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != `{"work_records":[]}` {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNotExistsWorkRecord(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/work_records/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Not Found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'WorkRecord not found'. Got '%s'", m["error"])
	}
}

func TestCreateWorkRecord(t *testing.T) {
	clearTable()
	param := `
{
"work_date": "2014-10-10T00:00:00+09:00",
"begin_work_time": "2014-10-10T10:00:00+09:00",
"end_work_time": "2014-10-10T19:00:00+09:00",
"begin_break_time": "2014-10-10T12:00:00+09:00",
"end_break_time": "2014-10-10T13:00:00+09:00"
}
`
	payload := []byte(param)
	req, _ := http.NewRequest("POST", "/work_records", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]models.WorkRecord
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["work_record"].WorkDate.Format(time.RFC3339) != "2014-10-10T00:00:00+09:00" {
		t.Errorf("Expected work_date to be '2014-10-10T00:00:00+09:00'. Got '%v'", m["work_record"].WorkDate.Format(time.RFC3339))
	}

	if m["work_record"].BeginWorkTime.Format(time.RFC3339) != "2014-10-10T10:00:00+09:00" {
		t.Errorf("Expected begin_work_time to be '2014-10-10T10:00:00+09:00'. Got '%v'", m["work_record"].BeginWorkTime.Format(time.RFC3339))
	}

	if m["work_record"].EndWorkTime.Format(time.RFC3339) != "2014-10-10T19:00:00+09:00" {
		t.Errorf("Expected end_work_time to be '2014-10-10T19:00:00+09:00'. Got '%v'", m["work_record"].EndWorkTime.Format(time.RFC3339))
	}

	if m["work_record"].BeginBreakTime.Format(time.RFC3339) != "2014-10-10T12:00:00+09:00" {
		t.Errorf("Expected begin_break_time to be '2014-10-10T12:00:00+09:00'. Got '%v'", m["work_record"].BeginBreakTime.Format(time.RFC3339))
	}

	if m["work_record"].EndBreakTime.Format(time.RFC3339) != "2014-10-10T13:00:00+09:00" {
		t.Errorf("Expected end_break_time to be '2014-10-10T13:00:00+09:00'. Got '%v'", m["work_record"].EndBreakTime.Format(time.RFC3339))
	}
}

func TestGetWorkRecord(t *testing.T) {
	clearTable()
	addWorkRecord(1)

	var work models.WorkRecord
	a.DB.First(&work)
	path := `/work_records/` + fmt.Sprint(work.ID)

	req, _ := http.NewRequest("GET", path, nil)

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]models.WorkRecord
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["work_record"].WorkDate.Format(time.RFC3339) != work.WorkDate.Format(time.RFC3339) {
		t.Errorf(
			"Expected work_date to be '%v'. Got '%v'",
			work.WorkDate.Format(time.RFC3339),
			m["work_record"].WorkDate.Format(time.RFC3339),
		)
	}

	if m["work_record"].BeginWorkTime.Format(time.RFC3339) != work.BeginWorkTime.Format(time.RFC3339) {
		t.Errorf(
			"Expected begin_work_time to be '%v'. Got '%v'",
			work.BeginWorkTime.Format(time.RFC3339),
			m["work_record"].BeginWorkTime.Format(time.RFC3339),
		)
	}

	if m["work_record"].EndWorkTime.Format(time.RFC3339) != work.EndWorkTime.Format(time.RFC3339) {
		t.Errorf(
			"Expected end_work_time to be '%v'. Got '%v'",
			work.EndWorkTime.Format(time.RFC3339),
			m["work_record"].EndWorkTime.Format(time.RFC3339),
		)
	}

	if m["work_record"].BeginBreakTime.Format(time.RFC3339) != work.BeginBreakTime.Format(time.RFC3339) {
		t.Errorf(
			"Expected begin_break_time to be '%v'. Got '%v'",
			work.BeginBreakTime.Format(time.RFC3339),
			m["work_record"].BeginBreakTime.Format(time.RFC3339),
		)
	}

	if m["work_record"].EndBreakTime.Format(time.RFC3339) != work.EndBreakTime.Format(time.RFC3339) {
		t.Errorf(
			"Expected end_break_time to be '%v'. Got '%v'",
			work.EndBreakTime.Format(time.RFC3339),
			m["work_record"].EndBreakTime.Format(time.RFC3339),
		)
	}
}

func TestUpdateWorkRecord(t *testing.T) {
	clearTable()
	addWorkRecord(1)

	var work models.WorkRecord
	a.DB.First(&work)
	path := `/work_records/` + fmt.Sprint(work.ID)

	req, _ := http.NewRequest("GET", path, nil)
	response := executeRequest(req)
	var before map[string]models.WorkRecord
	json.Unmarshal(response.Body.Bytes(), &before)
	beforeWorkRecord := before["work_record"]

	param := `
{
"work_date": "2016-10-10T00:00:00+09:00",
"begin_work_time": "2017-10-10T10:00:00+09:00",
"end_work_time": "2018-10-10T19:00:00+09:00",
"begin_break_time": "2019-10-10T12:00:00+09:00",
"end_break_time": "2020-10-10T13:00:00+09:00"
}
`
	payload := []byte(param)
	req, _ = http.NewRequest("PUT", path, bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusNoContent, response.Code)

	if response.Header()["Location"][0] != "http://localhost:8080/" + fmt.Sprint(work.ID) {
		t.Errorf(
			"Expected the Location Haader to remain the same (%v). Got %v",
			"http://localhost:8080/" + fmt.Sprint(work.ID),
			response.Header()["Location"][0],
		)
	}

	var afterWorkRecord models.WorkRecord
	a.DB.First(&afterWorkRecord, beforeWorkRecord.ID)

	if afterWorkRecord.ID != beforeWorkRecord.ID {
		t.Errorf(
			"Expected the id to remain the same (%v). Got %v",
			beforeWorkRecord.ID,
			afterWorkRecord.ID,
		)
	}

	if afterWorkRecord.WorkDate.Format(time.RFC3339) != "2016-10-10T00:00:00+09:00" {
		t.Errorf(
			"Expected the WorkDate to Update Failed 2016-10-10T00:00:00+09:00. Got %v",
			afterWorkRecord.WorkDate.Format(time.RFC3339),
		)
	}

	if afterWorkRecord.BeginWorkTime.Format(time.RFC3339) != "2017-10-10T10:00:00+09:00" {
		t.Errorf(
			"Expected the BeginWorkTime to Update Failed 2017-10-10T10:00:00+09:00. Got %v",
			afterWorkRecord.BeginWorkTime.Format(time.RFC3339),
		)
	}

	if afterWorkRecord.EndWorkTime.Format(time.RFC3339) != "2018-10-10T19:00:00+09:00" {
		t.Errorf(
			"Expected the EndWorkTime to Update Failed 2018-10-10T19:00:00+09:00. Got %v",
			afterWorkRecord.EndWorkTime.Format(time.RFC3339),
		)
	}

	if afterWorkRecord.BeginBreakTime.Format(time.RFC3339) != "2019-10-10T12:00:00+09:00" {
		t.Errorf(
			"Expected the BeginBreakTime to Update Failed 2019-10-10T12:00:00+09:00. Got %v",
			afterWorkRecord.BeginBreakTime.Format(time.RFC3339),
		)
	}

	if afterWorkRecord.EndBreakTime.Format(time.RFC3339) != "2020-10-10T13:00:00+09:00" {
		t.Errorf(
			"Expected the EndBreakTime to Update Failed 2020-10-10T13:00:00+09:00. Got %v",
			afterWorkRecord.EndBreakTime.Format(time.RFC3339),
		)
	}
}

//func TestDeleteWorkRecord(t *testing.T) {
//	clearTable()
//	addWorkRecord(1)
//
//	var work models.WorkRecord
//	a.DB.First(&work)
//	path := `/work_records/` + fmt.Sprint(work.ID)
//
//	req, _ := http.NewRequest("GET", path, nil)
//	response := executeRequest(req)
//	checkResponseCode(t, http.StatusOK, response.Code)
//
//	req, _ = http.NewRequest("DELETE", path, nil)
//	response = executeRequest(req)
//	checkResponseCode(t, http.StatusOK, response.Code)
//
//	req, _ = http.NewRequest("GET", path, nil)
//	response = executeRequest(req)
//	checkResponseCode(t, http.StatusNotFound, response.Code)
//}

func addWorkRecord(count int) {
	var work models.WorkRecord
	if count < 1 {
		count = 1
	}

	t := time.Now()
	for i := 0; i < count; i++ {
		t2 := t.AddDate(0, 0, -i)

		work.WorkDate = t2
		work.BeginWorkTime = t2
		work.EndWorkTime = t2
		work.BeginBreakTime = t2
		work.EndBreakTime = t2

		a.DB.Create(&work)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM work_records")
}
