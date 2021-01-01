package core

import "github.com/gotk3/gotk3/gtk"

type Text struct {
	Text string
	label *gtk.Label
}

func (z *Text) Init() error {
	var err error
	z.label,err = gtk.LabelNew(z.Text)
	return err
}
func (z *Text) Build() (gtk.IWidget,error) {
	return z.label,nil
}