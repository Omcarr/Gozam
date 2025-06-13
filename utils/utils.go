package utils

import (
	"io"
	"math/rand"
	"os"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

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
