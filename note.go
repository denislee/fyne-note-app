package main

import (
	"strings"
)

type note struct {
	content string
}

func (n *note) title() string {
	if n.content == "" {
		return "Untitled"
	}
	return strings.SplitN(n.content, "\n", 2)[0]
}

type noteList struct {
	list []*note
}

func (l *noteList) add() *note {
	n := &note{}
	l.list = append([]*note{n}, l.list...)
	return n
}

func (l *noteList) remove(n *note) {
	if len(l.list) == 0 {
		return
	}

	for i, note := range l.list {
		if note != n {
			continue
		}
		if i == len(l.list)-1 {
			l.list = l.list[:i]
		} else {
			l.list = append(l.list[:i], l.list[i+1:]...)
		}
	}
}
