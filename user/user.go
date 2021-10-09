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
	Name     string
	Email    string
	Password string
}
func getUserById(w http.ResponseWriter, r *http.Request) {
	conn := database.InitiateMongoClient();
	UserCollection := conn.Database("rest").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user User
	id := r.URL.Query().Get("id")
	err := UserCollection.FindOne(ctx, bson.M{"id": id}).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}


func addUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	conn := database.InitiateMongoClient();
	UserCollection := conn.Database("rest").Collection("users")
	u1 := User{
		id:       uuid.New().String(),
		Name:	  r.Form["name"][0],
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
		"Name":     u1.Name,
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
	http.HandleFunc("/getUserById", getUserById)
	http.HandleFunc("/addUser", addUser)
}