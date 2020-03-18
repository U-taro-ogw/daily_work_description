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

	if body := response.Body.String(); body != "[]" {
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
"end_break_time": "2014-10-10T13:00:00+09:00",
}
`
	payload := []byte(param)
	req, _ := http.NewRequest("POST", "/work_records", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["work_date"] != "2014-10-10T00:00:00+09:00" {
		t.Errorf("Expected work_date to be '2014-10-10T00:00:00+09:00'. Got '%v'", m["work_date"])
	}

	if m["begin_work_time"] != "2014-10-10T10:00:00+09:00" {
		t.Errorf("Expected begin_work_time to be '2014-10-10T10:00:00+09:00'. Got '%v'", m["begin_work_time"])
	}

	if m["end_work_time"] != "2014-10-10T19:00:00+09:00" {
		t.Errorf("Expected end_work_time to be '2014-10-10T19:00:00+09:00'. Got '%v'", m["end_work_time"])
	}

	if m["begin_break_time"] != "2014-10-10T12:00:00+09:00" {
		t.Errorf("Expected begin_break_time to be '2014-10-10T12:00:00+09:00'. Got '%v'", m["begin_break_time"])
	}

	if m["end_break_time"] != "2014-10-10T13:00:00+09:00" {
		t.Errorf("Expected end_break_time to be '2014-10-10T13:00:00+09:00'. Got '%v'", m["end_break_time"])
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
}

//func TestUpdateWorkRecord(t *testing.T) {
//	clearTable()
//	addWorkRecord(1)
//
//	var work models.WorkRecord
//	a.DB.First(&work)
//	path := `/work_records/` + fmt.Sprint(work.ID)
//
//	req, _ := http.NewRequest("GET", path, nil)
//	response := executeRequest(req)
//	var beforeWorkRecord map[string]interface{}
//	json.Unmarshal(response.Body.Bytes(), &beforeWorkRecord)
//
//	param := `
//{
//"work_date": "2014-10-10T00:00:00+09:00",
//"begin_work_time": "2014-10-10T10:00:00+09:00",
//"end_work_time": "2014-10-10T19:00:00+09:00",
//"begin_break_time": "2014-10-10T12:00:00+09:00",
//"end_break_time": "2014-10-10T13:00:00+09:00",
//}
//`
//	payload := []byte(param)
//	req, _ = http.NewRequest("PUT", path, bytes.NewBuffer(payload))
//	response = executeRequest(req)
//
//	checkResponseCode(t, http.StatusOK, response.Code)
//
//	var m map[string]interface{}
//
//	if m["id"] != beforeWorkRecord["id"] {
//		t.Errorf("Expected the id to remain the same (%v). Got %v", beforeWorkRecord["id"], m["id"])
//	}
//
//	if m["work_date"] == beforeWorkRecord["work_date"] {
//		t.Errorf("Expected the work_date to remain the same (%v). Got %v", beforeWorkRecord["work_date"], m["work_date"])
//	}
//
//	if m["BeginWorkTime"] == beforeWorkRecord["BeginWorkTime"] {
//		t.Errorf("Expected the BeginWorkTime to remain the same (%v). Got %v", beforeWorkRecord["BeginWorkTime"], m["BeginWorkTime"])
//	}
//
//	if m["EndWorkDate"] == beforeWorkRecord["EndWorkDate"] {
//		t.Errorf("Expected the EndWorkDate to remain the same (%v). Got %v", beforeWorkRecord["EndWorkDate"], m["EndWorkDate"])
//	}
//
//	if m["BeginBreakTime"] == beforeWorkRecord["BeginBreakTime"] {
//		t.Errorf("Expected the BeginBreakTime to remain the same (%v). Got %v", beforeWorkRecord["BeginBreakTime"], m["BeginBreakTime"])
//	}
//
//	if m["EndBreakTime"] == beforeWorkRecord["EndBreakTime"] {
//		t.Errorf("Expected the EndBreakTime to remain the same (%v). Got %v", beforeWorkRecord["EndBreakTime"], m["EndBreakTime"])
//	}
//}

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
		work.EndWorkDate = t2
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
