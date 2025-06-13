package db

type Song struct {
	Title     string
	Artist    string
	YouTubeID string
}

func GetSongByID(songID uint32) (Song, bool, error) {
	// if songID != 1 {
	// 	// return false if not the test ID
	// 	return Song{}, false, nil
	// }

	// Hardcoded Viva La Vida entry
	song := Song{
		Title:     "Viva La Vida",
		Artist:    "Coldplay",
		YouTubeID: "dvgZkm1xWPE",
	}

	return song, true, nil
}
