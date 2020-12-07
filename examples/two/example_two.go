package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/matthewapeters/dali"
)

/**
* Draw a randomly located, randomly sized, random colored dot
 */
func doDot(wg *sync.WaitGroup, w *dali.Window, image *image.RGBA, mtx *sync.Mutex) {
	centerX := rand.Intn(350)
	centerY := rand.Intn(350)
	radius := rand.Intn(90)
	radiusf := float64(radius)

	clr := color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}

	/**
	 * the logic for painting a circle is to iterate over all of the pixels within the
	 * circle's bounding box, and check if the sqaure of the corrdinates (after offset
	 * by the center) are less than or equal to the square of the radius - this is the
	 * Pythagorean theorem.
	 */
	for y := centerY - radius; y <= centerY+radius && y <= 300; y++ {
		for x := centerX - radius; x <= centerX+radius && x <= 300; x++ {
			if float64((centerX-x)*(centerX-x)+(centerY-y)*(centerY-y)) <= radiusf*radiusf {
				mtx.Lock()
				image.Set(x, y, clr)
				mtx.Unlock()
			}
		}
	}
	wg.Done()
}

/**
 * This method will run as a Go routine. Every 250 milliseconds, a new image will be produced with
 * 20 randomly sized, colored, and located dots.  The image is pushed to the GUI.  The Go routine
 * watches for the closing of the provided channel, and ends when it sees it closed.
 */
func makeRandomImages(ch chan bool, w *dali.Window) {
	for {
		time.Sleep(250 * time.Millisecond)
		select {
		case <-ch:
			fmt.Println("Channel closed - stop making images")
			return
		default:
			image := image.NewRGBA(image.Rect(0, 0, 300, 300))

			/*Draw some dots on the picture*/
			var wg sync.WaitGroup
			mtx := sync.Mutex{}
			for i := 0; i < 20; i++ {
				wg.Add(1)
				/*
				 * Draw the dot asynchronously -- this will result in overlapping
				 * dots that "interrupt" and bleed into each other, not simply overlap
				 * the mutex is to illustrate that, at some point, race conditions
				 * should be addressed
				 */
				go doDot(&wg, w, image, &mtx)
			}
			// Wait until all of the dots are finished drawing
			wg.Wait()

			// Encode the image as a PNG
			buffer := new(bytes.Buffer)
			if err := png.Encode(buffer, image); err != nil {
				fmt.Println("writeImageWithTemplate unable to encode image", err)
				log.Fatalln("unable to encode image.")
			}
			// Encode the image data as Base64 (non-binary) encoding
			image64 := base64.StdEncoding.EncodeToString(buffer.Bytes())

			// this expression tells tells the web browser that the image content is here,
			// and does not need to be downloaded from a web resource
			imageDump := fmt.Sprintf("data:image/png;base64,%s", image64)

			// this JavaScript will the source content of the img tag
			scriptlet := `document.getElementById("randPic").src="%s"`

			// Send the javascript containing the image and the instruction to modify the image
			w.GetUI().Eval(fmt.Sprintf(scriptlet, imageDump))
		}
	}
}

func main() {
	Window := dali.NewWindow(400, 400, "", "")
	Head := dali.NewHeadElement()
	Title := &dali.TitleElement{Text: "Example Two: Interactive Graphics"}
	Head.Elements.AddElement(Title)
	Window.Elements.AddElement(Head)

	// Create an image element
	img := dali.NewImage("randPic", 300, 300, "")
	img.SetStyle("border:solid 1px #000000;")

	// Create a body element
	Body := dali.NewBodyElement("")
	// Put the image in the body ...
	Body.Elements.AddElement(img)
	// ... and put the body in the Window
	Window.Elements.AddElement(Body)

	/*Launch GUI*/
	Window.Start()

	/*Launch a go routine that will generate random images in the GUI until the GUI is closed*/
	ImageMakeChannel := make(chan bool)
	go makeRandomImages(ImageMakeChannel, Window)

	<-Window.GetUI().Done()
	close(ImageMakeChannel)
}
