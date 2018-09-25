package main

import (
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"




	"runtime"
	"sort"
	"fmt"
)

//--------------------------------------------------------------------
// 类型定义
//--------------------------------------------------------------------

type MyMainWindow struct {
	*walk.MainWindow
	model *TableViewModel
	tv  *walk.TableView
	te    *walk.TextEdit
}


type TableViewModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*ServerState
}

type ServerState struct {
	ServerId    int
	ServerName  string
	ServerState int			//服务器的状态
	Online int
	Memory int
	Cpu int
	IoRead int
	IoWrite int

	//Quux    time.Time
	checked bool

	Daemon int
}

var mwGlobal *MyMainWindow
var ServerListAll []ServerState
var outTE *walk.TextEdit

//--------------------------------------------------------------------
//  window
//--------------------------------------------------------------------
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //设置cpu的核的数量，从而实现高并发

	ServerListAll = make([]ServerState, 0)

	// 开启协程运行TCP处理
	go startTcp()

	// 开启监控游戏状态
	go watchGameServer()
	//for {
	//	time.Sleep(1*time.Second)
	//}
	startUI()

	defer func() {
		if e := recover(); e != nil {
			logerDump()
		}
	}()

}

func startUI() {
	defer func() {
		if e := recover(); e != nil {
			logerDump()
		}
	}()


	// 创建UI
	mw := &MyMainWindow{}
	mwGlobal = mw

	mw.model = NewTableViewModel()


	mWin := MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "游戏服务器监控工具",
		MinSize:  Size{1200, 800},
		Layout:   HBox{},
		Children: []Widget{
			Composite{
				MinSize:Size{700,800},
				//ColumnSpan:2,		// 占几个格子
				Layout: HBox{},
				Children: []Widget{
					TableView{
						AssignTo:              &mw.tv,
						Model: mw.model,
						AlternatingRowBGColor: walk.RGB(239, 239, 239),
						CheckBoxes:            true,
						ColumnsOrderable:      true,
						MultiSelection:        true,
						Columns: []TableViewColumn{
							{Title: "ServerID"},
							{Title: "ServerName", Width: 240},
							{Title: "ServerState", Width: 50,Alignment: AlignFar},
							{Title: "Online", Width: 50,Alignment: AlignFar},
							{Title: "Cpu(%)", Width: 50,Alignment: AlignFar},
							{Title: "Memory(M)", Width: 100,Alignment: AlignFar},
							{Title: "IoRead(M)", Width: 100,Alignment: AlignFar},
							{Title: "IoWrite(M)", Width: 100,Alignment: AlignFar},
							{Title: "Daemon", Width: 50,Alignment: AlignFar},
						},

						StyleCell: func(style *walk.CellStyle) {
							item := mw.model.items[style.Row()]

							if item.checked {
								if style.Row()%2 == 0 {
									style.BackgroundColor = walk.RGB(159, 215, 255)				//相邻行设置不同的颜色
								} else {
									style.BackgroundColor = walk.RGB(143, 199, 239)
								}
							}

							switch style.Col() {
							case 1,2:
								if item.ServerState == 0 {
									style.TextColor = walk.RGB(0, 0, 0)//
									//style.Image = goodIcon
								} else {
									style.TextColor = walk.RGB(255, 0, 0)// 红色
									//style.Image = badIcon
								}
							//case 1:
								//if canvas := style.Canvas(); canvas != nil {
								//	bounds := style.Bounds()
								//	bounds.X += 2
								//	bounds.Y += 2
								//	bounds.Width = int((float64(bounds.Width) - 4) / 5 * float64(len(item.ServerName)))
								//	bounds.Height -= 4
								//	//canvas.DrawBitmapPartWithOpacity(barBitmap, bounds, walk.Rectangle{0, 0, 100 / 5 * len(item.ServerName), 1}, 127)
								//
								//	bounds.X += 4
								//	bounds.Y += 2
								//	canvas.DrawText(item.ServerName, mw.tv.Font(), 0, bounds, walk.TextLeft)
								//}

							//case 2:   // 服务器状态
							//	if item.ServerState == 0 {
							//		style.TextColor = walk.RGB(0, 191, 0)		//绿色
							//		//style.Image = goodIcon
							//	} else {
							//		style.TextColor = walk.RGB(255, 0, 0)		// 红色
							//		//style.Image = badIcon
							//	}

							//case 3:
							//	if item.Quux.After(time.Now().Add(-365 * 24 * time.Hour)) {
							//		//style.Font = boldFont
							//	}
							}
						},

						OnSelectedIndexesChanged: func() {


						},
					},

				},
			},
			VSplitter{
				//ColumnSpan:0,
				//Layout: Grid{Columns: 1},
				Children: []Widget{
					PushButton{
						Text:"刷新列表",
						OnClicked: func() {
							mwGlobal.model.PublishRowsReset()		//5秒刷新一次列表
							mwGlobal.model.PublishRowChanged(0)
						},
					},
					PushButton{
						Text: "全选",
						OnClicked: func() {
							for _,v:=range mw.model.items{
								v.checked = true
							}
							//mw.model.ResetRows()
							mw.model.PublishRowsReset()
						},
					},
					PushButton{
						Text: "取消全选",
						OnClicked: func() {
							for _,v:=range mw.model.items{
								v.checked = false
							}
							mw.model.PublishRowsReset()
						},
					},
					PushButton{
						Text: "...",
						OnClicked: func() {

						},
					},
					//---------------------------------------------------------------------------
					// BY
					//---------------------------------------------------------------------------
					//PushButton{
					//	Text: "服务器维护",
					//	OnClicked: func() {
					//		checkList:= mw.model.GetSelectRows()
					//		if len(checkList) >= 0 {
					//			cmd := walk.MsgBox(mw, "提示", "确定进行服务器维护么？", walk.MsgBoxYesNo)
					//			if cmd == walk.DlgCmdYes {
					//				for _, v := range checkList {
					//					send_gm_cmd(mw.model.items[v].ServerId, "@RoomClose")
					//				}
					//			}
					//		} else {
					//			walk.MsgBox(mw, "提示", "需要选择一个游戏服务器进行通知!", walk.MsgBoxIconInformation)
					//		}
					//	},
					//},
					//PushButton{
					//	Text: "解除维护",
					//	OnClicked: func() {
					//		checkList := mw.model.GetSelectRows()
					//		if len(checkList) >= 0 {
					//			cmd := walk.MsgBox(mw, "提示", "确定解除维护么？", walk.MsgBoxYesNo)
					//			if cmd == walk.DlgCmdYes {
					//				for _, v := range checkList {
					//					send_gm_cmd(mw.model.items[v].ServerId, "@RoomOpen")
					//				}
					//			}
					//		} else {
					//			walk.MsgBox(mw, "提示", "需要选择一个游戏服务器进行通知!", walk.MsgBoxIconInformation)
					//		}
					//	},
					//},
					//PushButton{
					//	Text: "踢出所有玩家",
					//	OnClicked: func() {
					//		checkList := mw.model.GetSelectRows()
					//		if len(checkList) >= 0 {
					//			cmd := walk.MsgBox(mw, "提示", "确定踢出所有玩家么？", walk.MsgBoxYesNo)
					//			if cmd == walk.DlgCmdYes {
					//				for _, v := range checkList {
					//					send_gm_cmd(mw.model.items[v].ServerId, "@KickALLUser")
					//				}
					//			}
					//		} else {
					//			walk.MsgBox(mw, "提示", "需要选择一个游戏服务器进行通知!", walk.MsgBoxIconInformation)
					//		}
					//	},
					//},
					//PushButton{
					//	Text: "踢出所有超过3个小时的玩家",
					//	OnClicked: func() {
					//		checkList:= mw.model.GetSelectRows()
					//		if len(checkList) >= 0 {
					//			cmd := walk.MsgBox(mw, "提示", "确定踢出所有超过3个小时的玩家么？", walk.MsgBoxYesNo)
					//			if cmd == walk.DlgCmdYes {
					//				for _, v := range checkList {
					//					send_gm_cmd(mw.model.items[v].ServerId, "@KickALLUserTimeSoLong")
					//				}
					//			}
					//		} else {
					//			walk.MsgBox(mw, "提示", "需要选择一个游戏服务器进行通知!", walk.MsgBoxIconInformation)
					//		}
					//	},
					//},
					//PushButton{
					//	Text: "刷新IP白名单",
					//	OnClicked: func() {
					//		checkList:= mw.model.GetSelectRows()
					//		if len(checkList) >= 0 {
					//			cmd := walk.MsgBox(mw, "提示", "确定刷新IP白名单么？", walk.MsgBoxYesNo)
					//			if cmd == walk.DlgCmdYes {
					//				for _, v := range checkList {
					//					send_gm_cmd(mw.model.items[v].ServerId, "@RefreshIpList")
					//				}
					//			}
					//		} else {
					//			walk.MsgBox(mw, "提示", "需要选择一个游戏服务器进行通知!", walk.MsgBoxIconInformation)
					//		}
					//	},
					//},
					//PushButton{
					//	Text: "排行榜奖励发放(测试使用)",
					//	OnClicked: func() {
					//		f, _ := ini.Load("Setting.ini")
					//		rankID := f.Section("TestRank").Key("rankID").Value()
					//		timeStart := f.Section("TestRank").Key("timeStart").Value()
					//		timeEnd := f.Section("TestRank").Key("timeEnd").Value()
					//		send_gm_cmd(-100, "@SendRankAward "+rankID+" "+timeStart+" "+timeEnd)
					//	},
					//},
					//---------------------------------------------------------------------------
					//
					//---------------------------------------------------------------------------
					PushButton{
						Text: "...",
						OnClicked: func() {

						},
					},
					PushButton{
						Text: "...",
						OnClicked: func() {
							//mw.model.items[0].ServerState = 1 - mw.model.items[0].ServerState
							//mw.model.PublishRowChanged(0)
						},
					},
						PushButton{
						Text: "...",
						OnClicked: func() {
							//mw.model.items[0].ServerState = 1 - mw.model.items[0].ServerState
							//mw.model.PublishRowChanged(0)
						},
					},
					PushButton{
						Text: "...",
						OnClicked: func() {

							//mw.model.InsertRows(&ServerState{
							//	ServerId:2,ServerName:"22",ServerState:1,})

						},
					},
					PushButton{
						Text: "...",
						OnClicked: func() {
							//i := 3
							//mw.model.DeleteRows(i)

						},
					},
					TextEdit{AssignTo: &outTE,
						Text: "", Enabled: false},
				},
			},
		},
	}
	mWin.Run()
}

