package main

import (
	app "urd/app"
	cache "urd/cache"
	"urd/db"
)

func main() {
	db.StartDbEngine()
	go app.StartRoute()
	go cache.InitCache()

  select{}
}
