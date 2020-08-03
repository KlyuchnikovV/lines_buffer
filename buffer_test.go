package lines_buffer

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDeleteBackward(t *testing.T) {
	var originalText = []byte("Hello, \nworld!")
	b := NewBuffer(string(originalText))

	testCases := []struct {
		desc           string
		expected       string
		expectedRow    int
		expectedColumn int
		action         func()
	}{
		{
			desc:           "Delete '!' symbol",
			action:         func() { b.DeleteBackward() },
			expected:       "Hello, \nworld",
			expectedRow:    1,
			expectedColumn: 5,
		},
		{
			desc: "Delete '\\n' symbol",
			action: func() {
				assert.Equal(t, true, b.SetPosition(1, 0))
				b.DeleteBackward()
			},
			expected:       "Hello, world",
			expectedRow:    0,
			expectedColumn: 7,
		},
		{
			desc: "Try to delete -1 symbol",
			action: func() {
				assert.Equal(t, true, b.SetPosition(0, 0))
				b.DeleteBackward()
			},
			expected:       "Hello, world",
			expectedRow:    0,
			expectedColumn: 0,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.action()
			assert.Equal(t, string(tC.expected), b.String())
			assert.Equal(t, tC.expectedRow, b.row)
			assert.Equal(t, tC.expectedColumn, b.column)
		})
	}
}

func TestDeleteForward(t *testing.T) {
	var originalText = []byte("Hello, \nworld!")
	b := NewBuffer(string(originalText))

	testCases := []struct {
		desc           string
		expected       string
		expectedRow    int
		expectedColumn int
		action         func()
	}{
		{
			desc: "Delete '!' symbol",
			action: func() {
				assert.Equal(t, true, b.SetPosition(b.row, b.column-1))
				b.DeleteForward()
			},
			expected:       "Hello, \nworld",
			expectedRow:    1,
			expectedColumn: 5,
		},
		{
			desc:           "Try to delete (last + 1) symbol",
			action:         func() { b.DeleteForward() },
			expected:       "Hello, \nworld",
			expectedRow:    1,
			expectedColumn: 5,
		},
		{
			desc: "Delete '\\n' symbol",
			action: func() {
				assert.Equal(t, true, b.SetPosition(0, 7))
				b.DeleteForward()
			},
			expected:       "Hello, world",
			expectedRow:    0,
			expectedColumn: 7,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.action()
			assert.Equal(t, string(tC.expected), b.String())
			assert.Equal(t, tC.expectedRow, b.row)
			assert.Equal(t, tC.expectedColumn, b.column)
		})
	}
}

func TestInsertRune(t *testing.T) {
	var originalText = []byte("Hello, world")
	b := NewBuffer(string(originalText))

	testCases := []struct {
		desc           string
		expected       string
		expectedRow    int
		expectedColumn int
		action         func()
	}{
		{
			desc:           "Insert '!' symbol",
			expected:       "Hello, world!",
			expectedRow:    0,
			expectedColumn: 12,
			action: func() {
				b.InsertRune('!')
			},
		},
		{
			desc:           "Insert '\n' symbol",
			expected:       "Hello, \nworld!",
			expectedRow:    1,
			expectedColumn: 0,
			action: func() {
				b.SetPosition(0, 7)
				b.InsertRune('\n')
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.action()
			assert.Equal(t, string(tC.expected), b.String())
			assert.Equal(t, tC.expectedRow, b.row)
			assert.Equal(t, tC.expectedColumn, b.column)
		})
	}
}
