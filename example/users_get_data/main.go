package main

import (
	"fmt"
	"os"

	"github.com/goforbroke1006/vksdk"
)

func main() {
	vkToken := os.Getenv("VK_API_TOKEN")
	if len(vkToken) == 0 {
		panic("specify VK_API_TOKEN env var")
	}

	client := vksdk.New(vkToken)
	response, err := client.UsersGet(
		[]interface{}{199177108, 90829611, "7moky_owl"},
		[]string{"sex", "bdate", "city", "country", "has_photo", "has_mobile", "status", "last_seen", "occupation", "relation"},
		vksdk.Nom,
	)
	if err != nil {
		panic(err)
	}

	for _, userData := range response.Response {
		fmt.Println("Name   : ", userData.FirstName, userData.LastName)
		fmt.Println("Birth  : ", userData.BDate)
		fmt.Println("Status : ", userData.Status)
	}
}
