package main

import (
	"webpcd/filesystem"
)

func main() {
	var w *filesystem.FsRead
	w = filesystem.NewWatcher()
	w.Run()
	w.Log()
	defer w.Watcher.Close()
}
