package interact

import (
	"image/color"
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/s-daehling/fyne-charts/pkg/style"
)

type Legend struct {
	widget.BaseWidget
	les         []*LegendEntry
	location    style.LegendLocation
	style       style.LabelStyle
	interactive bool
}

func NewLegend() (l *Legend) {
	l = &Legend{
		les: make([]*LegendEntry, 0),
	}
	l.SetStyle(style.LegendLocationRight, style.DefaultLegendLabelStyle(), true)
	l.ExtendBaseWidget(l)
	return
}

func (l *Legend) CreateRenderer() (r fyne.WidgetRenderer) {
	r = newLegendRenderer(l)
	return
}

func (l *Legend) AddEntry(le *LegendEntry) {
	le.setInteractiveness(l.interactive)
	le.setStyle(l.style)
	if le.super == "" {
		l.les = append(l.les, le)
	} else {
		superFound := false
		insertIndex := -1
		for i := range l.les {
			if l.les[i].name == le.super {
				superFound = true
				continue
			}
			if superFound && l.les[i].super != le.super {
				insertIndex = i
				break
			}
		}
		if insertIndex == -1 {
			insertIndex = len(l.les)
		}
		l.les = slices.Insert(l.les, insertIndex, le)
	}
	l.updateSubDepiction()
}

func (l *Legend) RemoveEntry(name string, super string) {
	for i := range l.les {
		if l.les[i].name == name && l.les[i].super == super {
			l.les = slices.Delete(l.les, i, i+1)
			break
		}
	}
	l.updateSubDepiction()
}

func (l *Legend) Location() (loc style.LegendLocation) {
	loc = l.location
	return
}

func (l *Legend) SetStyle(loc style.LegendLocation, s style.LabelStyle, interactive bool) {
	l.location = loc
	l.style = s
	l.interactive = interactive
	for i := range l.les {
		l.les[i].setStyle(s)
		l.les[i].setInteractiveness(interactive)
	}
	l.updateSubDepiction()
	l.Refresh()
}

func (l *Legend) updateSubDepiction() {
	onlyOneSuper := true
	super := ""
	for i := range l.les {
		if l.les[i].super == "" {
			if super != "" {
				onlyOneSuper = false
				break
			}
			super = l.les[i].name
		} else {
			if l.les[i].super != super {
				onlyOneSuper = false
				break
			}
		}
	}
	for i := range l.les {
		if l.location == style.LegendLocationBottom || l.location == style.LegendLocationTop {
			l.les[i].setSubDepiction(false, !onlyOneSuper)
		} else {
			l.les[i].setSubDepiction(true, false)
		}
	}
}

type legendRenderer struct {
	l   *Legend
	col int
	row int
}

func newLegendRenderer(l *Legend) (lr *legendRenderer) {
	lr = &legendRenderer{
		l:   l,
		col: 1,
		row: 1,
	}
	return
}

func (lr *legendRenderer) Layout(size fyne.Size) {
	sEntry := lr.entrySize()
	lr.col, lr.row = lr.NumColRow(size)
	totWidth := float32(lr.col) * sEntry.Width
	xOff := (size.Width - totWidth) / 2.0
	y := float32(0.0)
	x := xOff
	colCount := 0
	for i := range lr.l.les {
		if (lr.l.location == style.LegendLocationBottom || lr.l.location == style.LegendLocationTop) && !lr.l.les[i].showBox {
			continue
		}
		xAlign := float32(0)
		switch lr.l.style.Alignment {
		case fyne.TextAlignCenter:
			xAlign = (sEntry.Width - lr.l.les[i].MinSize().Width) / 2
		case fyne.TextAlignTrailing:
			xAlign = sEntry.Width - lr.l.les[i].MinSize().Width
		}
		lr.l.les[i].Resize(lr.l.les[i].MinSize())
		lr.l.les[i].Move(fyne.NewPos(x+xAlign, y))
		colCount++
		if colCount < lr.col {
			x += sEntry.Width
		} else {
			x = xOff
			y += sEntry.Height
			colCount = 0
		}
	}
}

func (lr *legendRenderer) MinSize() (size fyne.Size) {
	size.Width = lr.entrySize().Width
	size.Height = lr.entrySize().Height * float32(lr.row)
	return
}

func (lr *legendRenderer) Refresh() {
	for i := range lr.l.les {
		lr.l.les[i].RefreshTheme()
		lr.l.les[i].box.Refresh()
		lr.l.les[i].label.Refresh()
	}
}

