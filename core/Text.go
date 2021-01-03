package core

import (
	"github.com/gotk3/gotk3/gtk"
)

type Text struct {
	Text  *string
	label *gtk.Label
}

func (z *Text) Init() error {
	var err error
	z.label, err = gtk.LabelNew(*(z.Text))
	if err := z.Update(); err != nil {
		return err
	}
	return err
}
func (z *Text) Build() (gtk.IWidget, error) {
	return z.label, nil
}
func (z *Text) Update() error {
	z.label.SetText(*(z.Text))
	return nil
}
