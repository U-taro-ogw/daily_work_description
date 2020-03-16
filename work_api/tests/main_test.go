package main

import (
	"log"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func clearTable()  {
	a.DB.Exec("DELETE FROM work_records")
}