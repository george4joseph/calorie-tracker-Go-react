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
	router.GET("/entries",router.GetEntries)
	router.GET("/entry/:id",router.GetEntry)
	router.GET("/incredients/:id",router.GetIncredients)

	router.PUT("/entry/update/:id",router.UpdateEntry)
	router.PUT("/incredient/update/:id",router.UpdateIncredient)
	router.DELETE("/entry/delete/:id",router.DeleteEntry)

	router.Run(":8000")



}