# GDUF
Golang Decarative UI Framework based on GTK3.

### You can write code in this way:
```go
    core.Window{
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
```