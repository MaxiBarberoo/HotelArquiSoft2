package main

import (
	app "fichadehotel/app"
	"fichadehotel/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
