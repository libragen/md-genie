package util

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func scanFiles(dir string) ([]os.FileInfo, error) {
	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}
	exDir := filepath.Dir(ex)
	archivesDir := path.Join(exDir, "archives")
	return ioutil.ReadDir(archivesDir)
}

func LsArchivesMdFiles(dir string) (map[string][]string, error) {
	files, err := scanFiles(dir)
	if err != nil {
		return nil, err
	}
	newsItems := []string{}
	movieItems := []string{}

	for _, f := range files {
		fileName := f.Name()
		if strings.Contains(fileName, "hacknews_") {
			newsItems = append([]string{fileName}, newsItems...)
		}
		if strings.Contains(fileName, "movie_") {
			movieItems = append([]string{fileName}, movieItems...)
		}
	}
	return map[string][]string{"Hack News List": newsItems, "Chinese Movie Board": movieItems}, nil
}
