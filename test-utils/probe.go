package testutils

import gologicgates "github.com/jhtrauntvein/go-logic-gates"

type Probe struct {
   Value bool
   Count int
}

func (p *Probe) OnLineChanged(line *gologicgates.Line, value bool) {
   p.Value = value
   p.Count++
}
