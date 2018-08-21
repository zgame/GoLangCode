package main

import (
	"log"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var mw *walk.MainWindow
	//var outTE *walk.TextEdit
	//
	//animal := new(Animal)

	if _, err := (MainWindow{
		AssignTo: &mw,
		Title:    "Walk Data Binding Example",
		MinSize:  Size{700, 500},
		//Layout:   VBox{},
		//Layout: HBox{},
		Layout: Grid{Columns: 8},
		Children: []Widget{
			Composite{
				ColumnSpan:6,
				Layout: Grid{Columns:0},
				Children: []Widget{
					TableView{
						//AssignTo:              &tv,
						AlternatingRowBGColor: walk.RGB(239, 239, 239),
						CheckBoxes:            true,
						ColumnsOrderable:      true,
						MultiSelection:        true,
						Columns: []TableViewColumn{
							{Title: "#"},
							{Title: "Bar"},
							{Title: "Baz", Alignment: AlignFar},
							{Title: "Quux", Format: "2006-01-02 15:04:05", Width: 150},
						},
					},
				},
			},
			Composite{
				ColumnSpan:1,
				Layout: Grid{Columns: 1},
				Children: []Widget{
					Label{Text: "dddddddddd"},
					PushButton{
						Text: "Edit Animal",
						OnClicked: func() {
						},
					},
					PushButton{
						Text: "Edit Animal",
						OnClicked: func() {
						},
					},
					PushButton{
						Text: "Edit Animal",
						OnClicked: func() {
						},
					},
					PushButton{
						Text: "Edit Animal",
						OnClicked: func() {
						},
					},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
				},
			},
			Composite{
				ColumnSpan:1,
				Layout: VBox{},
				Children: []Widget{
					Label{Text: "dddddddddd"},
					PushButton{
						Text: "Edit Animal",
						OnClicked: func() {
						},
					},
					PushButton{
						Text: "Edit Animal",
						OnClicked: func() {
						},
					},
					PushButton{
						Text: "Edit Animal",
						OnClicked: func() {
						},
					},
					PushButton{
						Text: "Edit Animal",
						OnClicked: func() {
						},
					},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
					Label{Text: "dddddddddd"},
				},
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}
