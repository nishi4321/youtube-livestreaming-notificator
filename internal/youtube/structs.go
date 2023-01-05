package youtube

import "time"

type Channels struct {
	Kind     string `json:"kind"`
	Etag     string `json:"etag"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []ChannelsItem `json:"items"`
}

type ChannelsItem struct {
	Kind    string `json:"kind"`
	Etag    string `json:"etag"`
	ID      string `json:"id"`
	Snippet struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		CustomURL   string    `json:"customUrl"`
		PublishedAt time.Time `json:"publishedAt"`
		Thumbnails  struct {
			Default struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"default"`
			Medium struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"medium"`
			High struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"high"`
		} `json:"thumbnails"`
		DefaultLanguage string `json:"defaultLanguage,omitempty"`
		Localized       struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"localized"`
		Country string `json:"country"`
	} `json:"snippet"`
	ContentDetails struct {
		RelatedPlaylists struct {
			Likes   string `json:"likes"`
			Uploads string `json:"uploads"`
		} `json:"relatedPlaylists"`
	} `json:"contentDetails"`
}

type Videos struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	Items         []struct {
		Kind    string `json:"kind"`
		Etag    string `json:"etag"`
		ID      string `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
				Standard struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"standard"`
				Maxres struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"maxres"`
			} `json:"thumbnails"`
			ChannelTitle string `json:"channelTitle"`
			PlaylistID   string `json:"playlistId"`
			Position     int    `json:"position"`
			ResourceID   struct {
				Kind    string `json:"kind"`
				VideoID string `json:"videoId"`
			} `json:"resourceId"`
			VideoOwnerChannelTitle string `json:"videoOwnerChannelTitle"`
			VideoOwnerChannelID    string `json:"videoOwnerChannelId"`
		} `json:"snippet"`
	} `json:"items"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
}

type VideoDetails struct {
	Kind     string      `json:"kind"`
	Etag     string      `json:"etag"`
	Items    []VideoItem `json:"items"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
}

type VideoItem struct {
	Kind    string `json:"kind"`
	Etag    string `json:"etag"`
	ID      string `json:"id"`
	Snippet struct {
		PublishedAt time.Time `json:"publishedAt"`
		ChannelID   string    `json:"channelId"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Thumbnails  struct {
			Default struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"default"`
			Medium struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"medium"`
			High struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"high"`
			Standard struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"standard"`
			Maxres struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"maxres"`
		} `json:"thumbnails"`
		ChannelTitle         string   `json:"channelTitle"`
		Tags                 []string `json:"tags"`
		CategoryID           string   `json:"categoryId"`
		LiveBroadcastContent string   `json:"liveBroadcastContent"`
		Localized            struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"localized"`
		DefaultAudioLanguage string `json:"defaultAudioLanguage"`
	} `json:"snippet"`
	LiveStreamingDetails struct {
		ScheduledStartTime time.Time `json:"scheduledStartTime"`
		ActiveLiveChatID   string    `json:"activeLiveChatId"`
	} `json:"liveStreamingDetails"`
}
