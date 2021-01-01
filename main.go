package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"main.go/core"
)

func main() {
	gtk.Init(nil)
	app,_ := core.NewApp("rx.rx")
	mainWin := core.Window{
		Title: "TestWindow",
		DefaultSize: core.Vector2{
			X: 200,
			Y: 500,
		},
		Body: &core.Column{
			Children:       []core.Element{
				&core.Button{
					Text:    "Here",
					OnClick: func() {
						fmt.Println("Here")
					},
				},
			},
			ElementPadding: 4,
		},
	}
	app.AddWindow(&mainWin)
	app.Run()
}