func ShowAllServerNum() int {
	num := 0
	num = len(mwGlobal.model.items)

	if outTE !=nil {
		outTE.SetText(fmt.Sprintf("版本号：V1.16  (增加中心服守护)    服务器总数：%d", num))
	}
	return 0
}

//--------------------------------------------------------------------
// model
//--------------------------------------------------------------------
func NewTableViewModel() *TableViewModel {
	m := new(TableViewModel)
	m.ResetRows()
	return m
}
// 刷新全部列表
func (m *TableViewModel) ResetRows() {
	// Create some random data.
	m.items = make([]*ServerState, len(ServerListAll))

	//now := time.Now()

	for i,v := range ServerListAll {
		m.items[i] = &ServerState{
			ServerId:    v.ServerId,
			ServerName:  v.ServerName,
			ServerState: v.ServerState,
			Online:v.Online,
			Cpu:v.Cpu,
			Memory:v.Memory,
			IoRead:v.IoRead,
			IoWrite:v.IoWrite,
			//Quux:  time.Unix(rand.Int63n(now.Unix()), 0),
			Daemon:v.Daemon,
		}
	}

	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()
	//m.Sort(m.sortColumn, m.sortOrder)
}

// 更新tableview一行
func (m *TableViewModel) UpdateRows(state int, online int, cpu int, mem int, read int ,write int, updateIndex int) {
	m.items[updateIndex].ServerState = state
	m.items[updateIndex].Online = online

	m.items[updateIndex].Cpu = cpu
	m.items[updateIndex].Memory = mem/1024/1024
	m.items[updateIndex].IoRead = read/1024/1024
	m.items[updateIndex].IoWrite = write/1024/1024
	//m.PublishRowChanged(0)
}

