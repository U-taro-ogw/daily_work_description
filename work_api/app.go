package main

import (
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"

    "github.com/U-taro-ogw/daily_work_description/work_api/db/connect"
)

type App struct {
    Router *mux.Router
    DB *gorm.DB
}

func (a *App) Initialize() {
    a.DB = connect.MysqlConnection()
    a.Router = mux.NewRouter()
}

func (a *App) Run(addr string)  {

}

