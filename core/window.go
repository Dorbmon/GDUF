package core

import "github.com/gotk3/gotk3/gtk"

type Vector2 struct {
	X int
	Y int
}
type Window struct {
	Title       string
	DefaultSize Vector2
	window      *gtk.Window
	Body        beyond
}

func (z *Window) Init() error {
	var err error
	z.window, err = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return err
	}
	if z.Body != nil {
		if err := z.Body.Init(); err != nil {
			return err
		}
		body, err := toIWidget(z.Body)
		if err != nil {
			return err
		}
		z.window.Add(body)
	}
	z.window.SetTitle(z.Title)
	z.window.SetDefaultSize(z.DefaultSize.X, z.DefaultSize.Y)
	return nil
}
func (z *Window) Update() error {
	return z.Body.Update()
}