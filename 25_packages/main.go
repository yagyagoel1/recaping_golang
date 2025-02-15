package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/yagyagoel1/recaping_golang/auth"

	"github.com/yagyagoel1/recaping_golang/user"
)

func main() {
	auth.LoginWithCredentials("ram", "shyam")
	session := auth.GetSession()
	fmt.Println("session", session)

	myUser := user.User{
		Email: "user@mgami.com",
		Name:  "johndoe",
	}
	fmt.Println("username", myUser.Name)

	fmt.Println("myuser", myUser)
	color.Cyan("Prints text in cyan.")
}
