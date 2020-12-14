package main

import "github.com/matthewapeters/dali"

func main() {
	// Create the Window
	W := dali.NewWindow(1000, 800, "", "")

	// Add Components to Window

	// Start Window
	W.Start()
	// Awaite window closure
	<-W.GetUI().Done()
}
