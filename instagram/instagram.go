package instagram

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	database "nimeshjohari02.com/restapi/database"
	filehandler "nimeshjohari02.com/restapi/filehandler"
)

type article struct {
	id          string
	title       string
	description string
	urlToImage  string
	publishedAt time.Time
	fileName    string
}

 func addInstaPost(w http.ResponseWriter, r *http.Request) {
	conn := database.InitiateMongoClient()
	db := conn.Database("rest")
	collection := db.Collection("InstaPost")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.ParseForm()
	post := article{
		id:          uuid.New().String(),
		title:       r.Form["title"][0],
		description: r.Form["description"][0],
		urlToImage:  r.Form["urlToImage"][0],
		publishedAt: time.Now(),
		fileName:    uuid.New().String() + ".jpg",
	}

	// result, err := collection.InsertOne(ctx, post)
	result, err := collection.InsertOne(ctx, bson.M{
		"id":          post.id,
		"title":       post.title,
		"description": post.description,
		"urlToImage":  post.urlToImage,
		"publishedAt": post.publishedAt,
		"fileName":    post.fileName,
	})

	filename := post.fileName
	fileSize := filehandler.UploadFile(post.urlToImage, filename)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Inserted a single document: %v\n", result.InsertedID)
	fmt.Fprintf(w, "File size: %v\n", fileSize)
	log.Printf("Write Data Successfully")
}
func init(){
	http.HandleFunc("/addInstaPost", addInstaPost)
}