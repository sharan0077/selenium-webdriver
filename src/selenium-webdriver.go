package main

import (
	"fmt"
	"github.com/twist2/common"
	"os"
	"path/filepath"
	"strings"
)

func setupJavaVersion(versionParts []string) {
	downloadLink := fmt.Sprintf("http://selenium-release.storage.googleapis.com/%s.%s/selenium-server-standalone-%s.%s.%s.jar",
		versionParts[0], versionParts[1], versionParts[0], versionParts[1], versionParts[2])
	if !common.DirExists("libs") {
		err := os.Mkdir("libs", common.NewDirectoryPermissions)
		if err != nil {
			fmt.Printf("Failed to create libs directory. %s\n", err.Error())
			os.Exit(1)
		}
	}

	absLibs, _ := filepath.Abs("libs")
	err := common.Download(downloadLink, absLibs)
	if err != nil {
		fmt.Printf("Failed to download %s. %s\n", downloadLink, err.Error())
		os.Exit(1)
	}
}

func main() {
	action := os.Getenv("selenium-webdriver_action")
	if action == "setup" {
		seleniumVersion := os.Getenv("selenium_version")
		if seleniumVersion == "" {
			fmt.Println("Finding latest version of selenium is not implemented. Provide selenium version")
			os.Exit(1)
		}

		versionParts := strings.Split(seleniumVersion, ".")
		if len(versionParts) != 3 {
			fmt.Println("Incorrect version format. It should be major.minor.patch")
			os.Exit(1)
		}

		language := os.Getenv("test_language")
		switch language {
		case "java":
			setupJavaVersion(versionParts)
		default:
			fmt.Println("Language not supported")
			os.Exit(1)
		}
	}
}
