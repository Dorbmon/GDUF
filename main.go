package main

import (
	"fmt"
	"main.go/core"
)

func main() {
	app, _ := core.NewApp("rx.rx")
	mainWin := core.Window{
		Title: "TestWindow",
		DefaultSize: core.Vector2{
			X: 200,
			Y: 500,
		},
		Body: &core.Column{
			Children: []core.Element{
				&core.Button{
					Text: "Here",
					OnClick: func() {
						fmt.Println("Here")
					},
				},
				&core.Text{Text: "HiHi"},
			},
			ElementPadding: 4,
		},
	}
	app.AddWindow(&mainWin)
	app.Run()
}
