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

	for _, ownerID := range []int64{199177108, 90829611, 12317860, 28045291, 291229116, 26133130, 352033509, 25873211} {
		response, err := client.PhotosGet(ownerID, vksdk.SystemAlbumSavedPhotos)
		if err != nil {
			panic(err)
		}
		if response.Error != nil {
			fmt.Printf("WARN: %d %s (%v)\n",
				response.Error.Code, response.Error.Message, response.Error.RequestParams)
			continue
		}

		for _, photoData := range response.Response.Items {
			fmt.Println("Owner ID : ", ownerID)
			fmt.Println("ID       : ", photoData.ID)
			fmt.Println("Title    : ", photoData.Title)
			fmt.Println("Text     : ", photoData.Text)
			fmt.Println("URL      : ", photoData.Sizes[len(photoData.Sizes)-1].Url)
		}
	}

}
