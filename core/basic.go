package core

import (
	"errors"
	"github.com/gotk3/gotk3/gtk"
	"reflect"
)

type Element interface {
	Init() error
	Build() (gtk.IWidget, error)
	Update() error // Update data
}
type CustomWidget interface {
	Build() Element
	Init() error
	Update() error
}
type Basic interface {
	Init() error
	Update() error
}
type show interface {
	toWindow()*gtk.Window
	Init() error
	BindOnClose(func(show show))error
	Show()error
	Hide()error
}

var customWidgetType = reflect.TypeOf((*CustomWidget)(nil)).Elem()
var builtInWidgetType = reflect.TypeOf((*Element)(nil)).Elem()

type App struct {
	application *gtk.Application
	windows     []show
}

func init() {
	gtk.Init(nil)
}
func toIWidget(widget Basic) (gtk.IWidget, error) {
	// If it is a built-in widget
	t := reflect.TypeOf(widget)
	if t.Implements(customWidgetType) {
		return toIWidget(widget.(CustomWidget).Build())
	}
	if t.Implements(builtInWidgetType) {
		return widget.(Element).Build()
	}
	return nil, errors.New("it is not a widget offered")
}
