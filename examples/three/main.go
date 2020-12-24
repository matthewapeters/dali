package main

import (
	"fmt"
	"image"
	"math"
	"math/cmplx"
	"os"
	"runtime"
	"strconv"
	"sync"

	"github.com/matthewapeters/dali"
)

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
				// Newtonian Iteration
				//2z3 - z2 - c)/(3z2 - 2z - c)
				//zn = (2*zn*zn*zn - zn*zn - z) / (3*zn*zn - 2*zn - z)

				// Traditional Mandelbrot
				// z^2 + c
				zn = zn*zn + z

				//
				//zn = complex(math.Sin(real(zn))*math.Cosh(imag(zn)), math.Cos(real(zn)*math.Sinh(imag(zn)))) + z

				//if cmplx.Abs(zn) > 0.01 {
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
func DrawMandelbrot(view *ViewPort, iterations int, display *dali.Canvas, progress *dali.ProgressElement) {

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
		go drawChunk(int(a), int(Chunk), height, iterations, view, &chunkWG, chout)
	}

	// Wait until all of the pixels are done are finished drawing
	for pixelCount := 0; pixelCount < width*height; pixelCount++ {
		i := <-chout
		if i != nil {
			image.Set(i.Pixel.X, i.Pixel.Y, view.Pallette(i))
			go progress.Status(float64(pixelCount))

		}
	}
	chunkWG.Wait()
	display.DrawImage(image, 0, 0)
	progress.Status(0.0)
}

//FindMandelbrotSet iterates over C to determine members of the Mandelbrot set
func FindMandelbrotSet(w *dali.Window, image *dali.Canvas, control *sync.Mutex, vp *ViewPort, progress *dali.ProgressElement) {
	for i := 0; i < 5000; i += 25 {
		control.Lock()
		w.GetUI().Eval(fmt.Sprintf(`document.getElementById("iterations").value="%d";`, i))
		DrawMandelbrot(
			vp,
			i,
			image,
			progress)
		control.Unlock()
	}
}

// ZoomIn will reduce the range of the C Plan within the viewport
func ZoomIn(image *dali.Canvas, iterations, zoomLevel *dali.InputElement, vp *ViewPort,
	control *sync.Mutex, progress *dali.ProgressElement) {
	control.Lock()
	defer control.Unlock()

	vp.ZoomLevel *= 0.9

	zoomLevel.Set(fmt.Sprintf("%.14f", vp.ZoomLevel))
	iv := iterations.Value()
	i, _ := strconv.Atoi((iv))
	DrawMandelbrot(
		vp,
		i,
		image,
		progress)
}

// ZoomOut will increase the range of the C Plan within the viewport
func ZoomOut(image *dali.Canvas, iterations, zoomLevel *dali.InputElement, vp *ViewPort,
	control *sync.Mutex, progress *dali.ProgressElement) {
	control.Lock()
	defer control.Unlock()

	vp.ZoomLevel *= (10.0 / 9.0)
	zoomLevel.Set(fmt.Sprintf("%.14f", vp.ZoomLevel))
	iv := iterations.Value()
	i, _ := strconv.Atoi((iv))
	DrawMandelbrot(
		vp,
		i,
		image,
		progress)
}

// PanLeft will shift the Complex Plane to the right within the viewport
func PanLeft(image *dali.Canvas, iterations, focalPointReal *dali.InputElement, vp *ViewPort,
	control *sync.Mutex, progress *dali.ProgressElement) {
	control.Lock()
	defer control.Unlock()
	length := 4 * vp.ZoomLevel
	vp.ImaginaryPlaneFocalPoint -= complex(0.1*length, 0)

	iv := iterations.Value()
	i, _ := strconv.Atoi((iv))
	DrawMandelbrot(
		vp,
		i,
		image,
		progress)

	focalPointReal.Set(fmt.Sprintf("%.14f", real(vp.ImaginaryPlaneFocalPoint)))
}

// PanRight will shift the Complex Plane to the left within the viewport
func PanRight(image *dali.Canvas, iterations, focalPointReal *dali.InputElement, vp *ViewPort,
	control *sync.Mutex, progress *dali.ProgressElement) {
	control.Lock()
	defer control.Unlock()
	length := 4 * vp.ZoomLevel
	vp.ImaginaryPlaneFocalPoint += complex(0.1*length, 0)

	iv := iterations.Value()
	i, _ := strconv.Atoi((iv))
	DrawMandelbrot(
		vp,
		i,
		image,
		progress)

	focalPointReal.Set(fmt.Sprintf("%.14f", real(vp.ImaginaryPlaneFocalPoint)))
}

