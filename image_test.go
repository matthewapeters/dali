package dali_test

import (
	"fmt"
	"testing"

	"github.com/matthewapeters/dali"
)

func TestImage(t *testing.T) {
	i := dali.NewImage("TestImageName", "TestImageID", 100, 200, "PathToImage")
	expected := `<image name="TestImageName" id="TestImageID" width="100" height="200" src="PathToImage" style="width:100;height:200;">`
	if fmt.Sprintf("%s", i) != expected {
		t.Errorf(`expected 
		"%s"
		 but got 
		"%s"`, expected, i)
	}
	i.SetStyle(`border:solid 1px #123456`)
	expected = `<image name="TestImageName" id="TestImageID" width="100" height="200" src="PathToImage" style="border:solid 1px #123456">`
	if fmt.Sprintf("%s", i) != expected {
		t.Errorf(`expected "%s" but got "%s"`, expected, i)
	}
	err := i.AddMapArea(dali.Default, dali.Coordinates{}, "altText", dali.Function, "doAltText")
	if err != nil {
		t.Errorf("%s", err)
	}
	if !i.Clickable() {
		t.Error("Expected image to be clickable after adding an area map")
		fmt.Println(i.AreaMap.Areas)
	}

	expected = `<image name="TestImageName" id="TestImageID" width="100" height="200" src="PathToImage" style="border:solid 1px #123456" usemap="#TestImageID_map"><map id="TestImageID_map"><area shape="default" coords="0,0,100,200" onclick="doAltText()" alt="altText"></map>`
	if fmt.Sprintf("%s", i) != expected {
		t.Errorf(`expected 
"%s" 
but got
"%s"`, expected, i)
	}
}
