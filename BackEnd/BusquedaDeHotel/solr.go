package main

import (
	"fmt"
	solr "github.com/vanng822/go-solr/solr"
)

var solrInterface *solr.NewSolrInterface

func InitializeSolr() {
	// Establish a connection to Solr
	si, err := solr.NewSolrInterface("http://localhost:8983/solr", "hotels")
	if err != nil {
		fmt.Println("Error initializing solr")
	}

}
