package main

import (
	"fmt"
	"math/rand"

	"github.com/matthewapeters/dali"
)

//IpsumLorem returns Ipsum Lorem text with line breaks
func IpsumLorem() string {
	return `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum fermentum turpis eros, id ornare nunc vulputate ac. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus bibendum metus risus, non condimentum nibh posuere eu. Fusce pharetra magna non tincidunt luctus. Proin eu efficitur lectus, eget lobortis lorem. Etiam tristique nunc eget mi eleifend mollis. Phasellus eu augue ipsum. Praesent lobortis, libero quis pharetra mattis, arcu velit dictum eros, et bibendum metus enim ut orci. Ut id nisi tempus, molestie ipsum eu, porttitor ipsum. Etiam dignissim et nibh a faucibus. Nulla nec velit nulla. Nullam nibh eros, ornare ac orci eget, ultricies vulputate erat.
<br/><br/>
Proin pharetra suscipit mauris a auctor. Donec placerat dui non diam fermentum, in tempus velit euismod. Sed vel diam at ex elementum mattis. In sed laoreet ipsum. Nunc posuere, nisi at pellentesque gravida, massa nisl lobortis nibh, fringilla aliquet diam dui non nunc. Praesent sem libero, tincidunt eu eros nec, bibendum placerat justo. Nulla eget viverra nibh. Praesent arcu lorem, posuere vitae faucibus ut, sollicitudin ac mauris. Nullam a tempor est. Mauris blandit sagittis tincidunt. Ut sollicitudin ut augue ac commodo.
<br/><br/>
Nullam non mauris eu urna blandit consectetur vitae et leo. Nulla ultrices consectetur mauris a consequat. Vivamus id dolor iaculis, posuere est in, iaculis nibh. Curabitur nec tellus rutrum metus feugiat mollis. Etiam ac finibus magna. Vivamus bibendum arcu euismod neque varius, in posuere arcu vehicula. Pellentesque lacinia, urna at mattis ultrices, lorem nibh mattis elit, in tincidunt metus felis vitae enim. Donec at blandit lorem. Curabitur nec est in diam accumsan fermentum at sed sem.
<br/><br/>
Nulla sapien enim, pharetra sed rhoncus eu, gravida pellentesque neque. Nam enim dui, faucibus vitae erat non, egestas semper elit. Sed venenatis luctus massa nec elementum. Quisque in dui sem. Etiam vehicula convallis sapien a facilisis. Quisque mattis sapien arcu, ac sollicitudin urna dictum vitae. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam massa justo, volutpat eget lobortis ac, blandit sagittis leo. Mauris facilisis dolor a felis volutpat, vel faucibus erat dictum. Suspendisse ullamcorper, odio a rutrum sollicitudin, orci mauris ultricies lectus, at tempor leo orci eu lorem. Nam quam eros, facilisis et arcu sed, varius malesuada lectus.
<br/><br/>
Morbi cursus eu ex quis maximus. Cras viverra nisl in sapien tempus accumsan. Morbi mollis est nec leo feugiat lobortis. Sed non nunc sit amet est elementum vehicula. Cras rhoncus erat ut enim accumsan, viverra tempus sapien posuere. Phasellus condimentum quam et porttitor mollis. Curabitur ut elit volutpat, rhoncus justo egestas, finibus enim. Aliquam accumsan, arcu sed tincidunt ultricies, metus mi dapibus libero, sed tincidunt nunc turpis in arcu. Praesent sed tellus eu eros interdum congue. Cras eu libero viverra urna tempus mollis sed sed massa.`
}

// PickARandomNumber picks a randim integer server-side, and modifies a client-side element's inner HTML
// This is done through lorca's bindings
func PickARandomNumber(w *dali.Window) {
	scriptlet := ` document.getElementById("randomNumber").innerHTML="%d"; `
	randomNumber := rand.Int()
	fmt.Printf("server picked %d \n", randomNumber)
	w.GetUI().Eval(fmt.Sprintf(scriptlet, randomNumber))
}

func main() {
	/**
	 * Example one shows the creation of a window with a title.
	 */
	Window := dali.NewWindow(600, 400, "", "")
	Title := &dali.TitleElement{Text: "Dali Example One"}
	Head := dali.NewHeadElement()
	Head.Elements.AddElement(Title)
	Window.Elements.AddElement(Head)

	/*Add a body with no on-load function, and a blue background*/
	Body := dali.NewBodyElement("")
	Body.Style = "background:#5080FF;"

	/*Add an H1 banner named pageBanner to the GUI*/
	Banner := dali.NewHeader(dali.H1, "pageBanner", "This is Example One")
	Body.Elements.AddElement(Banner)
	Body.Elements.AddElement(dali.LineBreak())

	/*Add some exposition*/
	ScrollableDiv := dali.NewDiv("")
	ScrollableDiv.Style = "overflow-y:scroll;height:100;border:solid 1px #000011;padding:3"
	blahBlah := dali.Text(IpsumLorem())
	ScrollableDiv.Elements.AddElement(blahBlah)
	Body.Elements.AddElement(ScrollableDiv)
	Body.Elements.AddElement(dali.LineBreak())

	/*Demonstate Promise bindings by having a button trigger a random number selection*/
	nbr := dali.Span{Base: dali.Base{ID: "randomNumber", Style: "width: 100%;padding:5;"}, Text: "Click button to pick a random number..."}
	picker := dali.NewButton("Pick A Number:", "picker", "do_pick_a_number_server_side")
	picker.Binding.BoundFunction = func() { PickARandomNumber(Window) }
	pickerDiv := dali.NewDiv("pickerDiv")
	pickerDiv.Elements.AddElement(picker)
	pickerDiv.Elements.AddElement(&nbr)
	Body.Elements.AddElement(pickerDiv)

	Window.Elements.AddElement(Body)

	/*Start the GUI*/
	Window.Start()

	/*Wait for the GUI to close*/
	<-Window.GetUI().Done()
	fmt.Println("The GUI has closed.")
}
