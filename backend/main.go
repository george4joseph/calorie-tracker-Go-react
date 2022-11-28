package main

import (
	"github.com/george4joseph/calorie-tracker-Go-react/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create",routes.AddEntry)
	router.GET("/entries",routes.GetEntries)
	router.GET("/entry/:id",routes.GetEntryID)
	router.GET("/incredients/:id",routes.GetEntriesByIncredient)
	// router.GET("/incredients/:id",routes.GetIncredientsID)

	router.PUT("/entry/update/:id",routes.UpdateEntry)
	router.PUT("/incredient/update/:id",routes.UpdateIncredient)
	router.DELETE("/entry/delete/:id",routes.DeleteEntry)

	router.Run(":8000")



}