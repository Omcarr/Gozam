package audiofingerprint

import (
	// "fmt"
	"gozam/db"
	"gozam/models"
	"log"

	// "log"
	"math"
	"sort"
	"time"
)

const (
	maxFreqBits    = 9
	maxDeltaBits   = 14
	targetZoneSize = 5
)

type Match struct {
	SongID     uint32
	SongTitle  string
	SongArtist string
	YouTubeID  string
	Timestamp  uint32
	Score      float64
}

// Fingerprint generates fingerprints from a list of peaks and stores them in an array.
// Each fingerprint consists of an address and a couple.
// The address is a hash. The couple contains the anchor time and the song ID.
func CreateFingerprint(peaks []Peak, songID uint32) map[uint32]models.Couple {
	fingerprints := map[uint32]models.Couple{}

	for i, anchor := range peaks {
		for j := i + 1; j < len(peaks) && j <= i+targetZoneSize; j++ {
			target := peaks[j]

			address := createAddress(anchor, target)
			anchorTimeMs := uint32(anchor.Time * 1000)

			fingerprints[address] = models.Couple{anchorTimeMs, songID}
		}
	}

	return fingerprints
}

// createAddress generates a unique address for a pair of anchor and target points.
// The address is a 32-bit integer where certain bits represent the frequency of
// the anchor and target points, and other bits represent the time difference (delta time)
// between them. This function combines these components into a single address (a hash).
func createAddress(anchor, target Peak) uint32 {
	anchorFreq := int(real(anchor.Freq))
	targetFreq := int(real(target.Freq))
	deltaMs := uint32((target.Time - anchor.Time) * 1000)

	// Combine the frequency of the anchor, target, and delta time into a 32-bit address
	address := uint32(anchorFreq<<23) | uint32(targetFreq<<14) | deltaMs

	return address
}

// FindMatchesFGP uses the sample fingerprint to find matching songs in the database.
func FindMatchesFGP(sampleFingerprint map[uint32]uint32) ([]Match, time.Duration, error) {
	startTime := time.Now()
	// logger := utils.GetLogger()

	addresses := make([]uint32, 0, len(sampleFingerprint))
	for address := range sampleFingerprint {
		addresses = append(addresses, address)
	}

	redisClient, err := db.NewRedisClient()
	if err != nil {
		return nil, time.Since(startTime), err
	}
	defer redisClient.Close()

	m, err := db.GetCouples(redisClient, addresses)
	if err != nil {
		return nil, time.Since(startTime), err
	}

	matches := map[uint32][][2]uint32{}        // songID -> [(sampleTime, dbTime)]
	timestamps := map[uint32]uint32{}          // songID -> earliest timestamp
	targetZones := map[uint32]map[uint32]int{} // songID -> timestamp -> count

	for address, couples := range m {
		for _, couple := range couples {
			matches[couple.SongID] = append(
				matches[couple.SongID],
				[2]uint32{sampleFingerprint[address], couple.AnchorTimeMs},
			)

			if existingTime, ok := timestamps[couple.SongID]; !ok || couple.AnchorTimeMs < existingTime {
				timestamps[couple.SongID] = couple.AnchorTimeMs
			}

			if _, ok := targetZones[couple.SongID]; !ok {
				targetZones[couple.SongID] = make(map[uint32]int)
			}
			targetZones[couple.SongID][couple.AnchorTimeMs]++
		}
	}

	// matches = filterMatches(10, matches, targetZones)

	scores := analyzeRelativeTiming(matches)

	// log.Print(scores)
	var matchList []Match

	for songID, points := range scores {
		song, songExists, err := db.GetSongByID(songID)
		log.Print(song)

		if !songExists {
			// logger.Info(fmt.Sprintf("song with ID (%v) doesn't exist", songID))
			continue
		}
		if err != nil {
			// logger.Info(fmt.Sprintf("failed to get song by ID (%v): %v", songID, err))
			continue
		}

		match := Match{songID, song.Title, song.Artist, song.YouTubeID, timestamps[songID], points}
		matchList = append(matchList, match)
	}

	sort.Slice(matchList, func(i, j int) bool {
		return matchList[i].Score > matchList[j].Score
	})

	return matchList, time.Since(startTime), nil
}

// analyzeRelativeTiming calculates a score for each song based on the
// relative timing between the song and the sample's anchor times.
func analyzeRelativeTiming(matches map[uint32][][2]uint32) map[uint32]float64 {
	scores := make(map[uint32]float64)
	for songID, times := range matches {
		count := 0
		for i := 0; i < len(times); i++ {
			for j := i + 1; j < len(times); j++ {
				sampleDiff := math.Abs(float64(times[i][0] - times[j][0]))
				dbDiff := math.Abs(float64(times[i][1] - times[j][1]))
				if math.Abs(sampleDiff-dbDiff) < 100 { // Allow some tolerance
					count++
				}
			}
		}
		scores[songID] = float64(count)
	}
	return scores
}
