package core

import "github.com/gotk3/gotk3/gtk"

type Entry struct {
	entry *gtk.Entry
	Text *string
}

func (z *Entry) Init () error {
	var err error
	z.entry,err = gtk.EntryNew()
	if err != nil {
		return err
	}
	return z.Update()
}
func (z *Entry) Build () (gtk.IWidget,error) {
	return z.entry,nil
}
func (z *Entry) Update () error {
	if z.Text != nil {
		z.entry.SetText(*z.Text)
	}
	return nil
}
func (z *Entry) GetText() (string,error){
	return z.entry.GetText()
}