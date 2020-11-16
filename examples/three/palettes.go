package main

import (
	"image/color"
	"math"

	colorful "github.com/lucasb-eyer/go-colorful"
)

//GlowPallette returns a glowy-pallette color
func GlowPallette(i *Iterant) color.Color {
	//hue := math.Mod(float64(i.N), 1000) / 1000.0 * 360.0
	maxIterations := float64(i.Threshold)
	currentIteration := float64(i.N)
	logIterationRatio := math.Log10(currentIteration) / math.Log10(maxIterations)

	hue := logIterationRatio * 360.0
	value := math.Pow(math.Cos(math.Pi*logIterationRatio/2.0), 2.0)

	render := colorful.Hsv(hue, value, (1 - value))
	R, G, B := render.RGB255()
	return color.RGBA{R, G, B, 255}
}

//Glow3Pallette returns a glowy-pallette color
func Glow3Pallette(i *Iterant) color.Color {
	//hue := math.Mod(float64(i.N), 1000) / 1000.0 * 360.0
	maxIterations := float64(i.Threshold)
	currentIteration := float64(i.N)
	logIterationRatio := math.Log10(currentIteration) / math.Log10(maxIterations)

	hue := logIterationRatio * 360.0
	value := math.Pow(math.Cos(math.Pi*logIterationRatio/2.0), 2.0)

	render := colorful.Hsv(360-hue, value, (1 - value))
	R, G, B := render.RGB255()
	return color.RGBA{R, G, B, 255}
}

//Glow2Pallette returns a glowy-pallette color
func Glow2Pallette(i *Iterant) color.Color {
	//hue := math.Mod(float64(i.N), 1000) / 1000.0 * 360.0
	maxIterations := float64(i.Threshold)
	currentIteration := float64(i.N)
	logIterationRatio := math.Log10(currentIteration) / math.Log10(maxIterations)

	hue := logIterationRatio * 360.0
	value := math.Pow(math.Cos(math.Pi*logIterationRatio/2.0), 2.0)

	render := colorful.Hsv(hue, 1-value, (1 - value))
	R, G, B := render.RGB255()
	return color.RGBA{R, G, B, 255}
}

//DarkPallette returns a glowy-pallette color
func DarkPallette(i *Iterant) color.Color {
	//hue := math.Mod(float64(i.N), 1000) / 1000.0 * 360.0
	maxIterations := float64(i.Threshold)
	currentIteration := float64(i.N)
	logIterationRatio := math.Log10(currentIteration) / math.Log10(maxIterations)

	hue := logIterationRatio * 360.0
	value := math.Pow(math.Cos(math.Pi*logIterationRatio/2.0), 2.0)

	render := colorful.Hsv(hue, 0, value)
	R, G, B := render.RGB255()
	return color.RGBA{R, G, B, 255}
}

//Dark2Pallette returns a glowy-pallette color
func Dark2Pallette(i *Iterant) color.Color {
	//hue := math.Mod(float64(i.N), 1000) / 1000.0 * 360.0
	maxIterations := float64(i.Threshold)
	currentIteration := float64(i.N)
	logIterationRatio := math.Log(currentIteration) / math.Log(maxIterations)

	hue := logIterationRatio * 360.0
	value := math.Pow(math.Cos(math.Pi*logIterationRatio/2.0), 2.0)

	render := colorful.Hsv(360-hue, value, value)
	R, G, B := render.RGB255()
	return color.RGBA{R, G, B, 255}
}

//ColorfulPallette returns a glowy-pallette color
func ColorfulPallette(i *Iterant) color.Color {
	//hue := math.Mod(float64(i.N), 1000) / 1000.0 * 360.0
	maxIterations := float64(i.Threshold)
	currentIteration := float64(i.N)
	logIterationRatio := math.Log10(currentIteration) / math.Log10(maxIterations)

	hue := logIterationRatio * 360.0
	value := math.Pow(math.Cos(math.Pi*logIterationRatio/2.0), 2.0)

	render := colorful.Hsv(hue, 1, value)
	R, G, B := render.RGB255()
	return color.RGBA{R, G, B, 255}
}

