package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestURLWithOnlyPath(t *testing.T) {
	src := "https://homepages.cae.wisc.edu/~ece533/images/airplane.png"
	dest, err := os.Getwd()
	if err != nil {
		t.Fatal("Error in getting destination")
	}

	fileName, err := DownloadAndSave(src, dest)
	if err != nil {
		t.Fatal("Download failed")
	}

	if _, err = os.Stat(filepath.Join(dest, fileName)); os.IsNotExist(err) {
		t.Fatal("Download failed")
	}
}

func TestURLWithQueryParams(t *testing.T) {
	src := "https://homepages.cae.wisc.edu/~ece533/images/boat.png?q=10&fft=101"
	dest, err := os.Getwd()
	if err != nil {
		t.Fatal("Error in getting destination")
	}

	fileName, err := DownloadAndSave(src, dest)
	if err != nil {
		t.Fatal("Download failed")
	}

	if _, err = os.Stat(filepath.Join(dest, fileName)); os.IsNotExist(err) {
		t.Fatal("Download failed")
	}
}

func TestDestOtherThanWorkingDirectory(t *testing.T) {
	src := "https://homepages.cae.wisc.edu/~ece533/images/boat.png?q=10&fft=101"
	dest, err := os.Getwd()
	if err != nil {
		t.Fatal("Error in getting destination")
	}
	dest = filepath.Join(dest, "temp")
	err = os.Mkdir(dest, 0755)
	if err != nil {
		t.Fatal("Error in creating destination folder")
	}

	fileName, err := DownloadAndSave(src, dest)
	if err != nil {
		t.Fatal("Download failed")
	}

	if _, err = os.Stat(filepath.Join(dest, fileName)); os.IsNotExist(err) {
		t.Fatal("Download failed")
	}
}

func TestDownloadAlreadyExistingFile(t *testing.T) {
	src := "https://homepages.cae.wisc.edu/~ece533/images/fruits.png"
	dest, err := os.Getwd()
	if err != nil {
		t.Fatal("Error in getting destination")
	}
	dest = filepath.Join(dest, "temp")
	if _, err = os.Stat(dest); os.IsNotExist(err) {
		err = os.Mkdir(dest, 0755)
		if err != nil {
			t.Fatal("Error in creating destination folder")
		}
	}

	fileName, err := DownloadAndSave(src, dest)
	if err != nil {
		t.Fatal("Download failed")
	}

	if _, err = os.Stat(filepath.Join(dest, fileName)); os.IsNotExist(err) {
		t.Fatal("Download failed")
	}

	fileName, err = DownloadAndSave(src, dest)
	if err != nil {
		t.Fatal("Download failed")
	}

	if _, err = os.Stat(filepath.Join(dest, fileName)); os.IsNotExist(err) {
		t.Fatal("Download failed")
	}
}
