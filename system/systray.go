package system

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
)

type Systray struct {
	Title   string
	Tooltip string
}

func NewSystray() *Systray {
	return &Systray{
		Title:   "WebPcd",
		Tooltip: "WebP Container Daemon",
	}
}

func (s *Systray) Run() {
	systray.Run(s.onReady, s.onExit)
}

func (s *Systray) onExit() {
	fmt.Println("Ending process . . .")
	os.Exit(4)
}
func (s *Systray) onReady() {
	systray.SetTitle(s.Title)
	systray.SetTooltip(s.Tooltip)
	mQuitOrig := systray.AddMenuItem("Quit", "Exit")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()
}
