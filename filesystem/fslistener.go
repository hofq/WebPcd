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

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-f.Watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					if event.Name[len(event.Name)-4:] == "webp" {
						f.Stream <- event.Name
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

	// Add a path.
	// err = watcher.Add("%userprofile%\\Downloads")
	var err error
	err = f.Watcher.Add("C:\\Users\\<user>\\Downloads")
	if err != nil {
		log.Fatal(err)
	}

}
func (f *FsRead) Log() {
	for {
		select {
		case log := <-f.Stream:
			fmt.Println(log)

		}
	}
}
