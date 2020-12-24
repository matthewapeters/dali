package main

import (
	"fmt"

	"github.com/matthewapeters/dali"
)

func main() {
	// Create the Window
	W := dali.NewWindow(1000, 800, "", "")

	// Add Components to Window

	b := dali.NewBodyElement("")

	v := dali.NewVideoElement("vid", "vid", 900, 700)
	v.SetStyle(`display:inline-block;width:900px;height:700px;boder:solid 1px #338811`)
	start := dali.NewButton("Start Video", "startButton", "startButton", "startTracks")
	start.BindFunction(dali.ClickEvent, "startTracks", func() { _ = v.StartTracks() })
	stop := dali.NewButton("Stop Video", "stopButton", "stopButton", "stopTracks")
	stop.BindFunction(dali.ClickEvent, "stopTracks", func() { _ = v.StopTracks() })

	b.Elements.AddElement(v)
	b.Elements.AddElement(dali.NewBreak())
	b.Elements.AddElement(start)
	b.Elements.AddElement(stop)

	//b.Elements.AddElement(script)
	W.Elements.AddElement(b)

	// Start Window
	err := W.StartTLS("https-server.crt", "https-server.key")
	if err != nil {
		fmt.Println(err)
		W.Close()
	}
	// Awaite window closure
	<-W.GetUI().Done()

}
