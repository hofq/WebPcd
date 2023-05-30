package main

import (
	"flag"
	"webpcd/converter"
	"webpcd/filesystem"
)

func main() {
	var configpath string
	// init
	flag.StringVar(&configpath, "config", "../config.yaml", "path to config file")
	flag.Parse()

	w := filesystem.NewWatcher(configpath)
	w.Run()
	defer w.Watcher.Close()

	c := converter.New(configpath)
	go c.Convert(w.Stream)
	<-make(chan struct{})
}
