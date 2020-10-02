package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoteTitle(t *testing.T) {
	myNote := &note{content: "Hello"}
	assert.Equal(t, "Hello", myNote.title())

	myNote = &note{content: "note\nline2"}
	assert.Equal(t, "note", myNote.title())

	myNote = &note{content: ""}
	assert.Equal(t, "Untitled", myNote.title())

}

func TestNoteListAdd(t *testing.T) {
	notes := &noteList{}

	notes.add()
	assert.Equal(t, 1, len(notes.list))
}

func TestNoteListRemove(t *testing.T) {
	first := &note{content: "remove me"}
	second := &note{content: "remove me"}
	notes := &noteList{list: []*note{first, second}}

	assert.Equal(t, 2, len(notes.list))
	notes.remove(first)
	assert.Equal(t, 1, len(notes.list))
	notes.remove(second)
	assert.Equal(t, 0, len(notes.list))

}
