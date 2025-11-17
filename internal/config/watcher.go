package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func Watch(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("Watcher error:", err)
		return
	}
	defer watcher.Close()

	watcher.Add(path)

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&(fsnotify.Write|fsnotify.Create) != 0 {
				log.Println("Config reloaded")
				Load(path)
			}
		case err := <-watcher.Errors:
			log.Println("Watcher error:", err)
		}
	}
}
