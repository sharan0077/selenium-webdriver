package main

import (
	"fmt"
	"github.com/twist2/common"
	"os"
	"path/filepath"
	"strings"
	"regexp"
	"errors"
)

const (
	WebdriverBrowserFactory = "WebdriverBrowserFactory.java"
	WebdriverHooks = "WebdriverHooks.java"
	WebdriverPropertiesFile = "webdriver.properties"
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

	// Copies Files WebdriverBrowserFactory.java && WebdriverHooks.java into project src 
	pluginsPath,err := common.GetPluginsPath()
	if(err != nil){
		fmt.Printf("Could not find plugins path")
		os.Exit(1)
	}
	javaPluginsPath := filepath.Join(pluginsPath,"selenium-webdriver","0.0.0","skel","java")
	projectRoot,err := common.GetProjectRoot()
	if(err != nil){
		fmt.Printf("failed to get Project Source")
		os.Exit(1)
	}
	projectSrc := filepath.Join(projectRoot,"src")
	err = common.CopyFile(filepath.Join(javaPluginsPath,WebdriverBrowserFactory) , filepath.Join(projectSrc,WebdriverBrowserFactory))
	if(err != nil){
		fmt.Printf("failed to Copy File %s ",WebdriverBrowserFactory)
		os.Exit(1)
	}
	err = common.CopyFile(filepath.Join(javaPluginsPath,WebdriverHooks) , filepath.Join(projectSrc,WebdriverHooks))
	if(err != nil){
		fmt.Printf("failed to Copy File %s ",WebdriverHooks)
		os.Exit(1)
	}
	projectEnv := filepath.Join(projectRoot,"env","default")
	err = common.CopyFile(filepath.Join(javaPluginsPath, WebdriverPropertiesFile) , filepath.Join(projectEnv, WebdriverPropertiesFile))
	if (err != nil) {
		fmt.Printf("failed to Copy File %s ", WebdriverPropertiesFile)
		os.Exit(1)
	}
}




func findLatestVersion() (string, error) {
	downloadLink := "http://docs.seleniumhq.org/download/"
	tmp := os.TempDir()
	err := common.Download(downloadLink, tmp) 
	if err != nil{
		return "" , errors.New(fmt.Sprintf("Failed to download %s\n", downloadLink))
	}
	downloadedFileContents := common.ReadFileContents(filepath.Join(tmp, "download"))
	seleniumDownloadLinkPattern := regexp.MustCompile("http://selenium-release.storage.googleapis.com/.*/selenium-server-standalone-.*.jar")
	latestVersion := strings.Split((seleniumDownloadLinkPattern.FindString(downloadedFileContents)),"-")
	seleniumVersion := strings.Split((latestVersion[len(latestVersion)-1]),".jar")[0]
	return seleniumVersion , nil;
}

func main() {
	action := os.Getenv("selenium-webdriver_action")
	language := os.Getenv("test_language")
	seleniumVersion := os.Getenv("selenium_version")
	if action == "setup" {
		switch language {
			case "java" :
				if seleniumVersion == "" {
					latestVersion, err := findLatestVersion()
					if(err != nil){
						fmt.Println(err.Error())
						os.Exit(1)
					}
					seleniumVersion = latestVersion
				}
				versionParts := strings.Split(seleniumVersion, ".")
				if len(versionParts) != 3 {
					fmt.Println("Incorrect version format. It should be major.minor.patch")
					os.Exit(1)
				}
				setupJavaVersion(versionParts)	
			default :
				fmt.Println("Language not supported")
				os.Exit(1)
		}
	}
}		
