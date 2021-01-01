package core

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func NewApp(AppId string) (*App, error) {
	app, err := gtk.ApplicationNew(AppId, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		return nil, err
	}
	return &App{
		application: app,
		windows:     make([]*Window, 0),
	}, nil
}
func (z *App) Run() error {
	z.application.Connect("activate", func() {
		// Prepare windows
		for _, window := range z.windows {
			if err := window.Init(); err != nil {
				return
			}
			window.window.SetApplication(z.application)
			z.application.AddWindow(window.window)
			window.window.ShowAll()
		}
	})
	z.application.Run(nil)
	return nil
}
func (z *App) AddWindow(window *Window) error {
	z.windows = append(z.windows, window)
	return nil
}
