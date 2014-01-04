package main

import "code.google.com/p/go.exp/winfsnotify"
import "fmt"
import "strings"
import "io/ioutil"

const MAIL_DIR = "c:/Programs/Servers/MDaemon/Users/aldaviva.com/ben/Newsfeeds.IMAP"

func main() {
	watcher, _ := winfsnotify.NewWatcher()
	watcher.AddWatch(MAIL_DIR, winfsnotify.FS_CREATE)

	fmt.Printf("Watching %s...\n", MAIL_DIR)

	for {
		event := <-watcher.Event
		filename := event.Name
		onNewFile(filename)
	}
}

func onNewFile(filename string) {
	if strings.HasSuffix(filename, ".msg") {
		contentsWithAd, _ := ioutil.ReadFile(filename)

		contentsWithoutAd := removeAd(string(contentsWithAd))

		ioutil.WriteFile(filename, []byte(contentsWithoutAd), 0777)

		fmt.Printf("Removed ad from %s\n", filename)
	}
}

func removeAd(body string) string {
	return strings.Replace(body, "style=3D\"border:1=\r\npx solid #555555;", "style=3D\"display:=\r\n none;", -1)
}