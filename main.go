package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type ui struct {
	current *note
	notes   *noteList

	content *widget.Entry
	list    *widget.Box
}

func (u *ui) addNote() {
	newNote := u.notes.add()
	u.setNote(newNote)
}

func (u *ui) setNote(n *note) {
	u.current = n
	u.content.SetText(n.content)
	u.refreshList()
}

func (u *ui) refreshList() {
	u.list.Children = nil
	for _, n := range u.notes.list {
		thisNote := n
		button := widget.NewButton(n.title(), func() {
			u.setNote(thisNote)
		})
		if n == u.current {
			button.Style = widget.PrimaryButton
		}

		u.list.Append(button)
	}

}

func (u *ui) loadUI() fyne.CanvasObject {
	u.content = widget.NewMultiLineEntry()

	u.list = widget.NewVBox()
	u.refreshList()

	if len(u.notes.list) > 0 {
		u.setNote(u.notes.list[0])
	}

	u.content.OnChanged = func(content string) {
		if u.current == nil {
			return
		}
		u.current.content = content
		u.refreshList()
	}

	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			u.addNote()
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {}),
	)

	side := fyne.NewContainerWithLayout(layout.NewBorderLayout(bar, nil, nil, nil),
		bar, u.list)

	split := widget.NewHSplitContainer(side, u.content)
	split.Offset = 0.25
	return split
}

func main() {
	a := app.New()
	w := a.NewWindow("notes")

	list := &noteList{
		list: []*note{
			&note{content: "note1\ncontent1"},
			&note{content: "note2\ncontent2"},
		},
	}

	notesUI := &ui{notes: list}
	w.SetContent(notesUI.loadUI())
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
