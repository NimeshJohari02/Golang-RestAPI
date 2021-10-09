package user

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	database "nimeshjohari02.com/restapi/database"
)


func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

type User struct {
	id       string
	Email    string
	Password string
}

func addUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	conn := database.InitiateMongoClient();
	UserCollection := conn.Database("rest").Collection("users")
	u1 := User{
		id:       uuid.New().String(),
		Email:    r.Form["email"][0],
		Password: r.Form["password"][0],
	}
	hash, err := HashPassword(u1.Password)
	if err != nil {
		fmt.Println(err)
	}
	 ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	 defer cancel()

	res, err := UserCollection.InsertOne(ctx, bson.M{
		"id":       u1.id,
		"email":    u1.Email,
		"password": hash,
	})
	if err != nil {
		panic(err)
	}
	id := res.InsertedID
	fmt.Println(id)
}
func init(){
	http.HandleFunc("/addUser", addUser)
}