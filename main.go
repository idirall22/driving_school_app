package main

import (
	"log"
	"path/filepath"

	"github.com/jinzhu/gorm"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"

	"driving_school/backend"

	_ "github.com/mattn/go-sqlite3"
)

var databaseDirectory = "db"
var databaseFileName = "db.sqlite3"

// Test model
type Test struct {
}

// Data model
type Data struct {
	ID    int
	Title string `json:"title"`
}

// Hello method
func (t *Test) Hello(s string) *Data {
	return &Data{Title: s + "idir"}
}

func main() {
	test := &Test{}
	if err := service.CreateDatabaseDirFile(databaseDirectory,
		databaseFileName); err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(databaseDirectory, databaseFileName)
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	service.InitService(db)
	defer service.CloseService()

	service := service.MainService
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "driving_school",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(service)
	app.Bind(test)

	app.Run()
}
