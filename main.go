package main

import (
	"driving_school/backend"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {
	db, err := gorm.Open("sqlite3", "db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	service.InitService(db)
	defer service.CloseService()
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

	app.Bind(&service.MainService)

	app.Run()
}
