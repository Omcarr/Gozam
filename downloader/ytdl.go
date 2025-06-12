package downloader

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
)

type YTVideoInfo struct {
	Title       string `json:"title"`
	ID          string `json:"id"`
	WebpageURL  string `json:"webpage_url"`
	Uploader    string `json:"uploader"`
	ChannelID   string `json:"channel_id"`
	Duration    int    `json:"duration"`
	Description string `json:"description"`
	UploadDate  string `json:"upload_date"`
	ViewCount   int    `json:"view_count"`
	LikeCount   int    `json:"like_count"`
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
		log.Fatal("yt-dlp failed:", err)
	}

	return nil
}

func GetVideoDetails(id string) error {
	url := id
	cmd := exec.Command("yt-dlp", "-j", url)
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to get video info: %v", err)
	}

	var info YTVideoInfo
	if err := json.Unmarshal(out, &info); err != nil {
		log.Fatalf("Failed to parse video info: %v", err)
	}

	log.Printf("Title: %s\nUploader: %s\nDuration: %d seconds\nViews: %d\nURL: %s\n",
		info.Title, info.Uploader, info.Duration, info.ViewCount, info.WebpageURL)

	return nil
}
