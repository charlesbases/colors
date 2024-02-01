package colors

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"unsafe"
)

// Color .
type Color []attribute

type attribute int

const escape = "\x1b"

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
	F_Black attribute = iota + 30
	F_Red
	F_Green
	F_Yellow
	F_Blue
	F_Purple
	F_Cyan
	F_White
)

// Foreground Hi-Intensity text colors
const (
	F_HiBlack attribute = iota + 90
	F_HiRed
	F_HiGreen
	F_HiYellow
	F_HiBlue
	F_HiPurple
	F_HiCyan
	F_HiWhite
)

// Background text colors
const (
	B_Black attribute = iota + 40
	B_Red
	B_Green
	B_Yellow
	B_Blue
	B_Purple
	B_Cyan
	B_White
)

// Background Hi-Intensity text colors
const (
	B_HiBlack attribute = iota + 100
	B_HiRed
	B_HiGreen
	B_HiYellow
	B_HiBlue
	B_HiPurple
	B_HiCyan
	B_HiWhite
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
func (c Color) sprint(s string) []byte {
	buff := newBuffer()
	defer bufferPool.Put(buff)

	// head
	buff.WriteString(escape)
	buff.WriteString("[")
	buff.WriteString(c.sequence())
	buff.WriteString("m")

	// mess
	buff.WriteString(s)

	// tail
	buff.WriteString(escape)
	buff.WriteString("[")
	buff.WriteString(reset.toString())
	buff.WriteString("m")

	return buff.Bytes()
}

// Sprint .
func (c Color) Sprint(a ...interface{}) string {
	mess := c.sprint(fmt.Sprint(a...))
	return *(*string)(unsafe.Pointer(&mess))
}

// Sprintf .
func (c Color) Sprintf(format string, a ...interface{}) string {
	mess := c.sprint(fmt.Sprintf(format, a...))
	return *(*string)(unsafe.Pointer(&mess))
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
