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
	flags()
	// init
	systray := system.NewSystray()
	go systray.Run()

	f := filesystem.NewWatcher(configpath)
	f.Run()
	defer f.Watcher.Close()

	c := converter.New(configpath)
	go c.Convert(f.Stream)
	<-make(chan struct{})
}

func flags() {
	flag.StringVar(nil, "foo", "bar", "placeholder")
	flag.Parse()
}
