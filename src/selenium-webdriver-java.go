package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	downloadLink = "http://selenium-release.storage.googleapis.com/2.41/selenium-server-standalone-2.41.0.jar"
)

func downloadFile(url, dest string) error {
	out, err := os.Create(filepath.Join(dest, "selenium-server-standalone-2.41.0.jar"))
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(downloadLink)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	action := os.Getenv("selenium-webdriver-java_action")
	if action == "setup" {
		downloadFile(downloadLink, "libs")
		fmt.Println("the selenium webdriver has been setup")
	}
}
