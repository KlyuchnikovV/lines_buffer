package lines_buffer

import (
	"log"
	"strings"
)

type Buffer struct {
	lines       []string
	row, column int
}

func NewBuffer(s string) *Buffer {
	lines := strings.Split(s, "\n")
	return &Buffer{
		lines:  lines,
		row:    len(lines) - 1,
		column: len(lines[len(lines)-1]),
	}
}

func (b *Buffer) Insert(s string) {
	if s == "\n" {
		b.NewLine()
		return
	}
	b.lines[b.row] = string(
		append(
			append(
				[]byte(b.lines[b.row])[:b.column], []byte(s)...,
			),
			[]byte(b.lines[b.row])[b.column:]...,
		),
	)
	b.column += len(s)
}

func (b *Buffer) NewLine() {
	var newLine string
	b.lines[b.row], newLine = b.lines[b.row][:b.column], b.lines[b.row][b.column:]
	var originalLines = b.Lines()[b.row+1:]
	b.lines = append(
		append(
			b.lines[:b.row+1],
			newLine,
		),
		originalLines...,
	)
	b.NextRune()
}

func (b *Buffer) DeleteBackward() {
	if b.column == 0 {
		if b.row == 0 {
			return
		}

		b.row--
		b.column = len(b.CurrentRow())
		b.mergeLineWithNext()
	} else {
		log.Printf("BUFFER: r: %d, c: %d", b.row, b.column)
		b.lines[b.row] = string(append([]byte(b.CurrentRow())[:b.column-1], []byte(b.CurrentRow())[b.column:]...))
		b.column--
	}
}

func (b *Buffer) DeleteForward() {
	if b.column == len(b.CurrentRow()) {
		if b.row == len(b.lines)-1 {
			return
		}
		b.mergeLineWithNext()
	} else {
		b.lines[b.row] = string(append([]byte(b.CurrentRow())[:b.column], []byte(b.CurrentRow())[b.column+1:]...))
	}
}

func (b *Buffer) mergeLineWithNext() {
	b.lines[b.row] = b.CurrentRow() + b.lines[b.row+1]
	b.lines = append(b.lines[:b.row+1], b.lines[b.row+2:]...)
}

func (b *Buffer) SetPosition(row, column int) bool {
	if row < 0 || row > len(b.lines)-1 {
		return false
	}
	if column < 0 || column > len(b.lines[row]) {
		return false
	}
	b.row, b.column = row, column
	return true
}

func (b *Buffer) String() string {
	return strings.Join(b.lines, "\n")
}

func (b *Buffer) CurrentRow() string {
	return b.lines[b.row]
}

func (b *Buffer) Lines() []string {
	var tempLines = make([]string, len(b.lines))
	copy(tempLines, b.lines)
	return tempLines
}

func (b *Buffer) RowNum() int {
	return b.row
}

func (b *Buffer) ColumnNum() int {
	return b.column
}

func (b *Buffer) PrevLine() {
	if b.row == 0 {
		return
	}
	b.row--
	if len(b.CurrentRow()) < b.column {
		b.column = len(b.CurrentRow())
	}
}

func (b *Buffer) NextLine() {
	if b.row == len(b.lines)-1 {
		return
	}
	b.row++
	if len(b.CurrentRow()) < b.column {
		b.column = len(b.CurrentRow())
	}
}

func (b *Buffer) NextRune() {
	if b.column == len(b.CurrentRow()) {
		b.NextLine()
		b.column = 0
	} else {
		b.column++
	}
}

func (b *Buffer) PrevRune() {
	if b.column == 0 {
		b.PrevLine()
		b.column = len(b.CurrentRow())
	} else {
		b.column--
	}
}

func (b *Buffer) MoveForward(n int) {
	for i := 0; i < n; i++ {
		b.NextRune()
	}
}

func (b *Buffer) MoveBackward(n int) {
	for i := 0; i < n; i++ {
		b.PrevRune()
	}
}

func (b *Buffer) MoveDown(n int) {
	for i := 0; i < n; i++ {
		b.NextLine()
	}
}

func (b *Buffer) MoveUp(n int) {
	for i := 0; i < n; i++ {
		b.PrevLine()
	}
}
