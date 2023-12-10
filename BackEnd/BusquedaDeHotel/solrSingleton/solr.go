package solr

import (
	"github.com/rtt/Go-Solr"
	log "github.com/sirupsen/logrus"
)

var ClienteSolr *solr.Connection

func InitSolr() {

	var err error

	ClienteSolr, err = solr.Init("solr", 8983, "Hotels")
	if err != nil {
		log.Info("Failed to connect to Solr")
		log.Fatal(err)
	} else {
		log.Info("Connected to Solr successfully")
	}
}