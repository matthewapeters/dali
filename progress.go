package dali

import (
	"fmt"
	"sync"
)

//ProgressElement is a progress bar
type ProgressElement struct {
	Base
	Max             float64
	CurrentValue    float64
	ProgressChannel *chan float64
	Mu              sync.Mutex
}

//NewProgressElement creates a new progress bar
func NewProgressElement(name string, max float64) *ProgressElement {

	return &ProgressElement{
		Base: Base{ID: name},
		Max:  max,
	}
}

//Children returns an empty slice of Element for ProgressElement
func (p *ProgressElement) Children() *Elements {
	return &Elements{}
}

//Value returs the current value of the ProgressElement (as found in the DOM)
func (p *ProgressElement) Value() string {
	return fmt.Sprintf(`%s`, (*p.GetUI()).Eval(fmt.Sprintf(`document.getElementById("%s").value`, p.Name())))
}

func (p *ProgressElement) String() string {
	style := ""
	if p.Style != "" {
		style = fmt.Sprintf(` style="%s"`, p.Style)
	}
	return fmt.Sprintf(`<progress id="%s" max="%f"%s></progress>`, p.Name(), p.Max, style)
}

func (p *ProgressElement) monitorProgressChannel() {
	for {
		v, ok := <-*p.ProgressChannel
		if !ok {
			return
		}
		p.Mu.Lock()
		if v >= p.CurrentValue {
			p.CurrentValue = v
			(*p.GetUI()).Eval(fmt.Sprintf(`document.getElementById("%s").value=%f`, p.Name(), p.CurrentValue))
		}
		p.Mu.Unlock()
	}
}

//Reset will return the progress bar to 0
func (p *ProgressElement) Reset() {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.CurrentValue = 0
	(*p.GetUI()).Eval(fmt.Sprintf(`document.getElementById("%s").value=0`, p.Name()))
	for {
		_, notDone := <-*p.ProgressChannel
		if !notDone {
			return
		}
	}
}

//Channel provides the progress channel as a send-only unidirectional interface
func (p *ProgressElement) Channel() chan<- float64 {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	if p.ProgressChannel != nil {
		p.Mu.Unlock()
		p.Reset()
		p.Mu.Lock()
	}
	pc := make(chan float64, int(p.Max+1))
	p.ProgressChannel = &pc

	go p.monitorProgressChannel()
	return *p.ProgressChannel
}
