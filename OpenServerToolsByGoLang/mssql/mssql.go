package main

import (
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"strconv"
	"strings"
)

func main() {

	Engine = getDataBase()

	mw := &MyMainWindow{model: NewTableViewModel()}
	//var db *walk.DataBinder
	//var ServerStateT *ServerState
	//
	//ServerStateT = &ServerState{111,"dddddddd","dddddddd", 22,2,2,true}

	MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "批量开服工具",
		MinSize:  Size{1500, 800},
		//DataBinder:DataBinder{
		//	AssignTo:&db,
		//	DataSource:	ServerStateT,
		//	ErrorPresenter:ToolTipErrorPresenter{},
		//},
		Layout: HBox{},
		Children: []Widget{

			Composite{
				MinSize: Size{700, 800},
				//ColumnSpan:2,		// 占几个格子
				Layout: HBox{},
				Children: []Widget{
					TableView{
						AssignTo:              &mw.tv,
						Model:                 mw.model,
						AlternatingRowBGColor: walk.RGB(239, 239, 239),
						CheckBoxes:            true,
						ColumnsOrderable:      true,
						MultiSelection:        true,
						Columns: []TableViewColumn{
							{Title: "ServerID"},
							{Title: "ServerName", Width: 240},
							{Title: "机器码", Width: 100, Alignment: AlignFar},
							{Title: "机器人数量", Width: 80, Alignment: AlignFar},
							{Title: "房间类型", Width: 250, Alignment: AlignFar},
							{Title: "机器人类型", Width: 200, Alignment: AlignFar},
						},

						StyleCell: func(style *walk.CellStyle) {
							//item := mw.model.items[style.Row()]
							//
							//if item.checked {
							//	if style.Row()%2 == 0 {
							//		style.BackgroundColor = walk.RGB(159, 215, 255) //相邻行设置不同的颜色
							//	} else {
							//		style.BackgroundColor = walk.RGB(143, 199, 239)
							//	}
							//}
							//
							//switch style.Col() {
							//case 1, 2:
							//	if item.ServerState == 0 {
							//		style.TextColor = walk.RGB(0, 0, 0) //
							//		//style.Image = goodIcon
							//	} else {
							//		style.TextColor = walk.RGB(255, 0, 0) // 红色
							//		//style.Image = badIcon
							//	}
							//}
						},

						//OnSelectedIndexesChanged: func() {
						//			fmt.Println("OnSelectedIndexesChanged 点击", mw.tv.CurrentIndex())
						//},

						//OnItemActivated: func() {
						//	fmt.Println("OnItemActivated 点击", mw.tv.CurrentIndex())
						//},

						// 点击选择tableview的一行
						OnCurrentIndexChanged: func() {
							//fmt.Println("OnCurrentIndexChanged  点击", mw.tv.CurrentIndex())
							// 首先清理掉之前的选择
							for _, v := range mw.model.items {
								v.checked = false
							}
							i := mw.tv.CurrentIndex()
							//fmt.Println("",i)
							if i >= len(mw.model.items) || i < 0 {
								//fmt.Println("越界啦")
								return
							}
							mw.model.items[i].checked = true
							mw.model.PublishRowsReset()

							mw.teid.SetText(strconv.Itoa(mw.model.items[i].ServerId))
							mw.temachine.SetText(mw.model.items[i].ServerMachine)
							mw.teandorid.SetText(strconv.Itoa(mw.model.items[i].AndroidCount))

							//ServerStateT = mw.model.items[i]
							//db.SetDataSource(ServerStateT)
							mw.tegametype.SetCurrentIndex(mw.model.items[i].GameRoomListIndex)
							mw.tesqltype.SetCurrentIndex(mw.model.items[i].SqlFileRobotListIndex)

						},
					},
				},
			},

			VSplitter{
				Children: []Widget{
					PushButton{
						Text: "增加一个房间",
						OnClicked: func() {

							server := &ServerState{ServerID, "ss", ServerMachine, AndroidNum, 0, 0, false}
							ServerID ++
							mw.model.InsertRows(server)
						},
					},
					PushButton{
						Text: "删除一个房间",
						OnClicked: func() {
							i := mw.tv.CurrentIndex()
							if i >= 0 {
								// 如果选择的列表index正确,那么删除这个位置的元素

								mw.model.DeleteRows(i)
							} else {
								walk.MsgBox(mw, "提示", "需要选择一个进行删除!", walk.MsgBoxIconInformation)
							}
						},
					},
					PushButton{
						Text: "保存到数据库",
						OnClicked: func() {

							cmd := walk.MsgBox(mw, "提示", "你已经检查好，准备开始执行数据库操作了么？", walk.MsgBoxYesNo)
							if cmd == walk.DlgCmdYes{
								go func() {
									for index :=range mw.model.items{
										CreateRoomSql(mw.model.items[index].GameRoomListIndex, mw.model.items[index].ServerId, mw.model.items[index].ServerName,mw.model.items[index].ServerMachine)
										CreateGameStore( mw.model.items[index].ServerId, mw.model.items[index].AndroidCount)
										CreateAndroid(mw.model.items[index].SqlFileRobotListIndex, mw.model.items[index].ServerId, mw.model.items[index].AndroidCount)
									}
								}()

							}
							//index := GetSelectIndex(mw)


							//if index >= 0 {
							//	CreateRoomSql(mw.model.items[index].GameRoomListIndex, mw.model.items[index].ServerId, mw.model.items[index].ServerName,mw.model.items[index].ServerMachine)
							//}else {
							//	walk.MsgBox(mw, "提示", "需要选择一个进行删除!", walk.MsgBoxIconInformation)
							//}
							//fmt.Println("开始保存到数据库中！")
							//saveDataToSql(mw.model.ip)
							//fmt.Println("end！")
						},
					},
					PushButton{
						Text: "更新当前修改",
						OnClicked: func() {
							index := GetSelectIndex(mw)
							//fmt.Println(" " ,index)

							if index >= 0 {

								//err := db.Submit()
								//if err != nil{
								//	fmt.Println("error:",err)
								//}
								//fmt.Printf("%v", ServerStateT)
								//fmt.Println("")
								//mw.model.UpdateRows()

								// 将填写的各个输入框的内容保存到tableview中
								ServerId, _ := strconv.Atoi(mw.teid.Text())
								android, _ := strconv.Atoi(mw.teandorid.Text())
								Version = mw.teversion.Text()
								mw.model.UpdateRows(ServerId, mw.temachine.Text(), android, mw.tegametype.CurrentIndex(), mw.tesqltype.CurrentIndex(), index)

								// 根据版本号生成每个服务器的名字
								for _, v := range mw.model.items {
									list := GameRoomListSelect()
									name := list[v.GameRoomListIndex].Name
									splits := strings.Split(name, ".")
									v.ServerName = splits[0] + "[" + strconv.Itoa(v.ServerId) + "]_" + Version
									v.checked = false
								}

								mw.model.PublishRowsReset()

								// 清空右侧的填写项
								mw.teid.SetText("")
								mw.temachine.SetText("")
								mw.teandorid.SetText("")


							} else {
								walk.MsgBox(mw, "提示", "需要选择一个进行修改!", walk.MsgBoxIconInformation)
							}
						},
					},
					HSplitter{

						Children: []Widget{
							Label{Text: "服务器ID：",},
							LineEdit{AssignTo: &mw.teid, Text: ""},
						},},
					HSplitter{
						Children: []Widget{
							Label{Text: "机器码：",},
							LineEdit{AssignTo: &mw.temachine, Text: ""},
						},},
					HSplitter{
						Children: []Widget{
							Label{Text: "机器人数：",},
							LineEdit{AssignTo: &mw.teandorid, Text: ""},
						},},
					HSplitter{
						Children: []Widget{
							Label{Text: "版本号：",},
							LineEdit{AssignTo: &mw.teversion, Text: Version},
						},},

					Label{Text: "选择房间类型：",},
					ComboBox{
						AssignTo:      &mw.tegametype,
						Value:         0,
						BindingMember: "Id",
						DisplayMember: "Name",
						Model:         GameRoomListSelect(),
					},
					Label{Text: "选择机器人类型：",},
					ComboBox{
						AssignTo:      &mw.tesqltype,
						Value:         0,
						BindingMember: "Id",
						DisplayMember: "Name",
						Model:         SqlFileRobotListSelect(),
					},

					Label{Text: "Version:1.01",},
				},
			},
			//	},
			//},
		},
	}.Run()

}

