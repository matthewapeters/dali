package dali

import (
	"fmt"
	"math"
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

	c := make(chan float64, int(max))

	return &ProgressElement{
		Base:            Base{ID: name},
		Max:             max,
		CurrentValue:    0,
		Mu:              sync.Mutex{},
		ProgressChannel: &c,
	}
}

//Children returns an empty slice of Element for ProgressElement
func (p *ProgressElement) Children() *Elements {
	return &Elements{}
}

//Value returs the current value of the ProgressElement (as found in the DOM)
func (p *ProgressElement) Value() string {
	return fmt.Sprintf(`%s`, (*p.GetUI()).Eval(fmt.Sprintf(`document.getElementById("%s").value`, p.ID()())))
}

func (p *ProgressElement) String() string {
	go p.monitorProgressChannel()
	style := ""
	if p.Style != "" {
		style = fmt.Sprintf(` style="%s"`, p.Style)
	}
	return fmt.Sprintf(`<progress id="%s" value="%f" max="%f"%s></progress>`, p.ID()(), p.CurrentValue, p.Max, style)
}

// monitors a progress channel.  If the channel is closed, returns.
func (p *ProgressElement) monitorProgressChannel() {
	for {
		v, ok := <-*p.ProgressChannel
		if !ok {
			return
		}
		p.Mu.Lock()
		p.CurrentValue = v
		p.Set(fmt.Sprintf("%f", p.CurrentValue))
		p.Mu.Unlock()
	}
}

// Status receives the status value to report in the progress bar
func (p *ProgressElement) Status(f float64) {
	if p.Max >= 1000.0 && math.Mod(f, p.Max/100.0) <= 0.5 || p.Max < 1000.0 {
		p.Mu.Lock()
		defer p.Mu.Unlock()
		if p.ProgressChannel == nil {
			c := make(chan float64, int(p.Max))
			p.ProgressChannel = &c
		}
		*p.ProgressChannel <- f

	}
}
