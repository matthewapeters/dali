package dali

import (
	"fmt"
	"testing"
)

func TestImage(t *testing.T) {
	i := NewImage("TestImageName", 100, 200, "PathToImage")
	expected := `<image name="TestImageName" width="100" height="200" src="PathToImage">`
	if fmt.Sprintf("%s", i) != expected {
		t.Errorf(`expected "%s" but got "%s"`, expected, i)
	}
	i.StyleName = `border:solid 1px #123456`
	expected = `<image name="TestImageName" width="100" height="200" src="PathToImage" style="border:solid 1px #123456">`
	if fmt.Sprintf("%s", i) != expected {
		t.Errorf(`expected "%s" but got "%s"`, expected, i)
	}
	err := i.AddMapArea(Default, Coordinates{}, "altText", Function, "doAltText")
	if err != nil {
		t.Errorf("%s", err)
	}
	if !i.Clickable() {
		t.Error("Expected image to be clickable after adding an area map")
		fmt.Println(i.AreaMap.Areas)
	}

	expected = `<image name="TestImageName" width="100" height="200" src="PathToImage" style="border:solid 1px #123456"usemap="#TestImageName_map"><map name="TestImageName_map"><area shape="default" coords="0,0,100,200" onclick="doAltText()" alt="altText"></map>`
	if fmt.Sprintf("%s", i) != expected {
		t.Errorf(`expected "%s" but got "%s"`, expected, i)
	}
}
