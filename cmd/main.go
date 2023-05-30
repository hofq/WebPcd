package main

import (
	"flag"
	"fmt"
	"webpcd/converter"
	"webpcd/filesystem"
	"webpcd/system"
)

func main() {
	fmt.Println("Started Application")
	var configpath string
	// init
	flag.StringVar(&configpath, "config", "./config.yaml", "path to config file")
	flag.Parse()
	systray := system.NewSystray()
	go systray.Run()

	f := filesystem.NewWatcher(configpath)
	f.Run()
	defer f.Watcher.Close()

	c := converter.New(configpath)
	go c.Convert(f.Stream)
	<-make(chan struct{})
}