func (lr *legendRenderer) Objects() (canObj []fyne.CanvasObject) {
	for i := range lr.l.les {
		if (lr.l.location == style.LegendLocationBottom || lr.l.location == style.LegendLocationTop) && !lr.l.les[i].showBox {
			continue
		}
		canObj = append(canObj, lr.l.les[i])
	}
	return
}

func (lr *legendRenderer) Destroy() {}

func (lr *legendRenderer) NumColRow(size fyne.Size) (nc int, nr int) {
	nc = 1
	nr = 0
	nles := len(lr.l.les)
	if nles == 0 {
		return
	}
	if lr.l.location == style.LegendLocationBottom || lr.l.location == style.LegendLocationTop {
		nles = 0
		for i := range lr.l.les {
			if !lr.l.les[i].showBox {
				continue
			}
			nles++
		}
		wEntry := lr.entrySize().Width
		nc = int(size.Width / wEntry)
		if nc == 0 {
			nc = 1
		}
	}
	nr = nles / nc
	if nles%nc > 0 {
		nr++
	}
	if nr == 1 {
		nc = nles
	}
	return
}

func (lr *legendRenderer) entrySize() (size fyne.Size) {
	size.Width = 0
	size.Height = 0
	for i := range lr.l.les {
		if lr.l.les[i].MinSize().Width > size.Width {
			size.Width = lr.l.les[i].MinSize().Width
		}
		if lr.l.les[i].MinSize().Height > size.Height {
			size.Height = lr.l.les[i].MinSize().Height
		}
	}
	return
}

type LegendEntry struct {
	widget.BaseWidget
	name         string
	super        string
	showBox      bool
	subIndent    bool
	subShowSuper bool
	box          *legendBox
	label        *canvas.Text
	style        style.LabelStyle
}

func NewLegendEntry(name string, super string, showBox bool, col color.Color, tapFct func()) (le *LegendEntry) {
	le = &LegendEntry{
		name:    name,
		super:   super,
		showBox: showBox,
		box:     NewLegendBox(col, tapFct),
		label:   canvas.NewText(name, theme.Color(theme.ColorNameForeground)),
	}
	le.label.Resize(le.label.MinSize())
	le.box.Resize(fyne.NewSize(le.label.MinSize().Height*0.8, le.label.MinSize().Height*0.8))
	le.ExtendBaseWidget(le)
	return
}

func (le *LegendEntry) CreateRenderer() (r fyne.WidgetRenderer) {
	r = newLegendEntryRenderer(le)
	return
}

func (le *LegendEntry) RefreshTheme() {
	le.label.Color = theme.Color(le.style.ColorName)
	le.label.TextSize = theme.Size(le.style.SizeName)
}

func (le *LegendEntry) SetSuper(super string) {
	le.super = super
}

func (le *LegendEntry) HideBox() {
	le.showBox = false
}

func (le *LegendEntry) setSubDepiction(indent bool, showSuper bool) {
	le.subIndent = indent
	le.subShowSuper = showSuper
	if le.super != "" {
		le.label.Text = le.name
		if showSuper {
			le.label.Text += " (" + le.super + ")"
		}
	}
}

func (le *LegendEntry) SetColor(col color.Color) {
	le.box.SetColor(col)
}

func (le *LegendEntry) Show() {
	le.box.ToRect()
}

func (le *LegendEntry) Hide() {
	le.box.ToCircle()
}

func (le *LegendEntry) setStyle(ls style.LabelStyle) {
	le.style = ls
	le.label.Color = theme.Color(ls.ColorName)
	le.label.TextSize = theme.Size(ls.SizeName)
	le.label.TextStyle = ls.TextStyle
}

func (le *LegendEntry) setInteractiveness(interactive bool) {
	le.box.setInteractiveness(interactive)
}

type legendEntryRenderer struct {
	le *LegendEntry
}

func newLegendEntryRenderer(le *LegendEntry) (ler *legendEntryRenderer) {
	ler = &legendEntryRenderer{
		le: le,
	}
	return
}

