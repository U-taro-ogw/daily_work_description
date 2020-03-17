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

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App)initializeRoutes() {
	a.Router.HandleFunc("/work_records/{id:[0-9]+}", a.getWorkRecord).Methods("GET")
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

func respondWithError(w http.ResponseWriter, code int, message string)  {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{})  {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}