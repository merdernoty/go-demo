package main

// import (
// 	"encoding/json"
// 	"fmt"
// )


// type User struct {
// 	ID int `json:"id"`
// 	Name string `json:"name"`
// 	Email string `json:"email"`
// 	Password string `json:"-"`

// }
// func main () {
// 	user := User{ 
// 		ID: 1,
// 		Name: "Евгений",
// 		Email: "ef42@gmail.com",
// 		Password: "123456",
// 	}
// 	jsonData, _ := json.Marshal(user)
// 	fmt.Printf("Json: %s\n:", jsonData)


// 	jsonUser :=  `{"id":2,"name":"Bob"}`
// 	var newUser User

// 	json.Unmarshal([]byte(jsonUser), &newUser)
// 	fmt.Print(newUser)
// }