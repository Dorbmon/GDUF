package core

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func Str(str string) *string {
	return &str
}
func F64(f64 float64)*float64 {
	return &f64
}
func NewApp(AppId string) (*App, error) {
	app, err := gtk.ApplicationNew(AppId, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		return nil, err
	}
	return &App{
		application: app,
		windows:     make([]show, 0),
	}, nil
}
func (z *App) Run() error {
	z.application.Connect("activate", func() {
		// Prepare windows
		for _, window := range z.windows {
			if err := window.Init(); err != nil {
				return
			}
			window.toWindow().SetApplication(z.application)
			z.application.AddWindow(window.toWindow())
			window.BindOnClose(func(show show) {
				z.application.RemoveWindow(window.toWindow())
				return
			})
			//window.toWindow().ShowAll()
			window.Show()
		}
	})
	z.application.Run(nil)
	return nil
}
func (z *App) AddWindow(window show) error {
	z.windows = append(z.windows, window)
	return nil
}
func (z *App) ShowDialog(dialog gtk.Dialog) {
	dialog.ShowAll()
}
