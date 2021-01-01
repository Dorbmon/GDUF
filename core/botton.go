package core

import (
	"github.com/gotk3/gotk3/gtk"
)

type Button struct {
	Text string
	OnClick func()
}

func (z *Button) Init() error {
	return nil
}
func (z *Button) Build() (gtk.IWidget,error) {
	button, err := gtk.ButtonNewWithLabel("Hello World")		//创建一个按钮
	if err != nil {
		return nil,err
	}
	if z.OnClick != nil {
		button.Connect("clicked",z.OnClick)
	}
	return button,nil
}