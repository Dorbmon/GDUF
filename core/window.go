package core

import "github.com/gotk3/gotk3/gtk"

type Vector2 struct {
	X int
	Y int
}
type Window struct {
	Title       *string
	DefaultSize *Vector2
	window      *gtk.Window
	Body        Basic
	TitleBar	*HeaderBar
}

func (z *Window) BindApp(application *gtk.Application) error {
	z.window.SetApplication(application)
	return nil
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
	if z.TitleBar != nil {
		err := z.TitleBar.Init()
		if err != nil {
			return err
		}
	}
	return z.Update()
}
func (z *Window) Update() error {
	if z.Title != nil {
		z.window.SetTitle(*z.Title)
	}
	if z.DefaultSize != nil {
		z.window.SetDefaultSize(z.DefaultSize.X, z.DefaultSize.Y)
	}
	if z.TitleBar != nil {
		bar,err := toIWidget(z.TitleBar)
		if err != nil {
			return err
		}
		z.window.SetTitlebar(bar)
	}
	return z.Body.Update()
}
func (z *Window) toWindow() *gtk.Window {
	return z.window
}
func (z *Window) BindOnClose(f func(show show)) error{
	_,ret := z.window.Connect("destroy", func() {
		f (z)
	})
	return ret
}
func (z *Window) Show () error {
	z.window.ShowAll()
	return nil
}
func (z *Window) Hide () error {
	z.window.Hide()
	return nil
}