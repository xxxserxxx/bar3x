package ui

import (
	"image"
	"image/color"
	"image/draw"
	"sync"
)

var _ ParentDrawable = (*Root)(nil)

type Root struct {
	im       *image.RGBA
	bounds   image.Rectangle
	bg       image.Image
	ctx      Context
	lock     sync.Mutex
	Inner    Drawable
	onNotify func()
}

func NewRoot(ctx Context, onNotify func()) *Root {
	return &Root{
		onNotify: onNotify,
		ctx:      ctx,
		bg:       image.NewUniform(color.Transparent),
	}
}

func (r *Root) Init() error {
	return nil
}
func (r *Root) Width() int {
	return r.Inner.Width()
}
func (r *Root) OnWidthChange(func(int))  {}
func (r *Root) OnHeightChange(func(int)) {}

func (r *Root) Height() int {
	return r.Inner.Height()
}

func (r *Root) Visible() bool {
	return true
}
func (r *Root) SetVisible(bool)            {}
func (r *Root) OnVisibleChange(func(bool)) {}

func (r *Root) Notify() {
	if r.Inner == nil {
		return
	}
	r.onNotify()
}

func (r *Root) SetContext(Context) {}
func (r *Root) Context() Context {
	return r.ctx
}
func (r *Root) Add(Drawable) {}
func (r *Root) Children() []Drawable {
	return []Drawable{r.Inner}
}
func (r *Root) ChildContext(int) Context {
	return r.ctx
}

func (r *Root) SendEvent(ev Event) bool {
	if r.Inner == nil {
		return false
	}
	return r.Inner.SendEvent(ev)
}
func (r *Root) Draw(x, y int, im draw.Image) {
}

func (r *Root) Paint() {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.Inner == nil {
		return
	}

	w, h := r.Inner.Width(), r.Inner.Height()
	if w == 0 || h == 0 {
		return
	}
	size := image.Rect(0, 0, w, h)
	r.ensureImage(size)
	draw.Draw(r.im, r.im.Rect, r.bg, image.Point{}, draw.Src)
	r.Inner.Draw(0, 0, r.im)
	r.bounds = size
}

func (r *Root) Image() *image.RGBA {
	if r.im == nil {
		w, h := r.Inner.Width(), r.Inner.Height()
		size := image.Rect(0, 0, w, h)
		r.ensureImage(size)
		//return nil
	}
	return r.im.SubImage(r.bounds).(*image.RGBA)
}

func (r *Root) ensureImage(size image.Rectangle) {
	if r.im == nil {
		goto Make
	}

	if r.im.Rect.Dx() < size.Dx() {
		goto Make
	}

	if r.im.Rect.Dy() < size.Dy() {
		goto Make
	}

	return

Make:
	r.im = image.NewRGBA(size)
}

func (b *Root) OnLeftClick() func(Event) bool {
	return b.Inner.OnLeftClick()
}
func (b *Root) OnRightClick() func(Event) bool {
	return b.Inner.OnRightClick()
}
func (b *Root) SetOnLeftClick(cb func(Event) bool) {
	b.Inner.SetOnLeftClick(cb)
}
func (b *Root) SetOnRightClick(cb func(Event) bool) {
	b.Inner.SetOnRightClick(cb)
}
func (b *Root) OnPointerMove() func(Event) bool {
	return b.Inner.OnPointerMove()
}
func (b *Root) SetOnPointerMove(cb func(Event) bool) {
	b.Inner.SetOnPointerMove(cb)
}
func (b *Root) OnPointerEnter() func(Event) bool {
	return b.Inner.OnPointerEnter()
}
func (b *Root) SetOnPointerEnter(cb func(Event) bool) {
	b.Inner.SetOnPointerEnter(cb)
}
func (b *Root) OnPointerLeave() func(Event) bool {
	return b.Inner.OnPointerLeave()
}
func (b *Root) SetOnPointerLeave(cb func(Event) bool) {
	b.Inner.SetOnPointerLeave(cb)
}
