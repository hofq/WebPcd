package filesystem

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

type FsRead struct {
	Watcher *fsnotify.Watcher
	Stream  chan string
}

func NewWatcher() *FsRead {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	return &FsRead{
		Watcher: watcher,
	}
}

func (f *FsRead) Run() {

	f.Stream = make(chan string)

	// Add a path.
	// err = watcher.Add("%userprofile%\\Downloads")
	var err error
	err = f.Watcher.Add("C:\\Users\\<user>\\Downloads")
	fmt.Println("Set Download Path")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		// Start listening for events.
		for {
			select {
			case event, ok := <-f.Watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					if event.Name[len(event.Name)-4:] == "webp" {
						f.Stream <- event.Name
						fmt.Println("Detected WebP File: ", event.Name)
					} else {
						fmt.Println("Detected Change in Path")
					}
				}
			case err, ok := <-f.Watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	fmt.Println("Exit Loop")

}
func (f *FsRead) Log() {
	for {
		select {
		case log := <-f.Stream:
			fmt.Println(log)

		}
	}
}