//Colorful2Pallette returns a glowy-pallette color
func Colorful2Pallette(i *Iterant) color.Color {
	//hue := math.Mod(float64(i.N), 1000) / 1000.0 * 360.0
	maxIterations := float64(i.Threshold)
	currentIteration := float64(i.N)
	logIterationRatio := math.Log10(currentIteration) / math.Log10(maxIterations)

	hue := logIterationRatio * 360.0
	value := math.Pow(math.Cos(math.Pi*logIterationRatio/2.0), 2.0)

	render := colorful.Hsv(hue, logIterationRatio, value)
	R, G, B := render.RGB255()
	return color.RGBA{R, G, B, 255}
}

//Colorful3Pallette returns a glowy-pallette color
func Colorful3Pallette(i *Iterant) color.Color {
	//hue := math.Mod(float64(i.N), 1000) / 1000.0 * 360.0
	maxIterations := float64(i.Threshold)
	currentIteration := float64(i.N)
	firstRenderedIteration := 0.2 * maxIterations
	mappedIteration := 1.25 * (currentIteration - firstRenderedIteration)
	sqrtIterationRatio := math.Sqrt(mappedIteration) / math.Sqrt(maxIterations)

	hue := sqrtIterationRatio * 360.0
	//value := math.Pow(math.Cos(math.Pi*sqrtIterationRatio/2.0), 2.0)
	value := (1.0 - math.Abs(hue-180.0)/180.0)

	render := colorful.Hsv(hue, 1, value)
	R, G, B := render.RGB255()
	return color.RGBA{R, G, B, 255}
}

//Colorful4Pallette returns a glowy-pallette color
func Colorful4Pallette(i *Iterant) color.Color {
	//hue := math.Mod(float64(i.N), 1000) / 1000.0 * 360.0
	maxIterations := float64(i.Threshold)
	currentIteration := float64(i.N)
	IterationRatio := math.Sqrt(currentIteration) / math.Sqrt(maxIterations)

	hue := IterationRatio * 360.0
	value := math.Pow(math.Cos(math.Pi*IterationRatio/2.0), 2.0)

	render := colorful.Hsv(hue, 1, value)
	R, G, B := render.RGB255()
	return color.RGBA{R, G, B, 255}
}

// SeussyPallette is the original pallette I came up with - quite stripey
func SeussyPallette(i *Iterant) color.Color {
	ratio := (math.Mod(float64(i.N), 1000.0) / 1000.0)
	R := ratio * 255.0
	G := math.Mod(R*255.0, 255.0)
	B := math.Mod(R*255.0*255.0, 255.0*255.0)
	return color.RGBA{uint8(R), uint8(G), uint8(B), 255}
}

// Seussy2Pallette is the original pallette I came up with - quite stripey
func Seussy2Pallette(i *Iterant) color.Color {
	ratio := (math.Mod(float64(i.N), 1000.0) / 1000.0)
	R := ratio * 255.0
	G := math.Mod(R*255.0, 255.0)
	B := math.Mod(R*255.0*255.0, 255.0*255.0)
	return color.RGBA{uint8(B), uint8(G), uint8(R), 255}
}

// Seussy3Pallette is the original pallette I came up with - quite stripey
func Seussy3Pallette(i *Iterant) color.Color {
	ratio := (math.Mod(float64(i.N), 1000.0) / 1000.0)
	R := ratio * 255.0
	G := math.Mod(R*255.0, 255.0)
	B := math.Mod(R*255.0*255.0, 255.0*255.0)
	return color.RGBA{uint8(B), uint8(R), uint8(G), 255}
}

//PickPallette returns a pallette based on a value
func PickPallette(value string) func(*Iterant) color.Color {
	switch value {
	case "1":
		return SeussyPallette
	case "2":
		return GlowPallette
	case "3":
		return DarkPallette
	case "4":
		return ColorfulPallette
	case "5":
		return Glow2Pallette
	case "6":
		return Dark2Pallette
	case "7":
		return Glow3Pallette
	case "8":
		return Seussy2Pallette
	case "9":
		return Seussy3Pallette
	case "10":
		return Colorful2Pallette
	case "11":
		return Colorful3Pallette
	case "12":
		return Colorful4Pallette
	default:
		return SeussyPallette
	}
}
