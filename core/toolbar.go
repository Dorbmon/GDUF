package core

import "github.com/gotk3/gotk3/gtk"

type Toolbar struct {
	toolBar *gtk.Toolbar
}

func (z *Toolbar) Init() error {
	var err error
	z.toolBar,err = gtk.ToolbarNew()
	if err != nil {
		return err
	}
	return z.Update()
}
func (z *Toolbar) Update() error {
	return nil
}