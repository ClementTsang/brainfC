package bftoc

import (
	"fmt"
	"strings"
)

// rightArrowToken represents the ">" token
type rightArrowToken struct {
	times int
}

func (r *rightArrowToken) convertToC() (result string) {
	result = fmt.Sprintf("p+=%d;\n", r.times)
	return
}
func (r *rightArrowToken) getTokenStr() (result string) {
	result = strings.Repeat(">", r.times)
	return
}
func (r *rightArrowToken) getTimes() int {
	return r.times
}
func (r *rightArrowToken) addTime() {
	r.times++
}

// leftArrowToken represents the "<" token
type leftArrowToken struct {
	times int
}

func (l *leftArrowToken) convertToC() (result string) {
	result = fmt.Sprintf("p-=%d;\n", l.times)
	return
}
func (l *leftArrowToken) getTokenStr() (result string) {
	result = strings.Repeat("<", l.times)
	return
}
func (l *leftArrowToken) getTimes() int {
	return l.times
}
func (l *leftArrowToken) addTime() {
	l.times++
}

// plusToken represents the "+" token
type plusToken struct {
	times int
}

func (p *plusToken) convertToC() (result string) {
	result = fmt.Sprintf("*p+=%d;\n", p.times)
	return
}
func (p *plusToken) getTokenStr() (result string) {
	result = strings.Repeat("+", p.times)
	return
}
func (p *plusToken) getTimes() int {
	return p.times
}
func (p *plusToken) addTime() {
	p.times++
}

// minusToken represents the "-" token
type minusToken struct {
	times int
}

func (m *minusToken) convertToC() (result string) {
	result = fmt.Sprintf("*p-=%d;\n", m.times)
	return
}
func (m *minusToken) getTokenStr() (result string) {
	result = strings.Repeat("-", m.times)
	return
}
func (m *minusToken) getTimes() int {
	return m.times
}
func (m *minusToken) addTime() {
	m.times++
}

// dotToken represents the "." token
type dotToken struct{}

func (d *dotToken) convertToC() (result string) {
	result = "putchar(*p);\n"
	return
}
func (d *dotToken) getTokenStr() (result string) {
	result = "."
	return
}

func (d *dotToken) getTimes() int {
	return 1
}

func (d *dotToken) addTime() {}

// commaToken represents the "," token
type commaToken struct{}

func (c *commaToken) convertToC() (result string) {
	result = "*p = getchar();\n"
	return
}

func (c *commaToken) getTokenStr() (result string) {
	result = ","
	return
}

func (c *commaToken) getTimes() int {
	return 1
}

func (c *commaToken) addTime() {}

// leftBracketToken represents the "[" token
type leftBracketToken struct{}

func (l *leftBracketToken) convertToC() (result string) {
	result = "while (*p) {\n"
	return
}

func (l *leftBracketToken) getTokenStr() (result string) {
	result = "["
	return
}

func (l *leftBracketToken) getTimes() int {
	return 1
}

func (l *leftBracketToken) addTime() {}

// rightBracketToken represents the "]" token
type rightBracketToken struct{}

func (r *rightBracketToken) convertToC() (result string) {
	result = "}\n"
	return
}

func (r *rightBracketToken) getTokenStr() (result string) {
	result = "]"
	return
}

func (r *rightBracketToken) getTimes() int {
	return 1
}

func (r *rightBracketToken) addTime() {}
