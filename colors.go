package colors

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// Color .
type Color []attribute

type attribute int

const (
	escape = "\x1b"
	ending = escape + "[0m"
)

// Base attributes
const (
	reset attribute = iota
	bold
	faint
	italic
	underline
	blinkslow
	blinkrapid
	reversevideo
	concealed
	crossedout
)

// Foreground text colors
const (
	ForegroundBlack attribute = iota + 30
	ForegroundRed
	ForegroundGreen
	ForegroundYellow
	ForegroundBlue
	ForegroundPurple
	ForegroundCyan
	ForegroundWhite
)

// Foreground Hi-Intensity text colors
const (
	ForegroundHiBlack attribute = iota + 90
	ForegroundHiRed
	ForegroundHiGreen
	ForegroundHiYellow
	ForegroundHiBlue
	ForegroundHiPurple
	ForegroundHiCyan
	ForegroundHiWhite
)

// Background text colors
const (
	BackgroundBlack attribute = iota + 40
	BackgroundRed
	BackgroundGreen
	BackgroundYellow
	BackgroundBlue
	BackgroundPurple
	BackgroundCyan
	BackgroundWhite
)

// Background Hi-Intensity text colors
const (
	BackgroundHiBlack attribute = iota + 100
	BackgroundHiRed
	BackgroundHiGreen
	BackgroundHiYellow
	BackgroundHiBlue
	BackgroundHiPurple
	BackgroundHiCyan
	BackgroundHiWhite
)

var bufferPool sync.Pool

// newBuffer .
func newBuffer() *bytes.Buffer {
	if v := bufferPool.Get(); v != nil {
		b := v.(*bytes.Buffer)
		b.Reset()
		return b
	}
	return bytes.NewBuffer(nil)
}

// New .
func New(a ...attribute) Color {
	return a
}

// sprint .
func (c Color) sprint(s string) string {
	buff := newBuffer()
	defer bufferPool.Put(buff)

	// head
	buff.WriteString(escape)
	buff.WriteString("[")
	buff.WriteString(c.sequence())
	buff.WriteString("m")

	buff.WriteString(s)

	buff.WriteString(ending)
	return buff.String()
}

// Sprint .
func (c Color) Sprint(a ...interface{}) string {
	return c.sprint(fmt.Sprint(a...))
}

// Sprintf .
func (c Color) Sprintf(format string, a ...interface{}) string {
	return c.sprint(fmt.Sprintf(format, a...))
}

// sequence .
func (c Color) sequence() string {
	switch len(c) {
	case 0:
		return ""
	case 1:
		return c[0].toString()
	default:
		var res = make([]string, 0, len(c))
		for _, att := range c {
			res = append(res, att.toString())
		}
		return strings.Join(res, ";")
	}
}

// toString .
func (a attribute) toString() string {
	return strconv.Itoa(int(a))
}
