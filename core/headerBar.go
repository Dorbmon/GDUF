package core

import "github.com/gotk3/gotk3/gtk"

type HeaderBar struct {
	headerBar *gtk.HeaderBar
	Title *string
	SubTitle *string
}

func (z *HeaderBar) Init() error {
	var err error
	z.headerBar,err = gtk.HeaderBarNew()
	if err != nil {
		return err
	}
	return z.Update()
}
func (z *HeaderBar) Update() error {
	if z.Title != nil {
		z.headerBar.SetTitle(*z.Title)
	}
	if z.SubTitle != nil {
		z.headerBar.SetSubtitle(*z.SubTitle)
	}
	return nil
}
func (z *HeaderBar) Build() (gtk.IWidget,error) {
	return z.headerBar,nil
}