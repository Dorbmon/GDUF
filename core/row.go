package core

import "github.com/gotk3/gotk3/gtk"

type Row struct {
	Children       []Element
	ElementPadding int
	box            *gtk.Box
}

func (z *Row) Init() error {
	var err error
	z.box, err = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 2)
	if err != nil {
		return err
	}
	for _, element := range z.Children {
		if err := element.Init(); err != nil {
			return err
		}
	}
	if err := z.Update();err != nil {
		return err
	}
	return nil
}
func (z *Row) Build() (gtk.IWidget, error) {
	if z.Children == nil {
		return z.box, nil
	}
	return z.box, nil
}
func (z *Row) Update() error {
	children := z.box.GetChildren()
	children.Foreach(func (Item interface{}) {
		z.box.Remove(Item.(gtk.IWidget))
	})
	for _, element := range z.Children {
		element.Update()
		widget, err := element.Build()
		if err != nil {
			return err
		}
		z.box.Add(widget)
	}
	return nil
}