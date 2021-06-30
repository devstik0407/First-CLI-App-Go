package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func DownloadAndSave(src, dest string) (string, error) {
	res, err := http.Get(src)
	if err != nil {
		return "", err
	}

	u, err := url.Parse(src)
	if err != nil {
		return "", err
	}

	pathSplit := strings.Split(u.Path, "/")
	fileName := pathSplit[len(pathSplit)-1]

	fileNameSplit := strings.Split(filepath.Join(dest, fileName), ".")
	if len(fileNameSplit) < 2 {
		return "", errors.New("invalid url for downloading file")
	}
	ext := fileNameSplit[1]

	_, err = os.Stat(filepath.Join(dest, fileName))
	for !os.IsNotExist(err) {
		fmt.Printf("A file with same name %s already exists in the directory", fileName)
		fmt.Println("Enter a name (without extension) for the file:")
		newName := ""
		fmt.Scanln(&newName)
		if newName == "" {
			newName = "untitled"
		}
		fileName = newName + "." + ext
		_, err = os.Stat(fileName)
	}

	f, err := os.Create(filepath.Join(dest, fileName))
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, res.Body)
	if err != nil {
		return "", err
	}
	return fileName, nil
}