func (ler *legendEntryRenderer) Layout(size fyne.Size) {
	x := float32(0.0)
	switch ler.le.style.Alignment {
	case fyne.TextAlignLeading:
		if ler.le.super != "" && ler.le.subIndent {
			x += ler.le.box.Size().Width
		}
		if ler.le.showBox {
			ler.le.box.Move(fyne.NewPos(x, (ler.le.label.Size().Height-ler.le.box.Size().Height)/2))
			x += ler.le.box.Size().Width + 5
		}
		ler.le.label.Move(fyne.NewPos(x, 0))
	case fyne.TextAlignCenter:
		if ler.le.showBox {
			ler.le.box.Move(fyne.NewPos(x, (ler.le.label.Size().Height-ler.le.box.Size().Height)/2))
			x += ler.le.box.Size().Width + 5
		}
		ler.le.label.Move(fyne.NewPos(x, 0))
	case fyne.TextAlignTrailing:
		x = ler.MinSize().Width - 5
		if ler.le.super != "" && ler.le.subIndent {
			x -= ler.le.box.Size().Width
		}
		if ler.le.showBox {
			x -= ler.le.box.Size().Width
			ler.le.box.Move(fyne.NewPos(x, (ler.le.label.Size().Height-ler.le.box.Size().Height)/2))
			x -= 5
		}
		x -= ler.le.label.MinSize().Width
		ler.le.label.Move(fyne.NewPos(x, 0))
	}
}

func (ler *legendEntryRenderer) MinSize() (size fyne.Size) {
	size.Width = 5 + ler.le.label.MinSize().Width
	if ler.le.super != "" && ler.le.subIndent {
		size.Width += ler.le.box.Size().Width
	}
	if ler.le.showBox {
		size.Width += ler.le.box.Size().Width + 5
	}
	size.Height = ler.le.label.MinSize().Height
	return
}

func (ler *legendEntryRenderer) Refresh() {
	ler.le.RefreshTheme()
	ler.le.box.Refresh()
	ler.le.label.Refresh()
}

func (ler *legendEntryRenderer) Objects() (canObj []fyne.CanvasObject) {
	if ler.le.showBox {
		canObj = append(canObj, ler.le.box)
	}
	canObj = append(canObj, ler.le.label)
	return
}

func (ler *legendEntryRenderer) Destroy() {}

type legendBox struct {
	widget.BaseWidget
	rectColor   color.Color
	rect        *canvas.Rectangle
	circle      *canvas.Circle
	interactive bool
	tapFct      func()
}

func NewLegendBox(col color.Color, tapFct func()) *legendBox {
	box := &legendBox{
		rect:        canvas.NewRectangle(col),
		circle:      canvas.NewCircle(col),
		rectColor:   col,
		interactive: true,
		tapFct:      tapFct,
	}
	box.ExtendBaseWidget(box)
	return box
}

func (box *legendBox) CreateRenderer() fyne.WidgetRenderer {
	// c := container.NewStack(box.rect, box.grad)
	c := container.NewStack(box.rect, box.circle)
	return widget.NewSimpleRenderer(c)
}

func (box *legendBox) Tapped(_ *fyne.PointEvent) {
	if box.interactive {
		box.tapFct()
	}
}

func (box *legendBox) MouseIn(me *desktop.MouseEvent) {
	if !box.interactive {
		return
	}
	r, g, b, a := box.rectColor.RGBA()
	rb, gb, bb, _ := theme.Color(theme.ColorNameBackground).RGBA()
	// box.rect.FillColor = color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: 0xaaaa}
	box.rect.FillColor = color.RGBA64{R: uint16(float32(r+rb) * 0.5), G: uint16(float32(g+gb) * 0.5), B: uint16(float32(b+bb) * 0.5), A: uint16(a)}
	box.rect.Refresh()
	box.circle.FillColor = color.RGBA64{R: uint16(float32(r+rb) * 0.5), G: uint16(float32(g+gb) * 0.5), B: uint16(float32(b+bb) * 0.5), A: uint16(a)}
	box.circle.Refresh()
}

func (box *legendBox) MouseMoved(me *desktop.MouseEvent) {}

func (box *legendBox) MouseOut() {
	if !box.interactive {
		return
	}
	box.rect.FillColor = box.rectColor
	box.rect.Refresh()
	box.circle.FillColor = box.rectColor
	box.circle.Refresh()
}

func (box *legendBox) SetColor(col color.Color) {
	box.rectColor = col
	box.rect.FillColor = col
	box.circle.FillColor = col
}

func (box *legendBox) ToCircle() {
	box.rect.Hide()
}

func (box *legendBox) ToRect() {
	box.rect.Show()
}

func (box *legendBox) setInteractiveness(interactive bool) {
	box.interactive = interactive
}
