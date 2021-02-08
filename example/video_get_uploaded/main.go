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
		response, err := client.VideoGet(ownerID, vksdk.VideoSystemAlbumUploaded)
		if err != nil {
			panic(err)
		}
		if response.Error != nil {
			fmt.Printf("WARN: %d %s (%v)\n",
				response.Error.Code, response.Error.Message, response.Error.RequestParams)
			continue
		}

		for _, videoData := range response.Response.Items {
			fmt.Println("Owner ID  : ", ownerID)
			fmt.Println("ID        : ", videoData.ID)
			fmt.Println("Title     : ", videoData.Title)
			fmt.Println("Dimension : ", videoData.Width, "x", videoData.Height)
			fmt.Println("Player    : ", videoData.Player)
			fmt.Println("File 720  : ", videoData.Files.MP4Quality720)
			fmt.Println("File 480  : ", videoData.Files.MP4Quality480)
			fmt.Println("File 360  : ", videoData.Files.MP4Quality360)
			fmt.Println("File 240  : ", videoData.Files.MP4Quality240)
		}
	}

}
