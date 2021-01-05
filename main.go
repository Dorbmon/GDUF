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
	process := float64(0)
	var mainWin core.Window
	entry := core.Entry{Text: core.Str("Initial Text")}
	mainWin = core.Window{
		Title: core.Str("TestWindow"),
		TitleBar: &core.HeaderBar{
			Title:    core.Str("TestWindow"),
			SubTitle: core.Str("Subtitle"),
		},
		Body: &core.Column{
			Children: []core.Basic{
				&core.Button{
					Text: core.Str("点我来增加"),
					OnClick: func() {
						fmt.Println(entry.GetText())
						num++
						numStr = strconv.Itoa(num)
						process += 0.01
						mainWin.Update()
					},
				},
				&core.Text{Text: &numStr},
				&entry,
				&core.Image{ImageFile: core.Str("t.png"), DefaultSize: &core.Vector2{
					X: 200,
					Y: 200,
				}},
				&core.ProcessBar{Fraction: &process},
			},
			ElementPadding: 4,
		},
	}
	app.AddWindow(&mainWin)
	app.Run()
}
