package main

import "image/color"

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
