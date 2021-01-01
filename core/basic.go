package core

import "github.com/gotk3/gotk3/gtk"

type Element interface {
	Init()error
	Build()(gtk.IWidget,error)
}
type App struct {
	application *gtk.Application
	windows []*Window
}