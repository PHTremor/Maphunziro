package main

import (
	"fmt"
	"reflect"
)

type User struct {
	UserID   string `tagA:"valueA1" tagB:"valueA2"`
	Email    string `tagB:"value"`
	Password string `tagC:"v1 v2"`
}

func main() {
	T := reflect.TypeOf(User{})

	// Accessing first Tag
	fieldUserId := T.Field(0)
	t := fieldUserId.Tag
	fmt.Println("Struct Tag is: ", t)
	v, _ := t.Lookup("TagA")
	fmt.Printf("TagA: %s\n", v)

	// Accesing second Tag
	fieldEmail, _ := T.FieldByName("Email")
	vEmail := fieldEmail.Tag.Get("tagB")
	fmt.Println("email tagB: ", vEmail)

	// Accessing Third tag
	fieldPassword,_ := T.FieldByName("Password")
	fmt.Printf("Password tags: [%s]\n", fieldPassword.Tag)
	fmt.Println(fieldPassword.Tag.Get("tagC"))
}
