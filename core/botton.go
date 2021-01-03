package core

import (
	"github.com/gotk3/gotk3/gtk"
)

type Button struct {
	Text    *string
	OnClick func()
	button  *gtk.Button
}

func (z *Button) Init() error {
	var err error
	z.button, err = gtk.ButtonNewWithLabel(*z.Text) //创建一个按钮
	if err != nil {
		return err
	}
	if z.OnClick != nil {
		z.button.Connect("clicked", func() {
			if z.OnClick != nil {
				z.OnClick()
			}
		})
	}
	return nil
}
func (z *Button) Build() (gtk.IWidget, error) {
	return z.button, nil
}
func (z *Button) Update() error {
	z.button.SetLabel(*z.Text)
	return nil
}
