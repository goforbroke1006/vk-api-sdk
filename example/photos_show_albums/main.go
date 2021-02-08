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
		response, err := client.PhotosGetAlbums(ownerID)
		if err != nil {
			panic(err)
		}
		if response.Error != nil {
			fmt.Printf("WARN: %d %s (%v)\n",
				response.Error.Code, response.Error.Message, response.Error.RequestParams)
			continue
		}

		for _, albumData := range response.Response.Items {
			//if albumData.ID != vksdk.PhotosSystemAlbumSavedPhotos {
			//	continue
			//}

			fmt.Println("ID          : ", albumData.ID)
			fmt.Println("Title       : ", albumData.Title)
			fmt.Println("Description : ", albumData.Description)
		}
	}

}
