package core

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Image struct {
	image *gtk.Image
	pixbuf *gdk.Pixbuf
	ImageFile *string
	DefaultSize *Vector2
}
func (z *Image) Init() error {
	var err error
	z.image,err = gtk.ImageNew()
	if err != nil {
		return err
	}
	return z.Update()
}
func (z *Image) Build() (gtk.IWidget,error) {
	return z.image,nil
}
func (z *Image) Update() error {
	var err error
	z.pixbuf,err = gdk.PixbufNewFromFile(*z.ImageFile)
	if err != nil {
		return err
	}
	if z.DefaultSize != nil {
		z.pixbuf,err = z.pixbuf.ScaleSimple(z.DefaultSize.X,z.DefaultSize.Y,gdk.INTERP_BILINEAR)
		if err != nil {
			return err
		}
	}
	z.image.SetFromPixbuf(z.pixbuf)
	return nil
}