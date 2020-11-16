package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/cmplx"
	"os"
	"runtime"
	"strconv"
	"sync"

	"github.com/matthewapeters/dali"
)

//Point describes a position in the image
type Point struct {
	X, Y int
}

// ImageBoundingBox Describes the dimensions of the "physical" viewport desribed by the image (pixel range)
type ImageBoundingBox struct{ UpperLeft, LowerRight Point }

// CBoundingBox Describes the dimensions of the "imaginary plane" viewport described by the image
type CBoundingBox struct{ UpperLeft, LowerRight complex128 }

//ViewPort describes the portion of C that is visible at the current zoom level
type ViewPort struct {
	Image                    ImageBoundingBox
	C                        CBoundingBox
	ImaginaryPlaneFocalPoint complex128
	ZoomLevel                float64
	Pallette                 func(*Iterant) color.Color
}

// Iterant represents the state of the complex number "Zee" at
// iteration "N" and the result when it when the function is applied
type Iterant struct {
	Pixel       Point
	N           int
	Threshold   int
	ZeeNPlusOne complex128
	ViewPort
}

func drawChunk(start, chunk, height, iterations int, view *ViewPort, wg *sync.WaitGroup, chout chan *Iterant) {
	max := start + chunk
	imageWidth := view.Image.LowerRight.X - view.Image.UpperLeft.X
	imageHeight := view.Image.LowerRight.Y - view.Image.UpperLeft.Y
	cPlaneReal := real(view.C.LowerRight) - real(view.C.UpperLeft)
	cPlaneImag := imag(view.C.UpperLeft) - imag(view.C.LowerRight)

	for a := start; a < max; a++ {
		for b := 0; b < height; b++ {
			offsetX := float64(a) - (float64(imageWidth) / 2.0)
			offsetY := float64(b) - (float64(imageHeight) / 2.0)
			z := complex(
				(offsetX/float64(imageWidth)*cPlaneReal*view.ZoomLevel + real(view.ImaginaryPlaneFocalPoint)),
				(offsetY/float64(imageHeight)*cPlaneImag*view.ZoomLevel + imag(view.ImaginaryPlaneFocalPoint)))
			zn := z
			var n int
			for n = 0; n < iterations; n++ {
				zn = zn*zn + z
				if cmplx.Abs(zn) > 2 {
					break
				}
			}
			iterant := &Iterant{
				Pixel:       Point{X: a, Y: b},
				Threshold:   iterations,
				ViewPort:    *view,
				N:           n,
				ZeeNPlusOne: zn,
			}
			chout <- iterant
		}
	}
	wg.Done()
}

// DrawMandelbrot will Draw Mandelbrot images
func DrawMandelbrot(view *ViewPort, iterations int, display *dali.Image) {

	chunkWG := sync.WaitGroup{}
	width := view.Image.LowerRight.X - view.Image.UpperLeft.X
	height := view.Image.LowerRight.Y - view.Image.UpperLeft.Y
	chout := make(chan *Iterant, width*height)
	defer close(chout)

	image := image.NewRGBA(image.Rect(0, 0, width, height))
	NoCPUs := runtime.NumCPU()
	Chunk := width / NoCPUs * int(math.Max((float64(iterations)/1000.0), 1.0))

	for a := 0; a < width; a += Chunk {
		chunkWG.Add(1)
		go drawChunk(a, Chunk, height, iterations, view, &chunkWG, chout)
	}

	// Wait until all of the pixels are done are finished drawing
	for pixelCount := 0; pixelCount < width*height; pixelCount++ {
		i := <-chout
		if i != nil {
			image.Set(i.Pixel.X, i.Pixel.Y, view.Pallette(i))
		}
	}
	chunkWG.Wait()
	display.Load(image)
}

