package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"

)

type WorkRecord struct {
	ID        int `json:"id" gorm:"primary_key"`
	WorkDate time.Time `json:"work_date"`
	BeginWorkTime time.Time `json:"begin_work_time"`
	EndWorkTime time.Time `json:"end_work_time"`
	BeginBreakTime time.Time `json:"begin_break_time"`
	EndBreakTime time.Time `json:"end_break_time"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (w *WorkRecord) GetWorkRecord(db *gorm.DB) error {
	return db.First(&w, w.ID).Error
}

func (w *WorkRecord) UpdateWorkRecord(db *gorm.DB) error {
	var record WorkRecord
	if err := db.First(&record, w.ID).Error; err != nil {
		return errors.New("Not Found")
	}

	return db.Model(&w).Updates(w).Error
}

func (w *WorkRecord) DeleteWorkRecord(db *gorm.DB) error {
	return errors.New("未実装")
}

func (w *WorkRecord) CreateWorkRecord(db *gorm.DB) error {
	db.Create(&w)
	return nil
}

func GetWorkRecords(db *gorm.DB) ([]WorkRecord, error) {
	var workRecords []WorkRecord
	db.Find(&workRecords)
	return workRecords, nil
}
