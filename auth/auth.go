package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	database "nimeshjohari02.com/restapi/database"
	users "nimeshjohari02.com/restapi/user"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func authentication(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
	conn:=database.InitiateMongoClient();
	UserCollection := conn.Database("rest").Collection("users")
		r.ParseForm()
		u1 := users.User{
			Email:    r.Form["email"][0],
			Password: r.Form["password"][0],
		}
		var result users.User
		err := UserCollection.FindOne(context.TODO(), bson.M{"email" : u1.Email,}).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
        fmt.Println("result Passwords " , result.Password)
        fmt.Println("User Passwords " , u1.Password)
		if CheckPasswordHash(u1.Password, result.Password){
			fmt.Println("Password Matched")
		}else{
			fmt.Println("Password Not Matched")
		}
	}
}
func init(){
	http.HandleFunc("/auth", authentication)
}
