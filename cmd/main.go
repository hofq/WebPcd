package main

import (
	"webpcd/converter"
	"webpcd/filesystem"
)

func main() {

	w := filesystem.NewWatcher()
	w.Run()
	defer w.Watcher.Close()

	c := converter.New("C:\\Users\\<user>\\Downloads\\", false)
	go c.Convert(w.Stream)
	<-make(chan struct{})
}
