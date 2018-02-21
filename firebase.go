package main

import (
	"golang.org/x/net/context"

	"firebase.google.com/go"
	"google.golang.org/api/option"
	"fmt"
)

func main() {
	opt := option.WithCredentialsFile("/Users/clay/Desktop/ptc/firebase/admin-key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
		//return nil, fmt.Errorf("error initializing app: %v", err)
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(auth)

	users := auth.Users(context.Background(), "")
	for {
		user, err := users.Next()

		if err != nil {
			fmt.Println(err)
			break
		}

		if user == nil {
			break
		}
	}
	fmt.Println(users)

	//auth, err := app.Auth(context.Background())
	//user := &auth.UserToCreate{}
	//user.Email("clay@pointc.io")
	//user.Password("w0rkr33f")
	//
	//userRecord, err := auth.CreateUser(context.Background(), user)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(userRecord)

	//user, err := auth.GetUser(context.Background(), "")
}
