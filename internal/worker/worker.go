package worker

import (
	"fmt"
	"log"
	"time"
	"youtube-livestreaming-notificator/internal/config"
	"youtube-livestreaming-notificator/internal/scheduler"
	"youtube-livestreaming-notificator/internal/slack"
	"youtube-livestreaming-notificator/internal/youtube"
)

var idResp youtube.Channels
var notified []string

func init() {
	log.Println("Initializing...")
	var err error
	idResp, err = youtube.GetPlaylistById(config.GetConfig().TARGET_ACCOUNTS)
	if err != nil {
		log.Fatalln(err)
	}
	notified = []string{}
}

func getLivestreamingUpdates() error {
	var videoIds []string
	for _, v := range idResp.Items {
		videosResp, err := youtube.GetPlaylistItems(v.ContentDetails.RelatedPlaylists.Uploads)
		if err != nil {
			log.Println(err)
			return err
		}
		for _, vv := range videosResp.Items {
			videoIds = append(videoIds, vv.Snippet.ResourceID.VideoID)
		}
	}
	videoDetails, err := youtube.GetVideos(videoIds)
	if err != nil {
		log.Println(err)
		return err
	}
	for _, v := range videoDetails.Items {
		if v.Snippet.LiveBroadcastContent == "upcoming" {
			// 配信枠検知
			isNotified := false
			for _, vv := range notified {
				if vv == v.ID {
					isNotified = true
					// Check change status (scheduled time, title)
					scheduledData := scheduler.GetScheduledData()[vv]
					if scheduledData.Video.Snippet.Title != v.Snippet.Title || scheduledData.Video.LiveStreamingDetails.ScheduledStartTime != v.LiveStreamingDetails.ScheduledStartTime {
						slack.SendSlack(scheduledData.Channel, v, "配信情報変更")
						// Change scheduled data
						scheduler.AddSchedule(scheduledData.Channel, v)
					}
					break
				}
			}
			if !isNotified {
				var channel youtube.ChannelsItem
				for _, vv := range idResp.Items {
					if vv.ID == v.Snippet.ChannelID {
						channel = vv
						break
					}
				}
				slack.SendSlack(channel, v, "配信枠検知")
				notified = append(notified, v.ID)
				log.Println("[" + v.Snippet.ChannelTitle + "] > " + v.Snippet.Title + " at " + v.LiveStreamingDetails.ScheduledStartTime.Local().Format("2006/01/02 15:04") + "~")
				scheduler.AddSchedule(channel, v)
			}
		}
	}
	return nil
}

func StartWorker() {
	log.Println("Started worker !")
	getLivestreamingUpdates()
	t := time.NewTicker(5 * time.Minute)
	defer func() {
		fmt.Println("Stopping ticker...")
		t.Stop()
	}()

	for {
		select {
		case <-t.C:
			getLivestreamingUpdates()
		}
	}
}
