package core

import "github.com/gotk3/gotk3/gtk"

type ScrollView struct {
	scrolledWindow *gtk.ScrolledWindow
	Child Basic
}

func (z *ScrollView) Init() error {
	var err error
	z.scrolledWindow,err = gtk.ScrolledWindowNew(nil,nil)
	if err != nil {
		return err
	}
	return z.Update()
}
func (z *ScrollView) Update () error {
	c,_ := z.scrolledWindow.GetChild()
	var newChild gtk.IWidget = nil
	if z.Child != nil {
		var err error
		newChild,err = toIWidget(z.Child)
		if err != nil {
			return err
		}
	}
	if c != nil && c != newChild {
		z.scrolledWindow.Remove(c)
	}
	if c != newChild {
		z.scrolledWindow.Add(newChild)
	}
	return nil
}
func (z *ScrollView) Build() (gtk.IWidget,error) {
	return z.scrolledWindow,nil
}