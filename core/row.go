package core

import "github.com/gotk3/gotk3/gtk"

type Row struct {
	Children []Element
	ElementPadding int
	box *gtk.Box
}

func (z *Row) Init() error {
	var err error
	z.box, err = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 2)
	if err != nil {
		return err
	}
	for _,element := range z.Children {
		if err := element.Init();err != nil {
			return err
		}
	}
	return nil
}
func (z *Row) Build() (gtk.IWidget,error) {
	if z.Children == nil {
		return z.box,nil
	}
	for _,element := range z.Children {
		widget,err := element.Build()
		if err != nil {
			return nil,err
		}
		z.box.Add(widget)
	}
	return z.box,nil
}