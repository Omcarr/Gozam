package utils

import (
	"io"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
var ytRegex = regexp.MustCompile(`^https?://(www\.)?(youtube\.com/watch\?v=|youtu\.be/)[\w-]{11}$`)

func MoveFile(sourcePath string, destinationPath string) error {
	srcFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}

	destFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	err = srcFile.Close()
	if err != nil {
		return err
	}

	err = os.Remove(sourcePath)
	if err != nil {
		return err
	}

	return nil
}

func GenerateUniqueID() uint32 {
	randomNumber := rnd.Uint32()
	return randomNumber
}

func IsValidYouTubeURL(url string) bool {
	return ytRegex.MatchString(url)
}

// normalizeString removes non-letter/digit characters and lowercases the string
func normalizeString(s string) string {
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}

func FindDownloadedFile(songTitle, downloadPath string) (string, error) {
	normalizedTitle := normalizeString(songTitle)

	var found string
	err := filepath.Walk(downloadPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		normalizedName := normalizeString(info.Name())
		if strings.Contains(normalizedName, normalizedTitle) {
			found = path
			return filepath.SkipDir
		}
		return nil
	})
	if found == "" {
		return "", os.ErrNotExist
	}
	return found, err
}
