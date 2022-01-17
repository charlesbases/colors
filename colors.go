package colors

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

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

// New .
func New(a ...attribute) *Color {
	var c Color = make([]attribute, 0, len(a))
	return c.Add(a...)
}

// Add .
func (c *Color) Add(a ...attribute) *Color {
	*c = append(*c, a...)
	return c
}

// Sprint .
func (c *Color) Sprint(a ...interface{}) string {
	return c.format() + fmt.Sprint(a...) + c.unformat()
}

// Sprintf .
func (c *Color) Sprintf(format string, a ...interface{}) string {
	return c.format() + fmt.Sprintf(format, a...) + c.unformat()
}

func (c *Color) format() string {
	return fmt.Sprintf("%s[%sm", escape, c.sequence())
}

func (c *Color) unformat() string {
	return fmt.Sprintf("%s[%dm", escape, reset)
}

func (c *Color) sequence() string {
	var list = make([]string, 0, len(*c))
	for _, item := range *c {
		list = append(list, item.string())
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	return strings.Join(list, ",")
}

func (a attribute) string() string {
	return strconv.Itoa(int(a))
}
