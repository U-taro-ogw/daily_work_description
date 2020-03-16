package main

import (
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"

    "os"
)

type App struct {
    Router *mux.Router
    DB *gorm.DB
}

func (a *App) Initialize() {

}

func (a *App) Run(addr string)  {

}