// PanUp will shift the Complex Plane down within the viewport
func PanUp(image *dali.Canvas, iterations, focalPointImaginary *dali.InputElement, vp *ViewPort,
	control *sync.Mutex, progress *dali.ProgressElement) {
	control.Lock()
	length := 4.0 * vp.ZoomLevel
	vp.ImaginaryPlaneFocalPoint += complex(0, 0.1*length)
	iv := iterations.Value()
	i, _ := strconv.Atoi((iv))
	DrawMandelbrot(
		vp,
		i,
		image,
		progress)

	focalPointImaginary.Set(fmt.Sprintf("%.14f", imag(vp.ImaginaryPlaneFocalPoint)))
	control.Unlock()
}

// PanDown will shift the Complex Plane up within the viewport
func PanDown(image *dali.Canvas, iterations, focalPointImaginary *dali.InputElement, vp *ViewPort,
	control *sync.Mutex, progress *dali.ProgressElement) {
	control.Lock()
	length := 4.0 * vp.ZoomLevel
	vp.ImaginaryPlaneFocalPoint -= complex(0, 0.1*length)
	iv := iterations.Value()
	i, _ := strconv.Atoi((iv))
	DrawMandelbrot(
		vp,
		i,
		image,
		progress)

	focalPointImaginary.Set(fmt.Sprintf("%.14f", imag(vp.ImaginaryPlaneFocalPoint)))
	control.Unlock()
}

