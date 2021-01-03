package core

import "github.com/gotk3/gotk3/gtk"

type AboutDialog struct {
	aboutDialog *gtk.AboutDialog
	Title       *string
	CopyRight   *string
	ProgramName *string
	Authors     []string
}

func (z *AboutDialog) Init() error {
	var err error
	z.aboutDialog, err = gtk.AboutDialogNew()
	if err != nil {
		return err
	}
	return z.Update()
}
func (z *AboutDialog) Update() error {
	if z.Title != nil {
		z.aboutDialog.SetTitle(*z.Title)
	}
	if z.Authors != nil {
		z.aboutDialog.SetAuthors(z.Authors)
	}
	if z.CopyRight != nil {
		z.aboutDialog.SetCopyright(*z.CopyRight)
	}
	if z.ProgramName != nil {
		z.aboutDialog.SetProgramName(*z.ProgramName)
	}
	z.aboutDialog.Run()
	return nil
}
func (z *AboutDialog)toWindow()*gtk.Window{
	return z.aboutDialog.ToWindow()
}
func (z *AboutDialog) BindOnClose(f func(show show)) error{
	_,ret := z.aboutDialog.Connect("destroy",func () {
		f (z)
	})
	return ret
}
func (z *AboutDialog) Show () error {
	z.aboutDialog.Run()
	return nil
}
func (z *AboutDialog) Hide () error {
	z.aboutDialog.Hide()
	return nil
}