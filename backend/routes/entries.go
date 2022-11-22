package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


var entryCollection *mongo.Collection = openCollection(Client, "calories")


func AddEntry(c *gin.Context) {

}

func GetEntry(c *gin.Context) {

}

func GetEntries(c *gin.Context) {

}

func GetIncredients(c *gin.Context) {

}

func UpdateEntry(c *gin.Context) {

}

func DeleteEntry(c *gin.Context) {
	entryId := c.Params.ByName("id")
	doc_id , _ := primitive.ObjectIDFromHex(entryId)

	var ctx, cancel =  context.WithTimeout(context.Background(), 100*time.Second)

	result , err := entryCollection.DeleteOne(ctx, bson.M{"id":doc_id})

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		fmt.Println(err)

		return
	}

	defer cancel() 
	c.JSON(http.StatusOK,gin.H{"result":result.DeletedCount})

}

func UpdateIncredient(c *gin.Context) {

}