//UpdateDisplay will redraw the Mandelbrot set based on Window values
func UpdateDisplay(VP *ViewPort, display *dali.Canvas, control *sync.Mutex, iterations,
	zoomLevel, focalPointReal, focalPointImaginary *dali.InputElement, progress *dali.ProgressElement) {
	control.Lock()
	defer control.Unlock()
	i, _ := strconv.Atoi(iterations.Value())
	zoom, _ := strconv.ParseFloat(zoomLevel.Value(), 64)
	fpReal, _ := strconv.ParseFloat(focalPointReal.Value(), 64)
	fpImag, _ := strconv.ParseFloat(focalPointImaginary.Value(), 64)
	VP.ImaginaryPlaneFocalPoint = complex(fpReal, fpImag)
	VP.ZoomLevel = zoom

	DrawMandelbrot(VP, i, display, progress)
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
	progress := dali.NewProgressElement("redrawProgressBarr", "redrawProgressBarr", 900*700)

	Window := dali.NewWindow(1280, 920, "", "")
	Head := dali.NewHeadElement()
	title := &dali.TitleElement{Text: "Example Three: Mandelbrot Set"}
	Head.Elements.AddElement(title)
	Head.Elements.AddElement(&dali.ScriptElement{Text: `
	function draw_mandelbrot_set(){}
	async function body_on_load(){
		await new Promise(r => setTimeout(r, 200));
		pick_favorite_spot();
		//draw_mandelbrot_set();
	}
	function name_favorite_spot(){
		document.getElementById("viewName").value=prompt("Name This View:");
	} `})
	Window.Elements.AddElement(Head)
	Body := dali.NewBodyElement("first_view")

	div := dali.NewDiv("displayDiv", "displayDiv")
	div.SetStyle(`background-color:#BBBBBB;width:1260;height:900;`)

	display := dali.NewCanvas("display", "display", 900, 700)
	display.SetStyle(`border:solid 1px #333333;display:block;margin:auto;`)
	div.Elements.AddElement(display)

	tabl := dali.NewTableElement("menus", "menus", 3, 3, []string{"", "Explore the Mandelbrot Set", ""})
	tabl.SetStyle("width:100%;padding:5px;")
	tabl.SetCommonStyles("padding:0px;margin:none;")
	a, _ := tabl.GetCell(0, 0)
	a.SetStyle("width:33%;")

	startButton := dali.NewButton("Start Iterations", "start", "start", "start_iterations")
	startButton.SetBoundFunction(dali.ClickEvent, func() {
		startButton.Disable()
		FindMandelbrotSet(Window, display, &control, VP, progress)
		startButton.Enable()
		Window.GetUI().Eval(`document.getElementById("start").disabled=false;`)
	})

	pauseButton := dali.NewButton("Pause Iteration", "pause", "pause", "pause_iteration")

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

	pauseButton.SetBoundFunction(dali.ClickEvent, func() { toggle <- true })

	palette := dali.NewSelectElement("palette", "palette", "pick_palette")
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

	zoomInButton := dali.NewButton("Zoom In", "zoomIn", "zoomIn", "do_zoom_in")
	zoomOutButton := dali.NewButton("Zoom Out", "zoomOut", "zoomOut", "do_zoom_out")

	zoomLevel := dali.NewInputElement("zoomLevel", "zoomLevel", dali.NumberInput)
	zoomLevel.SetStyle("width:14em;")
	zoomLevel.Text = "1.0"

	panLeftButton := dali.NewButton("Pan Left", "left", "left", "do_pan_left")
	panRightButton := dali.NewButton("Pan Right", "right", "right", "do_pan_right")
	panUpButton := dali.NewButton("Pan Up", "up", "up", "do_pan_up")
	panDownButton := dali.NewButton("Pan Down", "down", "down", "do_pan_down")
	panTable := dali.NewTableElement("panTable", "panTable", 3, 3, []string{})
	panTable.SetCommonStyles("align-items:center;text-align:center;")
	panCell, _ := panTable.GetCell(1, 0)
	panCell.Elements.AddElement(panUpButton)
	panCell, _ = panTable.GetCell(0, 1)
	panCell.Elements.AddElement(panLeftButton)
	panCell, _ = panTable.GetCell(2, 1)
	panCell.Elements.AddElement(panRightButton)
	panCell, _ = panTable.GetCell(1, 1)
	panCell.Elements.AddElement(panDownButton)

	focalPointTable := dali.NewTableElement("focalPointTable", "focalPointTable", 2, 2, []string{})
	focalPointTable.SetStyleProperty(dali.Float, "right")
	focalPointReal := dali.NewInputElement("focalPointReal", "focalPointReal", dali.NumberInput)
	focalPointReal.SetStyle("width:15em;")
	focalPointReal.Text = fmt.Sprintf("%f", real(VP.ImaginaryPlaneFocalPoint))
	focalPointReal.InputEventType = dali.OnBlur
	focalPointImaginary := dali.NewInputElement("focalPointImaginary", "focalPointImaginary", dali.NumberInput)
	focalPointImaginary.Text = fmt.Sprintf("%f", imag(VP.ImaginaryPlaneFocalPoint))
	focalPointImaginary.InputEventType = dali.OnBlur
	focalPointImaginary.SetStyle("width:15em;")
	fpCell, _ := focalPointTable.GetCell(0, 0)
	fpCell.Elements.AddElement(dali.NewSpanElement("realLable", "realLabel", "Focal Point: Real: "))
	fpCell, _ = focalPointTable.GetCell(1, 0)
	fpCell.Elements.AddElement(focalPointReal)
	fpCell, _ = focalPointTable.GetCell(0, 1)
	fpCell.Elements.AddElement(dali.NewSpanElement("imagLabel", "imagLabel", " Imaginary: "))
	fpCell, _ = focalPointTable.GetCell(1, 1)
	fpCell.Elements.AddElement(focalPointImaginary)

	iterationsDiv := dali.NewDiv("iterationMenu", "iterationMenu")
	iterations := dali.NewInputElement("iterations", "iterations", dali.NumberInput)
	zoomOutButton.SetBoundFunction(dali.ClickEvent, func() { ZoomOut(display, iterations, zoomLevel, VP, &control, progress) })
	zoomInButton.SetBoundFunction(dali.ClickEvent, func() { ZoomIn(display, iterations, zoomLevel, VP, &control, progress) })
	panLeftButton.SetBoundFunction(dali.ClickEvent, func() { PanLeft(display, iterations, focalPointReal, VP, &control, progress) })
	panRightButton.SetBoundFunction(dali.ClickEvent, func() { PanRight(display, iterations, focalPointReal, VP, &control, progress) })
	panDownButton.SetBoundFunction(dali.ClickEvent, func() { PanDown(display, iterations, focalPointImaginary, VP, &control, progress) })
	panUpButton.SetBoundFunction(dali.ClickEvent, func() { PanUp(display, iterations, focalPointImaginary, VP, &control, progress) })
	iterations.Text = "1000"
	iterations.SetStyle("width:10em;")
	iterationsDiv.Elements.AddElement(dali.NewSpanElement("iterLabel", "iterLabel", "Iterations: "))
	iterationsDiv.Elements.AddElement(iterations)
	iterationsDiv.Elements.AddElement(dali.NewBreak())
	iterationsDiv.Elements.AddElement(startButton)
	iterationsDiv.Elements.AddElement(pauseButton)

	c, _ := tabl.GetCell(0, 0)
	c.Elements.AddElement(iterationsDiv)
	c, _ = tabl.GetCell(1, 0)
	c.Elements.AddElement(dali.NewSpanElement("palLabel", "palLabel", "Palette: "))
	c.Elements.AddElement(palette)

	c, _ = tabl.GetCell(1, 1)
	c.SetStyle(c.Style() + ";text-align:center;item-align:center")
	//c.Elements.AddElement(dali.NewBreak())
	panTable.SetStyle("width:100%;self-align:center;border:solid 1px")
	c.Elements.AddElement(panTable)
	c, _ = tabl.GetCell(2, 0)
	c.SetStyle(c.Style() + ";item-align:right;text-align:right;")
	c.Elements.AddElement(focalPointTable)

	zoomMenu := dali.NewDiv("zoomMenu", "zoomMenu")
	zoomMenu.Elements.AddElement(zoomInButton)
	zoomMenu.Elements.AddElement(zoomLevel)
	zoomMenu.Elements.AddElement(zoomOutButton)
	c, _ = tabl.GetCell(1, 2)
	c.SetStyle("align:center")
	c.Elements.AddElement(zoomMenu)

	favs, err := NewFavorites()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	favs.SetBoundFunction(dali.ChangeEvent, func() {
		fv := favs.FavoriteSpots[favs.Value()]
		focalPointReal.Set(fmt.Sprintf("%.14f", fv.FocalPointReal))
		focalPointImaginary.Set(fmt.Sprintf("%.14f", fv.FocalPointImaginary))
		zoomLevel.Set(fmt.Sprintf("%.14f", fv.ZoomLevel))
		iterations.Set(fmt.Sprintf("%d", fv.Iterations))
		VP.ImaginaryPlaneFocalPoint = complex(fv.FocalPointReal, fv.FocalPointImaginary)
		VP.ZoomLevel = fv.ZoomLevel
		UpdateDisplay(VP, display, &control, iterations, zoomLevel, focalPointReal, focalPointImaginary, progress)
	})
	saveButton := dali.NewButton("Save View", "saveButton", "saveButton", "saveFavorite")
	saveButton.SetBoundFunction(dali.ClickEvent, func() {
		Window.GetUI().Eval(`name_favorite_spot();`)
		favName := fmt.Sprintf("%s", Window.GetUI().Eval(`document.getElementById("viewName").value`))
		favs.AddFavoriteSpot(favName, focalPointReal, focalPointImaginary, zoomLevel, iterations)
	})
	viewName := dali.NewInputElement("viewName", "viewName", dali.HiddenInput)

	c, _ = tabl.GetCell(0, 1)
	c.SetStyle("text-align:left")
	c.Elements.AddElement(dali.NewSpanElement("favLabel", "favLabel", "Favorites"))
	c.Elements.AddElement(dali.NewBreak())
	c.Elements.AddElement(favs)
	c.Elements.AddElement(saveButton)
	c.Elements.AddElement(viewName)
	c.Elements.AddElement(dali.NewBreak())
	c.Elements.AddElement(dali.NewBreak())

	c.Elements.AddElement(dali.NewSpanElement("rendrLabel", "rendrLabel", "Render Progress:"))
	c.Elements.AddElement(dali.NewBreak())
	c.Elements.AddElement(progress)

	div.Elements.AddElement(tabl)
	Body.Elements.AddElement(div)
	Window.Elements.AddElement(Body)

	palette.SetBoundFunction(dali.ChangeEvent, func() {
		v := palette.Value()
		VP.Pallette = PickPallette(v)
		UpdateDisplay(VP, display, &control, iterations, zoomLevel, focalPointReal, focalPointImaginary, progress)
	})

	zoomLevel.SetBoundFunction(dali.ChangeEvent, func() {
		UpdateDisplay(VP, display, &control, iterations, zoomLevel, focalPointReal, focalPointImaginary, progress)
	})
	focalPointReal.SetBoundFunction(dali.ChangeEvent, func() {
		UpdateDisplay(VP, display, &control, iterations, zoomLevel, focalPointReal, focalPointImaginary, progress)
	})
	focalPointImaginary.SetBoundFunction(dali.ChangeEvent, func() {
		UpdateDisplay(VP, display, &control, iterations, zoomLevel, focalPointReal, focalPointImaginary, progress)
	})
	iterations.SetBoundFunction(dali.ChangeEvent, func() {
		UpdateDisplay(VP, display, &control, iterations, zoomLevel, focalPointReal, focalPointImaginary, progress)
	})

	Window.Start()
	Window.GetUI().Bind("draw_mandelbrot_set",
		func() {
			UpdateDisplay(VP, display, &control, iterations, zoomLevel, focalPointReal, focalPointImaginary, progress)
		})
	<-Window.GetUI().Done()
}
