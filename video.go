package dali

import (
	"errors"
	"fmt"
)

//Video provides a video element
type Video struct {
	Base
	Width, Height int
	Elements      *Elements
}

// NewVideoElement creates a new Video element
func NewVideoElement(name, id string, width, height int) *Video {
	return &Video{
		Base: Base{ElementID: id, ElementName: name}, Elements: &Elements{}, Width: width, Height: height}
}

func (v *Video) String() string {
	style := ""
	if v.Style() == "" {
		style = fmt.Sprintf(` style=width: %d px; height: %dpx;`, v.Width, v.Height)
	} else {
		style = fmt.Sprintf(` style="%s;width: %d px; height: %dpx;"`, v.Style(), v.Width, v.Height)
	}

	return fmt.Sprintf(`<video id="%s" autoplay %s>
	<script> <!--
		var constraints = { audio: true, video: { width: %d, height: %d } };

		async function %s_startTracks(){
			navigator.mediaDevices.getUserMedia(constraints)
				.then(function(mediaStream) {
					var video = document.querySelector('video');
					video.srcObject = mediaStream;
					video.onloadedmetadata = function(e) {
						video.play();
					};
				})
				.catch(function(err) { console.log(err.name + ": " + err.message); } // always check for errors at the end.
				); 
		}

		async function %s_stopTracks(){
			var s=document.getElementById("%s").captureStream();
			s.getTracks()[0].stop();
			s.getTracks()[1].stop();
		}
	--></script>
	</video>`, v.ID(), style, v.Width, v.Height, v.ID(), v.ID(), v.ID())
}

// Children returns the child elements
func (v *Video) Children() *Elements { return v.Elements }

//StartTracks will start the camera and audio streams
func (v *Video) StartTracks() error {
	var err error
	e := (*v.GetUI()).Eval(fmt.Sprintf(`%s_startTracks();`, v.ID()))
	if e != nil {
		err = fmt.Errorf(fmt.Sprintf(`%s`, e))
	}
	return err
}

//StopTracks will terminate the camera and audio stream
func (v *Video) StopTracks() error {
	if v.GetUI() == nil {
		return errors.New("Window is not yet started")
	}
	var err error
	e := (*v.GetUI()).Eval(fmt.Sprintf(`%s_stopTracks(); `, v.ID()))
	if e != nil {
		err = fmt.Errorf(fmt.Sprintf(`%s`, e))
	}
	return err
}
