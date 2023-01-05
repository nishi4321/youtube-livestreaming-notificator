package slack

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"youtube-livestreaming-notificator/internal/config"
	"youtube-livestreaming-notificator/internal/youtube"

	"github.com/EdlinOrg/prominentcolor"
	"github.com/slack-go/slack"
)

func SendSlack(channel youtube.ChannelsItem, info youtube.VideoItem, footer string) error {
	attachment := slack.Attachment{
		Color:         generateThemeColor(channel.Snippet.Thumbnails.High.URL),
		AuthorName:    info.Snippet.ChannelTitle,
		AuthorSubname: channel.Snippet.CustomURL,
		AuthorLink:    "https://www.youtube.com/" + channel.Snippet.CustomURL,
		AuthorIcon:    channel.Snippet.Thumbnails.High.URL,
		Title:         info.Snippet.Title,
		TitleLink:     "https://www.youtube.com/watch?v=" + info.ID,
		Text:          info.LiveStreamingDetails.ScheduledStartTime.Local().Format("2006/01/02 15:04") + "~",
		ImageURL:      info.Snippet.Thumbnails.Maxres.URL,
		Footer:        footer,
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	err := slack.PostWebhook(config.GetConfig().SLACK, &msg)
	if err != nil {
		return err
	}
	return nil
}

func generateThemeColor(url string) string {
	resp, _ := http.Get(url)
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	cols, err := prominentcolor.KmeansWithArgs(prominentcolor.ArgumentNoCropping|prominentcolor.ArgumentDebugImage|prominentcolor.ArgumentAverageMean, img)
	if err != nil {
		log.Println(err)
	}
	return "#" + cols[1].AsString()
}
