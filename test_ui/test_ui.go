package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strings"
	"fmt"
	"os"
)

type MyMainWindow struct {
	*walk.MainWindow
	model *EnvModel
	lb    *walk.ListBox
	te    *walk.TextEdit
}

//listbox的点击选择事件
func (mw *MyMainWindow) lb_CurrentIndexChanged() {
	i := mw.lb.CurrentIndex()
	item := &mw.model.items[i]		//获取数据
	mw.te.SetText(item.value)		//设置文本框的显示变化
}
// listbox的双击事件
func (mw *MyMainWindow) lb_ItemActivated() {
	value := mw.model.items[mw.lb.CurrentIndex()].value
	walk.MsgBox(mw, "Value", value, walk.MsgBoxIconInformation)
}

type EnvItem struct {
	name  string
	value string
}

type EnvModel struct {
	walk.ListModelBase
	items []EnvItem
}
// 列表的数据
func NewEnvModel() *EnvModel {
	env := os.Environ()
	m := &EnvModel{items: make([]EnvItem, len(env))}		//分配item[]的内存
	for i, e := range env {
		j := strings.Index(e, "=")			//找到=，做分割
		if j == 0 {
			continue
		}
		name := e[0:j]						// = 前面
		value := strings.Replace(e[j+1:], ";", "\r\n", -1)		//= 后面内容
		m.items[i] = EnvItem{name, value}
	}
	return m
}
func (m *EnvModel) ItemCount() int {
	return len(m.items)
}

func (m *EnvModel) Value(index int) interface{} {
	return m.items[index].name
}




func main() {
	var inTE, outTE *walk.TextEdit
	mw := &MyMainWindow{model: NewEnvModel()}

	//data, _ := json.MarshalIndent(mw, "", " ")
	//fmt.Printf("%s\n", data)
	//fmt.Println("---------------------------------------------")

	MainWindow{
		AssignTo: &mw.MainWindow,
		Title:   "test window",
		MinSize: Size{600, 600},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					//TextEdit{AssignTo: &outTE, ReadOnly: true},
					ListBox{
						AssignTo: &mw.lb,
						Model:    mw.model,
						OnCurrentIndexChanged: mw.lb_CurrentIndexChanged,		//当选择的index变化的时候， 选择事件
						OnItemActivated:       mw.lb_ItemActivated,				//双击事件，如果不需要的话，可以设置nil
					},
				},
				MinSize: Size{Width:600 , Height: 400},
			},
			HSplitter{
				Children:[]Widget{
					Label{Text:"左面"},
					Label{Text:"中间"},
					Label{Text:"右面"},
				},

			},
			TextEdit{AssignTo: &mw.te,MinSize: Size{Width:600 , Height: 60},},			//这个编辑框用来显示选择的listbox的信息，lb_CurrentIndexChanged函数里面修改这个显示值
			PushButton{
				Text: "拷贝到另外一面",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
					fmt.Println(inTE.Text())
				},
			},
		},
	}.Run()
}


