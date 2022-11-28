package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/george4joseph/calorie-tracker-Go-react/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var entryCollection *mongo.Collection = OpenCollection(Client, "calories")

func AddEntry(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entry models.Entry

	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	validatationErr := validate.Struct(entry)
	if validatationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validatationErr.Error()})
		fmt.Println(validatationErr)
		return
	}
	entry.ID = primitive.NewObjectID()
	result, err := entryCollection.InsertOne(ctx, entry)
	if err != nil {
		msg := fmt.Sprintf("Item not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(err)
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, result)

}

func GetEntries(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var entries []bson.M
	cursor, err := entryCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	fmt.Println(entries)
	c.JSON(http.StatusOK, entries)

}

func GetEntryID(c *gin.Context) {

	EntryID := c.Params.ByName("id")
	doc_id, _ := primitive.ObjectIDFromHex(EntryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entry bson.M
	if err := entryCollection.FindOne(ctx, bson.M{"_id": doc_id}).Decode(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	fmt.Println(entry)
	c.JSON(http.StatusOK, entry)

}

func UpdateIncredient(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	EntryID := c.Params.ByName("id")
	doc_id, _ := primitive.ObjectIDFromHex(EntryID)

	type Ingredient struct {
		Ingredients *string `json:"ingredients"`
	}
	var ingredient Ingredient

	if err := c.BindJSON(&ingredient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	result, err := entryCollection.UpdateOne(ctx,
		bson.M{"_id": doc_id},
		bson.D{{Key: "$set", Value: bson.D{{Key: "ingredients", Value: ingredient.Ingredients}}}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, result.ModifiedCount)
}

func UpdateEntry(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	EntryID := c.Params.ByName("id")
	doc_id, _ := primitive.ObjectIDFromHex(EntryID)
	var entry models.Entry

	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	validatationErr := validate.Struct(entry)
	if validatationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validatationErr.Error()})
		fmt.Println(validatationErr)
		return
	}

	result, err := entryCollection.ReplaceOne(ctx,
		bson.M{"_id": doc_id},
		bson.M{
			"dish":        entry.Dish,
			"fat":         entry.Fat,
			"ingredients": entry.Ingredients,
			"calories":    entry.Calories,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	fmt.Println(result.ModifiedCount)
	c.JSON(http.StatusOK, result.ModifiedCount)

}

func DeleteEntry(c *gin.Context) {
	entryId := c.Params.ByName("id")
	doc_id, _ := primitive.ObjectIDFromHex(entryId)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	result, err := entryCollection.DeleteOne(ctx, bson.M{"_id": doc_id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)

		return
	}

	defer cancel()
	c.JSON(http.StatusOK, gin.H{"result": result.DeletedCount})

}

func GetEntriesByIncredient(c *gin.Context) {
	ingredient := c.Params.ByName("id")
	var entries []bson.M

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	cursor, err := entryCollection.Find(ctx, bson.M{"ingredients": ingredient})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)

		return
	}
	defer cancel()
	fmt.Println(entries)
	c.JSON(http.StatusOK, entries)

}
