package main

import (
	app "urd/app"
	cache "urd/cache"
	"urd/db"
	service "urd/services"
)

func main() {
	db.StartDbEngine()
	go app.StartRoute()
	go cache.InitCache()
	go service.GetTokenAmadeus()

	select {}
}
