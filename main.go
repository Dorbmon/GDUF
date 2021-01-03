package main

import (
	"fmt"
	"main.go/core"
	"strconv"
)

func main() {
	num := 1
	numStr := "1"
	app, _ := core.NewApp("rx.rx")
	var mainWin core.Window
	mainWin = core.Window{
		Title: "TestWindow",
		DefaultSize: core.Vector2{
			X: 200,
			Y: 500,
		},
		Body: &core.Column{
			Children: []core.Element{
				&core.Button{
					Text: core.Str("点我来增加"),
					OnClick: func() {
						fmt.Println("xx")
						num++
						numStr = strconv.Itoa(num)
						mainWin.Update()
					},
				},
				&core.Text{Text: &numStr},
				&core.Entry{Text: core.Str("Initial Text")},
				&core.Image{ImageFile: core.Str("t.png"),DefaultSize: &core.Vector2{
					X: 200,
					Y: 200,
				}},
			},
			ElementPadding: 4,
		},
	}
	app.AddWindow(&mainWin)
	app.Run()
}
