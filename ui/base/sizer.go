package base

import (
	"image/draw"

	"github.com/shimmerglass/bar3x/ui"
)

const (
	HAlighCenter = "center"
	HAlignLeft   = "left"
	HAlignRight  = "right"
)

const (
	VAlighMiddle = "middle"
	VAlignTop    = "top"
	VAlignBottom = "bottom"
)

var _ ui.ParentDrawable = (*Sizer)(nil)

type Sizer struct {
	Base

	setWidth  int
	setHeight int

	hAlign string
	vAlign string

	paddingTop    int
	paddingRight  int
	paddingBottom int
	paddingLeft   int

	inner ui.Drawable
}

func NewSizer(p ui.ParentDrawable) *Sizer {
	return &Sizer{
		Base:      NewBase(p),
		hAlign:    HAlighCenter,
		vAlign:    VAlighMiddle,
		setWidth:  -1,
		setHeight: -1,
	}
}

func (s *Sizer) Add(d ui.Drawable) {
	s.inner = d
	d.OnWidthChange(func(int) {
		s.updateSize()
	})
	d.OnHeightChange(func(int) {
		s.updateSize()
	})
	d.OnVisibleChange(func(bool) {
		s.updateSize()
	})
}

func (c *Sizer) ChildContext(i int) ui.Context {
	return c.ctx
}
func (r *Sizer) Children() []ui.Drawable {
	return []ui.Drawable{r.inner}
}
func (c *Sizer) SetContext(ctx ui.Context) {
	c.ctx = ctx
	if c.inner != nil {
		c.inner.SetContext(ctx)
	}
}

func (s *Sizer) SetWidth(v int) {
	s.setWidth = v
	s.updateSize()
}
func (s *Sizer) SetHeight(v int) {
	s.setHeight = v
	s.updateSize()
}

func (s *Sizer) HAlign() string {
	return s.hAlign
}
func (s *Sizer) SetHAlign(v string) {
	s.hAlign = v
}

func (s *Sizer) VAlign() string {
	return s.vAlign
}
func (s *Sizer) SetVAlign(v string) {
	s.vAlign = v
}

func (s *Sizer) PaddingTop() int {
	return s.paddingTop
}
func (s *Sizer) SetPaddingTop(v int) {
	s.paddingTop = v
	s.updateSize()
}
func (s *Sizer) PaddingRight() int {
	return s.paddingRight
}
func (s *Sizer) SetPaddingRight(v int) {
	s.paddingRight = v
	s.updateSize()
}
func (s *Sizer) PaddingBottom() int {
	return s.paddingBottom
}
func (s *Sizer) SetPaddingBottom(v int) {
	s.paddingBottom = v
	s.updateSize()
}
func (s *Sizer) PaddingLeft() int {
	return s.paddingLeft
}
func (s *Sizer) SetPaddingLeft(v int) {
	s.paddingLeft = v
	s.updateSize()
}

func (s Sizer) updateSize() {
	w, h := s.setWidth, s.setHeight
	ew, eh := 0, 0
	if s.inner != nil && s.inner.Visible() {
		ew, eh = s.inner.Width(), s.inner.Height()
	}

	if w == -1 {
		w = ew
		w += s.paddingLeft + s.paddingRight
	}
	if h == -1 {
		h = eh
		h += s.paddingTop + s.paddingBottom
	}

	s.width.set(w)
	s.height.set(h)
}

func (s Sizer) Draw(tx, ty int, im draw.Image) {
	if !s.inner.Visible() {
		return
	}

	w, h := s.width.v, s.height.v
	ew, eh := s.inner.Width(), s.inner.Height()

	var x, y int
	switch s.hAlign {
	case HAlighCenter:
		d := w - ew - s.paddingLeft - s.paddingRight
		x = s.paddingLeft + d/2
	case HAlignLeft:
		x = s.paddingLeft
	case HAlignRight:
		x = w - ew - s.paddingRight
	}
	switch s.vAlign {
	case VAlighMiddle:
		d := h - eh - s.paddingTop - s.paddingBottom
		y = s.paddingTop + d/2
	case VAlignTop:
		y = s.paddingTop
	case VAlignBottom:
		y = h - eh - s.paddingBottom
	}

	s.inner.Draw(tx+x, ty+y, im)
}