//FindMandelbrotSet iterates over C to determine members of the Mandelbrot set
func FindMandelbrotSet(w *dali.Window, image *dali.Image, control *sync.Mutex, vp *ViewPort) {
	for i := 0; i < 5000; i += 25 {
		control.Lock()
		w.GetUI().Eval(fmt.Sprintf(`document.getElementById("iterations").value="%d";`, i))
		DrawMandelbrot(
			vp,
			i,
			image)
		control.Unlock()
	}
}

// ZoomIn will reduce the range of the C Plan within the viewport
func ZoomIn(image *dali.Image, iterations, zoomLevel *dali.InputElement, vp *ViewPort, control *sync.Mutex) {
	control.Lock()
	defer control.Unlock()

	vp.ZoomLevel *= 0.9

	zoomLevel.Set(fmt.Sprintf("%.14f", vp.ZoomLevel))
	iv := iterations.Value()
	i, _ := strconv.Atoi((iv))
	DrawMandelbrot(
		vp,
		i,
		image)
}

// ZoomOut will increase the range of the C Plan within the viewport
func ZoomOut(image *dali.Image, iterations, zoomLevel *dali.InputElement, vp *ViewPort, control *sync.Mutex) {
	control.Lock()
	defer control.Unlock()

	vp.ZoomLevel *= (10.0 / 9.0)
	zoomLevel.Set(fmt.Sprintf("%.14f", vp.ZoomLevel))
	iv := iterations.Value()
	i, _ := strconv.Atoi((iv))
	DrawMandelbrot(
		vp,
		i,
		image)
}

// PanLeft will shift the Complex Plane to the right within the viewport
func PanLeft(image *dali.Image, iterations, focalPointReal *dali.InputElement, vp *ViewPort, control *sync.Mutex) {
	control.Lock()
	defer control.Unlock()
	length := 4 * vp.ZoomLevel
	vp.ImaginaryPlaneFocalPoint -= complex(0.1*length, 0)

	iv := iterations.Value()
	i, _ := strconv.Atoi((iv))
	DrawMandelbrot(
		vp,
		i,
		image)

	focalPointReal.Set(fmt.Sprintf("%.14f", real(vp.ImaginaryPlaneFocalPoint)))
}

// PanRight will shift the Complex Plane to the left within the viewport
func PanRight(image *dali.Image, iterations, focalPointReal *dali.InputElement, vp *ViewPort, control *sync.Mutex) {
	control.Lock()
	defer control.Unlock()
	length := 4 * vp.ZoomLevel
	vp.ImaginaryPlaneFocalPoint += complex(0.1*length, 0)

	iv := iterations.Value()
	i, _ := strconv.Atoi((iv))
	DrawMandelbrot(
		vp,
		i,
		image)

	focalPointReal.Set(fmt.Sprintf("%.14f", real(vp.ImaginaryPlaneFocalPoint)))
}

// PanUp will shift the Complex Plane down within the viewport
func PanUp(image *dali.Image, iterations, focalPointImaginary *dali.InputElement, vp *ViewPort, control *sync.Mutex) {
	control.Lock()
	length := 4.0 * vp.ZoomLevel
	vp.ImaginaryPlaneFocalPoint += complex(0, 0.1*length)
	iv := iterations.Value()
	i, _ := strconv.Atoi((iv))
	DrawMandelbrot(
		vp,
		i,
		image)

	focalPointImaginary.Set(fmt.Sprintf("%.14f", imag(vp.ImaginaryPlaneFocalPoint)))
	control.Unlock()
}

// PanDown will shift the Complex Plane up within the viewport
func PanDown(image *dali.Image, iterations, focalPointImaginary *dali.InputElement, vp *ViewPort, control *sync.Mutex) {
	control.Lock()
	length := 4.0 * vp.ZoomLevel
	vp.ImaginaryPlaneFocalPoint -= complex(0, 0.1*length)
	iv := iterations.Value()
	i, _ := strconv.Atoi((iv))
	DrawMandelbrot(
		vp,
		i,
		image)

	focalPointImaginary.Set(fmt.Sprintf("%.14f", imag(vp.ImaginaryPlaneFocalPoint)))
	control.Unlock()
}

