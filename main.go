package main

import(
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main() {
	// Set Slack Bot Token and Channel ID
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-xxxxxx-xxxxxxxx-xxxxxx") // replace with your SLACK_BOT_TOKEN
	os.Setenv("CHANNEL_ID", "XXXXXXX") // replace with your CHANNEL_ID

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"files/image.png"} // replace with your preferred file

	for i := 0; i<len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}

		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}