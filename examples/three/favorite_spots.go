package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/matthewapeters/dali"
)

// FavoriteView is the location of a favorite view
type FavoriteView struct {
	FocalPointReal      float64 `json:"focal_point_real"`
	FocalPointImaginary float64 `json:"focal_point_imaginary"`
	ZoomLevel           float64 `json:"zoom_level"`
	Iterations          int     `json:"iterations"`
}

// FavoriteSpot identifies a location and view of the Mandelbrot Set that the user has found worth revisiting.
type FavoriteSpot struct {
	FavoriteView
	dali.Base
}

// Favorites is a list of FavoriteSpots
type Favorites struct {
	FavoriteSpots map[string]FavoriteSpot
	*dali.SelectElement
}

//GetFavoriteViews returns the Favorite Views from the Favorite Spots
func (favs *Favorites) GetFavoriteViews() *map[string]FavoriteView {
	fv := map[string]FavoriteView{}
	for name, fs := range favs.FavoriteSpots {
		fv[name] = fs.FavoriteView
	}
	return &fv
}

//Save saves the Favorites to a file
func (favs *Favorites) Save() error {
	var err error
	f, err := os.Create("favorites.json")
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := json.Marshal(favs.GetFavoriteViews())
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

//NewFavorites will load all of the favorites from a file
func NewFavorites() (*Favorites, error) {
	var err error
	favs := &Favorites{
		SelectElement: dali.NewSelectElement("favorites", "favorites", "pick_favorite_spot"),
		FavoriteSpots: map[string]FavoriteSpot{},
	}

	f, err := os.Open("favorites.json")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()
	byts, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if len(byts) != 0 {
		newFavs := map[string]FavoriteSpot{}

		err := json.Unmarshal(byts, &newFavs)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		favs.FavoriteSpots = newFavs
		idx := 0
		for name := range favs.FavoriteSpots {
			favs.AddOption(strings.ReplaceAll(name, "_", " "), name)
			idx++
		}
	} else {
		fmt.Println("no favorites to load")
	}
	return favs, err
}

//PickFavoriteSpot moves the display image to the indicated spot, zoom, and iterations
func PickFavoriteSpot(display *dali.Image, focalPointReal, focalPointImaginary, zoomLevel, iterations *dali.InputElement, fav *FavoriteView) error {
	var err error
	focalPointReal.Set(fmt.Sprintf("%.14f", fav.FocalPointReal))
	focalPointImaginary.Set(fmt.Sprintf("%.14f", fav.FocalPointImaginary))
	zoomLevel.Set(fmt.Sprintf("%.14f", fav.ZoomLevel))
	iterations.Set(fmt.Sprintf("%d", fav.Iterations))
	return err
}

//AddFavoriteSpot Add a new FavoriteSpot to Favorites
func (favs *Favorites) AddFavoriteSpot(name string, focalPointReal, focalPointImaginary, zoomLevel, iterations *dali.InputElement) {
	name = strings.ReplaceAll(name, " ", "_")
	fpr, _ := strconv.ParseFloat(focalPointReal.Value(), 64)
	fpi, _ := strconv.ParseFloat(focalPointImaginary.Value(), 64)
	itr, _ := strconv.ParseInt(iterations.Value(), 10, 32)
	zl, _ := strconv.ParseFloat(zoomLevel.Value(), 64)
	newFav := FavoriteSpot{
		Base: dali.Base{
			ElementID:   name,
			ElementName: name,
		},
		FavoriteView: FavoriteView{
			FocalPointReal:      fpr,
			FocalPointImaginary: fpi,
			Iterations:          int(itr),
			ZoomLevel:           zl,
		},
	}

	//fmt.Printf("%s", &newFav)
	favs.FavoriteSpots[name] = newFav
	(*favs.GetUI()).Eval(fmt.Sprintf(` document.getElementById("%s").options.add(new Option("%s","%s"));`,
		favs.ID,
		strings.ReplaceAll(name, "_", " "),
		name))
	favs.Save()
}
