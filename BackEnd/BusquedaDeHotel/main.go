package main

import (
	app "busquedadehotel/app"
	service "busquedadehotel/services"
	solr "busquedadehotel/solrSingleton"
)

func main() {
	solr.InitSolr()

	go service.Consumer()

	app.StartRoute()

	select {}
}
