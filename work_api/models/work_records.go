package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type WorkRecord struct {
	gorm.Model

	WorkDate time.Date `json:"work_date"`
	BeginWorkTime time.Time `json:"begin_work_time"`
	EndWorkDate time.Time `json:"end_work_time"`
	BeginBreakTime time.Time `json:"begin_break_time"`
	EndBreakTime time.Time `json:"end_break_time"`
}

func (w *WorkRecord) GetWorkRecord(db *gorm.DB) error {
	return errors.New("未実装")
}

func (w *WorkRecord) UpdateWorkRecord(db *gorm.DB) error {
	return errors.New("未実装")
}

func (w *WorkRecord) DeleteWorkRecord(db *gorm.DB) error {
	return errors.New("未実装")
}

func (w *WorkRecord) CreateWorkRecord(db *gorm.DB) error {
	return errors.New("未実装")
}

func getWorkRecords(db *gorm.DB, begin_date, end_date time.Time) ([]WorkRecord, error) {
	return nil, errors.New("未実装")
}
