package main

// Import needed packages
import (
	"github.com/finspect/finspect"
	"gopkg.in/fsnotify.v1"
	"log"
)

// Example watcher from fsnotify
func ExampleNewWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/tmp")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func main() {
	ExampleNewWatcher()
}