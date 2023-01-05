package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"youtube-livestreaming-notificator/internal/config"
)

func GetPlaylistById(ids []string) (Channels, error) {
	idsString := strings.Join(ids, ",")
	url := "https://www.googleapis.com/youtube/v3/channels?part=contentDetails,snippet&id=" + idsString + "&key=" + config.GetConfig().YOUTUBE.APIKEY
	resp, err := http.Get(url)
	var channels Channels
	if err != nil {
		return channels, err
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(byteArray, &channels); err != nil {
		fmt.Println(err)
		return channels, err
	}
	return channels, nil
}

func GetPlaylistItems(id string) (Videos, error) {
	url := "https://www.googleapis.com/youtube/v3/playlistItems?part=snippet&playlistId=" + id + "&key=" + config.GetConfig().YOUTUBE.APIKEY
	resp, err := http.Get(url)
	var videos Videos
	if err != nil {
		return videos, err
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(byteArray, &videos); err != nil {
		return videos, err
	}
	return videos, nil
}

func GetVideos(ids []string) (VideoDetails, error) {
	idsString := strings.Join(ids, ",")
	url := "https://www.googleapis.com/youtube/v3/videos?part=liveStreamingDetails,snippet&id=" + idsString + "&key=" + config.GetConfig().YOUTUBE.APIKEY
	resp, err := http.Get(url)
	var videoDetails VideoDetails
	if err != nil {
		return videoDetails, err
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(byteArray, &videoDetails); err != nil {
		return videoDetails, err
	}
	return videoDetails, nil
}
