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

When you want to build your custom widget,you should make sure these functions are in your *widget:

```
func Build() Element
func Init() error
func Update() error
```
