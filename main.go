package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"

	"driving_school/backend"

	_ "github.com/mattn/go-sqlite3"
)

var folderName = "auto-ecole/"
var databaseDirectory = folderName + "db"
var databaseFileName = "db.sqlite3"

func main() {
	path := filepath.Join(
		service.GetRootDir(),
		databaseDirectory,
	)

	if err := service.CreateDatabaseDirFile(databaseDirectory,
		databaseFileName); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := gorm.Open("sqlite3", filepath.Join(path, databaseFileName))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	service.InitService(db)
	defer service.CloseService()

	service := service.MainService
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1280,
		Height: 720,
		Title:  "Auto école Noureddine",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(service)
	app.Run()
}
