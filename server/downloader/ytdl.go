package downloader

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"time"
)

type YouTubeResponse struct {
	Items []VideoItem `json:"items"`
}

type VideoItem struct {
	ID      string  `json:"id"`
	Snippet Snippet `json:"snippet"`
}

type Snippet struct {
	Title        string    `json:"title"`
	ChannelTitle string    `json:"channelTitle"`
	Thumbnail    Thumbnail `json:"thumbnails"`
}

type Thumbnail struct {
	Standard Standard `json:"standard"`
}

type Standard struct {
	Url string `json:"url"`
}

type YouTubePlaylistResponse struct {
	Items []PlaylistVideoItem `json:"items"`
}

type PlaylistVideoItem struct {
	Snippet        PlaylistSnippet        `json:"snippet"`
	ContentDetails PlaylistContentDetails `json:"contentDetails"`
}

type PlaylistSnippet struct {
	Title        string    `json:"title"`
	ChannelTitle string    `json:"videoOwnerChannelTitle"`
	Thumbnail    Thumbnail `json:"thumbnails"`
}

type PlaylistContentDetails struct {
	ID string `json:"videoId"`
}

func DownloadYTaudio(id, path string) error {
	url := id

	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		log.Fatal("Failed to create output directory:", err)
	}

	cmd := exec.Command(
		"yt-dlp",
		"-f", "bestaudio/bestaudio[ext=m4a]/140",
		"--no-playlist",
		"--no-check-certificate",
		"--force-ipv4",
		"-o", path+"/%(title)s.%(ext)s",
		url,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Println("yt-dlp failed:", err)
		return fmt.Errorf("yt-dlp error: %w", err)
	}

	return nil
}

func GetVideoDetails(id string) (*YouTubeResponse, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	parsedURL, err := url.Parse(id)
	if err != nil {
		return nil, err
	}

	parsedId := parsedURL.Query().Get("v")

	ytApiKey := os.Getenv("ytApiKey")
	apiURL := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?id=%s&key=%s&part=snippet,statistics", parsedId, ytApiKey)

	response, err := client.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API call failed: %s", response.Status)
	}

	var ytResp YouTubeResponse
	err = json.NewDecoder(response.Body).Decode(&ytResp)
	if err != nil {
		return nil, err
	}

	return &ytResp, nil
}

func GetPlaylistDetails(id string) (*YouTubePlaylistResponse, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	parsedURL, err := url.Parse(id)
	if err != nil {
		return nil, err
	}

	parsedId := parsedURL.Query().Get("list")

	ytApiKey := os.Getenv("ytApiKey")
	apiURL := fmt.Sprintf("https://www.googleapis.com/youtube/v3/playlistItems?playlistId=%s&key=%s&part=snippet,contentDetails&maxResults=50", parsedId, ytApiKey)

	response, err := client.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API call failed: %s", response.Status)
	}

	var ytResp YouTubePlaylistResponse
	err = json.NewDecoder(response.Body).Decode(&ytResp)
	if err != nil {
		return nil, err
	}

	return &ytResp, nil
}
