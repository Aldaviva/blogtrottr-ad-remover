package main

import "code.google.com/p/go.exp/winfsnotify"
import "fmt"
import "strings"
import "io/ioutil"
import "path/filepath"

const MAIL_DIR = "c:/Programs/Servers/MDaemon/Users/aldaviva.com/ben/Newsfeeds.IMAP"

func main() {
	removeExistingAds()
	watchForNewAds()
}

func removeExistingAds(){
	msgPaths, err := filepath.Glob(MAIL_DIR + "/*.msg")
	if err == nil && msgPaths != nil {
		for _, msgPath := range msgPaths {
			//fmt.Printf("File exists %s\n", msgPath)
			removeAdFromFile(msgPath)
		}
	}
}

func watchForNewAds(){
	watcher, _ := winfsnotify.NewWatcher()
	watcher.AddWatch(MAIL_DIR, winfsnotify.FS_CREATE)

	fmt.Printf("Watching %s...\n", MAIL_DIR)

	for {
		event := <-watcher.Event
		filename := event.Name
		removeAdFromFile(filename)
	}
}

func removeAdFromFile(filename string) {
	if strings.HasSuffix(filename, ".msg") {
		contentsWithAd, _ := ioutil.ReadFile(filename)

		contentsWithoutAd := removeAdFromBody(string(contentsWithAd))

		ioutil.WriteFile(filename, []byte(contentsWithoutAd), 0777)

		fmt.Printf("Removed ad from %s\n", filename)
	}
}

func removeAdFromBody(body string) string {
	// return strings.Replace(body, "style=3D\"border:1=\r\npx solid #555555;", "style=3D\"display:=\r\n none;", -1)
	topAdRemoved := strings.Replace(body, "border=3D\"0\" cellpadding=3D\"0\" cellspacing=3D\"0\"", "style=3D\"display: none;\"", -1)
	topAndBottomAdRemoved := strings.Replace(topAdRemoved, "cellpadding=3D\"0\" width=3D\"100%\"", "style=3D\"display: none;\"", -1)
	return topAndBottomAdRemoved
}