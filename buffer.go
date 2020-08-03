package lines_buffer

import "strings"

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

func (b *Buffer) InsertRune(r rune) {
	if r == '\n' {
		b.NewLine()
		return
	}
	b.lines[b.row] = string(
		append(
			append(
				[]byte(b.lines[b.row])[:b.column], byte(r),
			),
			[]byte(b.lines[b.row])[b.column:]...,
		),
	)
}

func (b *Buffer) NewLine() {
	var newLine string
	b.lines[b.row], newLine = b.lines[b.row][:b.column], b.lines[b.row][b.column:]
	b.lines = append(b.lines, newLine)
	b.column = 0
	b.row++
}

func (b *Buffer) DeleteBackward() {
	if b.column == 0 {
		if b.row == 0 {
			return
		}

		b.row--
		b.column = len(b.currentRow())
		b.mergeLineWithNext()
	} else {
		b.lines[b.row] = string(append([]byte(b.currentRow())[:b.column-1], []byte(b.currentRow())[b.column:]...))
		b.column--
	}
}

func (b *Buffer) DeleteForward() {
	if b.column == len(b.currentRow()) {
		if b.row == len(b.lines)-1 {
			return
		}
		b.mergeLineWithNext()
	} else {
		b.lines[b.row] = string(append([]byte(b.currentRow())[:b.column], []byte(b.currentRow())[b.column+1:]...))
	}
}

func (b *Buffer) mergeLineWithNext() {
	b.lines[b.row] = b.currentRow() + b.lines[b.row+1]
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

func (b *Buffer) currentRow() string {
	return b.lines[b.row]
}
