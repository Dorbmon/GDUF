package core

import "github.com/gotk3/gotk3/gtk"

type ProcessBar struct {
	processBar *gtk.ProgressBar
	Fraction *float64
}

func (z *ProcessBar) Init () error {
	var err error
	z.processBar,err = gtk.ProgressBarNew()
	if err != nil {
		return err
	}
	return z.Update()
}
func (z *ProcessBar) Update () error {
	if z.Fraction != nil {
		z.processBar.SetFraction(*z.Fraction)
	}
	return nil
}
func (z *ProcessBar) Build() (gtk.IWidget,error) {
	return z.processBar,nil
}