//UpdateDisplay will redraw the Mandelbrot set based on Window values
func UpdateDisplay(VP *ViewPort, display *dali.Image, control *sync.Mutex, iterations, zoomLevel, focalPointReal, focalPointImaginary *dali.InputElement) {
	control.Lock()
	defer control.Unlock()
	i, _ := strconv.Atoi(iterations.Value())
	zoom, _ := strconv.ParseFloat(zoomLevel.Value(), 64)
	fpReal, _ := strconv.ParseFloat(focalPointReal.Value(), 64)
	fpImag, _ := strconv.ParseFloat(focalPointImaginary.Value(), 64)
	VP.ImaginaryPlaneFocalPoint = complex(fpReal, fpImag)
	VP.ZoomLevel = zoom
	DrawMandelbrot(VP, i, display)
}

func main() {
	control := sync.Mutex{}

	// VP is a ViewPort that maps the pixels to the imaginary number plane C
	// Zooming and Panning are controlled by
	VP := &ViewPort{
		Image:                    ImageBoundingBox{UpperLeft: Point{0, 0}, LowerRight: Point{900, 700}},
		C:                        CBoundingBox{UpperLeft: complex(-2.5, -1.5), LowerRight: complex(1.5, 1.5)},
		ImaginaryPlaneFocalPoint: complex(0, 0),
		ZoomLevel:                1.0,
		Pallette:                 SeussyPallette,
	}

	//TODO Create application window
	Window := dali.NewWindow(1280, 920, "", "")
	Head := dali.NewHeadElement()
	title := &dali.TitleElement{Text: "Example Three: Mandelbrot Set"}
	Head.Elements.AddElement(title)
	Head.Elements.AddElement(&dali.ScriptElement{Text: `
	function draw_mandelbrot_set(){}
	async function first_view(){
		await new Promise(r => setTimeout(r, 200));
		pick_favorite_spot();
		//draw_mandelbrot_set();
	}
	function name_favorite_spot(){
		document.getElementById("viewName").value=prompt("Name This View:");
	} `})
	Window.Elements.AddElement(Head)
	Body := dali.NewBodyElement("first_view")

	div := dali.NewDiv("displayDiv")
	div.Style = `background-color:#BBBBBB;width:1260;height:900;`
	Body.Elements.AddElement(div)
	display := dali.NewImage("display", 900, 700, "")
	display.Style = `border:solid 1px #333333;display:block;margin:auto;`
	div.Elements.AddElement(display)
	div.Elements.AddElement(dali.LineBreak())

	startButton := dali.NewButton("Start Iterations", "start", "start_iterations")
	startButton.BoundFunction = func() {
		startButton.Disable()
		FindMandelbrotSet(Window, display, &control, VP)
		startButton.Enable()
		Window.GetUI().Eval(`document.getElementById("start").disabled=false;`)
	}

	pauseButton := dali.NewButton("Pause Iteration", "pause", "pause_iteration")

	toggle := make(chan bool)
	go func() {
		i := -1
		for {
			<-toggle
			i *= -1
			if i > 0 {
				control.Lock()
				pauseButton.SetText("Resume Iteration")
			} else {
				control.Unlock()
				pauseButton.SetText("Pause Iteration")
			}
		}
	}()

	pauseButton.BoundFunction = func() { toggle <- true }

	palette := dali.NewSelectElement("palette", "pick_palette")
	palette.AddOption("Dr Seussy", "1")
	palette.AddOption("Dr Lucy", "8")
	palette.AddOption("Dr Spock", "9")
	palette.AddOption("Radiation", "2")
	palette.AddOption("Iradiated", "7")
	palette.AddOption("Radiation Too", "5")
	palette.AddOption("Black Hole Sun", "3")
	palette.AddOption("Black Hole", "6")
	palette.AddOption("Color-full", "4")
	palette.AddOption("Colorful Too", "10")
	palette.AddOption("Colorful Three", "12")
	palette.AddOption("High Zoom", "11")

	zoomInButton := dali.NewButton("Zoom In", "zoomIn", "do_zoom_in")
	zoomOutButton := dali.NewButton("Zoom Out", "zoomOut", "do_zoom_out")

	zoomLevel := dali.NewInputElement("zoomLevel", dali.NumberInput)
	zoomLevel.Style = "width:14em;"
	zoomLevel.Text = "1.0"
	zoomLevel.Binding.FunctionName = "zoomLevelUpdate"

	zoomLevel.InputEventType = dali.OnBlur

	panLeftButton := dali.NewButton("Pan Left", "left", "do_pan_left")
	panRightButton := dali.NewButton("Pan Right", "right", "do_pan_right")
	panUpButton := dali.NewButton("Pan Up", "up", "do_pan_up")
	panDownButton := dali.NewButton("Pan Down", "down", "do_pan_down")

	focalPointDiv := dali.NewDiv("focalPoint")
	focalPointDiv.Style = `display:block;margin:auto;width:50%;align-content:center;align-items:center`
	focalPointReal := dali.NewInputElement("focalPointReal", dali.NumberInput)
	focalPointReal.Style = "width:15em;"
	focalPointReal.Text = fmt.Sprintf("%f", real(VP.ImaginaryPlaneFocalPoint))
	focalPointReal.Binding.FunctionName = "updateFocalPointReal"
	focalPointReal.InputEventType = dali.OnBlur
	focalPointImaginary := dali.NewInputElement("focalPointImaginary", dali.NumberInput)
	focalPointImaginary.Text = fmt.Sprintf("%f", imag(VP.ImaginaryPlaneFocalPoint))
	focalPointImaginary.Binding.FunctionName = "updateFocalPointImaginary"
	focalPointImaginary.InputEventType = dali.OnBlur
	focalPointImaginary.Style = "width:15em;"
	focalPointDiv.Elements.AddElement(&dali.Span{Text: "Focal Point: Real: "})
	focalPointDiv.Elements.AddElement(focalPointReal)
	focalPointDiv.Elements.AddElement(&dali.Span{Text: " Imaginary: "})
	focalPointDiv.Elements.AddElement(focalPointImaginary)

	iterationsDiv := dali.NewDiv("iterationMenu")
	iterationsDiv.Style = `display:block;margin:auto;width:50%;align-content:center;align-items:center`
	iterations := dali.NewInputElement("iterations", dali.NumberInput)
	zoomOutButton.BoundFunction = func() { ZoomOut(display, iterations, zoomLevel, VP, &control) }
	zoomInButton.BoundFunction = func() { ZoomIn(display, iterations, zoomLevel, VP, &control) }
	panLeftButton.BoundFunction = func() { PanLeft(display, iterations, focalPointReal, VP, &control) }
	panRightButton.BoundFunction = func() { PanRight(display, iterations, focalPointReal, VP, &control) }
	panDownButton.BoundFunction = func() { PanDown(display, iterations, focalPointImaginary, VP, &control) }
	panUpButton.BoundFunction = func() { PanUp(display, iterations, focalPointImaginary, VP, &control) }
	iterations.Text = "1000"
	iterations.Style = "width:10em;"
	iterations.Binding.FunctionName = "updateIterations"
	iterations.InputEventType = dali.OnBlur
	iterationsDiv.Elements.AddElement(&dali.Span{Text: "Iterations: "})
	iterationsDiv.Elements.AddElement(iterations)
	iterationsDiv.Elements.AddElement(startButton)
	iterationsDiv.Elements.AddElement(pauseButton)
	iterationsDiv.Elements.AddElement(palette)
	div.Elements.AddElement(iterationsDiv)

	panMenu := dali.NewDiv("menu")
	panMenu.Style = `diplay:block;margin:auto;width:37%;align-content:center;align-items:center`

	panMenu.Elements.AddElement(panLeftButton)
	panMenu.Elements.AddElement(panRightButton)
	panMenu.Elements.AddElement(panUpButton)
	panMenu.Elements.AddElement(panDownButton)
	div.Elements.AddElement(dali.LineBreak())
	div.Elements.AddElement(panMenu)
	div.Elements.AddElement(focalPointDiv)

	zoomMenu := dali.NewDiv("menu")
	zoomMenu.Style = `diplay:block;margin:auto;width:37%;align-content:center;align-items:center`
	zoomMenu.Elements.AddElement(dali.LineBreak())
	zoomMenu.Elements.AddElement(zoomInButton)
	zoomMenu.Elements.AddElement(zoomLevel)
	zoomMenu.Elements.AddElement(zoomOutButton)
	div.Elements.AddElement(zoomMenu)
	favDiv := dali.NewDiv("favDiv")
	favDiv.Style = `display:block;margin:auto;width:45%;align-items:center;align-content:center;`

	favs, err := NewFavorites()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//favs.Style = `display:block;margin:auto`
	favs.BoundFunction = func() {
		fv := favs.FavoriteSpots[favs.Value()]
		focalPointReal.Set(fmt.Sprintf("%.14f", fv.FocalPointReal))
		focalPointImaginary.Set(fmt.Sprintf("%.14f", fv.FocalPointImaginary))
		zoomLevel.Set(fmt.Sprintf("%.14f", fv.ZoomLevel))
		iterations.Set(fmt.Sprintf("%d", fv.Iterations))
		VP.ImaginaryPlaneFocalPoint = complex(fv.FocalPointReal, fv.FocalPointImaginary)
		VP.ZoomLevel = fv.ZoomLevel
		UpdateDisplay(VP, display, &control, iterations, zoomLevel, focalPointReal, focalPointImaginary)
	}
	saveButton := dali.NewButton("Add This View to Favorites", "saveButton", "saveFavorite")
	saveButton.Binding.BoundFunction = func() {
		Window.GetUI().Eval(`name_favorite_spot();`)
		favName := fmt.Sprintf("%s", Window.GetUI().Eval(`document.getElementById("viewName").value`))
		favs.AddFavoriteSpot(favName, focalPointReal, focalPointImaginary, zoomLevel, iterations)
	}
	viewName := dali.NewInputElement("viewName", dali.HiddenInput)

	favDiv.Elements.AddElement(dali.LineBreak())
	favDiv.Elements.AddElement(favs)
	favDiv.Elements.AddElement(dali.Text(" "))
	favDiv.Elements.AddElement(saveButton)
	favDiv.Elements.AddElement(viewName)
	div.Elements.AddElement(favDiv)

	Window.Elements.AddElement(Body)

	palette.Binding.BoundFunction = func() {
		v := palette.Value()
		VP.Pallette = PickPallette(v)
		UpdateDisplay(VP, display, &control, iterations, zoomLevel, focalPointReal, focalPointImaginary)
	}
	zoomLevel.Binding.BoundFunction = func() {
		UpdateDisplay(VP, display, &control, iterations, zoomLevel, focalPointReal, focalPointImaginary)
	}
	focalPointReal.Binding.BoundFunction = func() {
		UpdateDisplay(VP, display, &control, iterations, zoomLevel, focalPointReal, focalPointImaginary)
	}
	focalPointImaginary.Binding.BoundFunction = func() {
		UpdateDisplay(VP, display, &control, iterations, zoomLevel, focalPointReal, focalPointImaginary)
	}
	iterations.Binding.BoundFunction = func() {
		UpdateDisplay(VP, display, &control, iterations, zoomLevel, focalPointReal, focalPointImaginary)
	}

	Window.Start()
	Window.GetUI().Bind("draw_mandelbrot_set",
		func() {
			UpdateDisplay(VP, display, &control, iterations, zoomLevel, focalPointReal, focalPointImaginary)
		})
	<-Window.GetUI().Done()
}