// 添加tableview一行
func (m *TableViewModel) InsertRows(server *ServerState) {
	m.items = append(m.items , server)
	m.PublishRowsInserted(0,0)
}
// 删除tableview一行
func (m *TableViewModel) DeleteRows( delIndex int ) {
	m.items = append(m.items[:delIndex] , m.items[delIndex+1:]...)
	m.PublishRowsRemoved(0,0)
}

// 获取所有选中的项
func (m *TableViewModel) GetSelectRows() []int {
	checkList := make([]int,0)
	for i :=range m.items{
		if m.Checked(i){
			checkList = append(checkList, i)
		}
	}
	return checkList
}

//--------------------------------------------------------------------
// 模板继承
//--------------------------------------------------------------------

// Called by the TableView from SetModel and every time the model publishes a
// RowsReset event.
func (m *TableViewModel) RowCount() int {
	return len(m.items)
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *TableViewModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.ServerId

	case 1:
		return item.ServerName

	case 2:
		return item.ServerState

	case 3:
		return item.Online
	case 4:
		return item.Cpu
	case 5:
		return item.Memory
	case 6:
		return item.IoRead
	case 7:
		return item.IoWrite
	case 8:
		return item.Daemon
	}

	panic("unexpected col")
}

// Called by the TableView to retrieve if a given row is checked.
func (m *TableViewModel) Checked(row int) bool {
	return m.items[row].checked
}

// Called by the TableView when the user toggled the check box of a given row.
func (m *TableViewModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked

	return nil
}




// Called by the TableView to sort the model.
func (m *TableViewModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.SliceStable(m.items, func(i, j int) bool {
		a, b := m.items[i], m.items[j]

		c := func(ls bool) bool {
			if m.sortOrder == walk.SortAscending {
				return ls
			}

			return !ls
		}

		switch m.sortColumn {
		case 0:
			return c(a.ServerId < b.ServerId)

		case 1:
			return c(a.ServerName < b.ServerName)

		case 2:
			return c(a.ServerState < b.ServerState)

		case 3:
			return c(a.Online < b.Online)

		case 4:
			return c(a.Cpu < b.Cpu)
		case 5:
			return c(a.Memory < b.Memory)
		case 6:
			return c(a.IoWrite < b.IoWrite)
		case 7:
			return c(a.IoRead < b.IoRead)
		case 8:
			return c(a.Daemon < b.Daemon)
		}

		panic("unreachable")
	})

	return m.SorterBase.Sort(col, order)
}
