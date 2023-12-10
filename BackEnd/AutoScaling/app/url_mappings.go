package app

import (
	controller "AutoScaling/controller/autoscaling"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Users Mapping
	router.GET("/services", controller.GetScalableServices)
	router.GET("/stats", controller.GetServicesAndStats)
	router.POST("/stats", controller.GetStatsByService)
	router.POST("/scale", controller.ScaleService)
	router.DELETE("/container/:id", controller.DeleteContainer)

	log.Info("Finishing mappings configurations")
}
