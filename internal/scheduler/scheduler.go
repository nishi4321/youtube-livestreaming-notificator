package scheduler

import (
	"fmt"
	"log"
	"time"
	"youtube-livestreaming-notificator/internal/slack"
	"youtube-livestreaming-notificator/internal/youtube"
)

type ScheduleInfo struct {
	Channel youtube.ChannelsItem
	Video   youtube.VideoItem
}

var schedules map[string]ScheduleInfo

func init() {
	schedules = map[string]ScheduleInfo{}
}

func StartScheduler() {
	log.Println("Started scheduler !")
	//
	t := time.NewTicker(1 * time.Second)
	defer func() {
		fmt.Println("Stopping ticker...")
		t.Stop()
	}()

	for {
		select {
		case <-t.C:
			for id, v := range schedules {
				if isEqualTime(v.Video.LiveStreamingDetails.ScheduledStartTime.Local(), time.Now()) {
					go func(id string, v ScheduleInfo) {
						// Get latest video info
						video, err := youtube.GetVideos([]string{id})
						if err != nil {
							// if failed, use old info.
							video.Items[0] = v.Video
						}
						slack.SendSlack(v.Channel, video.Items[0], "配信開始")
						delete(schedules, id)
						log.Println("Schedule notified and delete. " + video.Items[0].Snippet.ChannelTitle)
					}(id, v)
				}
			}
		}
	}
}

func AddSchedule(channel youtube.ChannelsItem, video youtube.VideoItem) {
	t := video.LiveStreamingDetails.ScheduledStartTime
	// 過去の日時なら登録しない
	if t.Local().Unix() < time.Now().Unix() {
		return
	}
	schedule := ScheduleInfo{
		Channel: channel,
		Video:   video,
	}
	schedules[video.ID] = schedule
	log.Println("Add scheduler at " + video.LiveStreamingDetails.ScheduledStartTime.Local().Format("2006/01/02 15:04"))
}

func isEqualTime(t1 time.Time, t2 time.Time) bool {
	if t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day() && t1.Hour() == t2.Hour() && t1.Minute() == t2.Minute() && t1.Second() == t2.Second() {
		return true
	} else {
		return false
	}

}

func GetScheduledData() map[string]ScheduleInfo {
	return schedules
}