func GetSelectIndex(mw *MyMainWindow) int {
	var index int
	index = -1
	// 获取到当前选择的是哪行
	for i := range mw.model.items {
		if mw.model.items[i].checked == true {
			index = i
			break
		}
	}
	return index
}

//------------------------------------------------------------
// 获取数据引擎， 读取配置文件
//------------------------------------------------------------
func getDataBase() *xorm.Engine {

	// 读取配置文件
	f, err := ini.Load("Setting.ini")
	if err != nil {
		fmt.Println("ini配置文件出错！", err)
		log.Fatal(err)
		return nil
	}
	ServerIP := f.Section("SQL").Key("ServerIP").Value()
	Database := f.Section("SQL").Key("Database").Value()
	uid := f.Section("SQL").Key("uid").Value()
	pwd := f.Section("SQL").Key("pwd").Value()

	ServerMachine = f.Section("Room").Key("ServerMachine").Value()
	ServerID, _ = f.Section("Room").Key("ServerId").Int()
	AndroidNum, _ = f.Section("Room").Key("AndroidNum").Int()
	Version = f.Section("Room").Key("Version").Value()

	fmt.Println("读取配置文件！")
	fmt.Println("开始连接数据库...")

	engine, err := xorm.NewEngine("mssql", "server="+ServerIP+";user id="+uid+";password="+pwd+";Database="+Database)
	//engine, err := xorm.NewEngine("mysql", uid+":"+pwd+"@tcp("+ServerIP+")/"+Database+"?charset=utf8")

	err = engine.Ping()
	if err != nil {
		fmt.Println("数据库连接出错！", err)
		log.Fatal(err)
		return nil
	}

	fmt.Println("数据库连接成功！")

	//engine.ShowSQL(true)
	engine.SetMapper(core.SameMapper{})
	return engine
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
	m.items = make([]*ServerState, 0)

	//now := time.Now()

	//for i,v := range ServerListAll {
	//	m.items[i] = &ServerState{
	//		ServerId:    v.ServerId,
	//		ServerName:  v.ServerName,
	//		ServerState: v.ServerState,
	//		Online:v.Online,
	//		Cpu:v.Cpu,
	//		Memory:v.Memory,
	//		IoRead:v.IoRead,
	//		IoWrite:v.IoWrite,
	//		//Quux:  time.Unix(rand.Int63n(now.Unix()), 0),
	//	}
	//}

	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()
	//m.Sort(m.sortColumn, m.sortOrder)
}

// 更新tableview一行
func (m *TableViewModel) UpdateRows(ServerId int, ServerMachine string, AndroidCount int, GameRoomListIndex int, SqlFileRobotListIndex int, updateIndex int) {
	m.items[updateIndex].ServerId = ServerId

	m.items[updateIndex].ServerMachine = ServerMachine
	m.items[updateIndex].AndroidCount = AndroidCount

	m.items[updateIndex].GameRoomListIndex = GameRoomListIndex
	m.items[updateIndex].SqlFileRobotListIndex = SqlFileRobotListIndex

	m.PublishRowChanged(0)
}

// 添加tableview一行
func (m *TableViewModel) InsertRows(server *ServerState) {
	m.items = append(m.items, server)
	m.PublishRowsInserted(0, 0)
}

// 删除tableview一行
func (m *TableViewModel) DeleteRows(delIndex int) {
	m.items = append(m.items[:delIndex], m.items[delIndex+1:]...)
	m.PublishRowsRemoved(0, 0)
}

// 获取所有选中的项
func (m *TableViewModel) GetSelectRows() []int {
	checkList := make([]int, 0)
	for i := range m.items {
		if m.Checked(i) {
			checkList = append(checkList, i)
		}
	}
	return checkList
}
