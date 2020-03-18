package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"github.com/U-taro-ogw/daily_work_description/work_api/db"
	"github.com/U-taro-ogw/daily_work_description/work_api/models"

	//"fmt"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize() {
	a.DB = db.MysqlConnect()
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App)initializeRoutes() {
	a.Router.HandleFunc("/work_records", a.getWorkRecords).Methods("GET")
	a.Router.HandleFunc("/work_records/{id:[0-9]+}", a.getWorkRecord).Methods("GET")
	a.Router.HandleFunc("/work_records", a.createWorkRecord).Methods("POST")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) getWorkRecords(w http.ResponseWriter, r *http.Request)  {
	wrs := []models.WorkRecord{}
	respondWithJSON(w, http.StatusOK, wrs)
}

func (a *App) getWorkRecord(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	wr := models.WorkRecord{Model: gorm.Model{ID: uint(id)}}
	if err := wr.GetWorkRecord(a.DB); err != nil {
		respondWithError(w, http.StatusNotFound, "Not Found")
		return
	}

	respondWithJSON(w, http.StatusOK, wr)
}

func (a *App) createWorkRecord(w http.ResponseWriter, r *http.Request) {
	var wr models.WorkRecord
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&wr); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := wr.CreateWorkRecord(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, wr)
}

func respondWithError(w http.ResponseWriter, code int, message string)  {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{})  {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}