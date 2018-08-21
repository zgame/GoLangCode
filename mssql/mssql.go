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
	if  i>=0 && i< len(mw.model.ip) {
		ipp := &mw.model.ip[i] //获取数据
		mw.te.SetText(*ipp)    //设置文本框的显示变化
	}
}

// listbox的双击事件
func (mw *MyMainWindow) lb_ItemActivated() {
	//value := mw.model.ip[mw.lb.CurrentIndex()]
	//walk.MsgBox(mw, "Value", value, walk.MsgBoxIconInformation)
}


type EnvModel struct {
	walk.ListModelBase
	ip []string
}

// 这个是数据库中的表结构，多行返回这个是一个本结构体的数组
type Test_ip struct {
	Ip string`xorm:"varchar(40)"`
}



// 获取列表的显示数据
func NewEnvModel() *EnvModel {
	m := &EnvModel{ip: make([]string, 0)}

	ip_list := getSqlData()
	for _, v := range ip_list {
		m.ip = append(m.ip, v.Ip)
	}

	return m
}
func (m *EnvModel) ItemCount() int {
	return len(m.ip)
}

func (m *EnvModel) Value(index int) interface{} {
	return m.ip[index]
}

// 切片的删除函数，要自己写
func remove_slice(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	mw := &MyMainWindow{model: NewEnvModel()}

	MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "IP白名单修改器",
		MinSize:  Size{500, 200},
		Layout:   VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					ListBox{
						AssignTo:              &mw.lb,
						Model:                 mw.model,
						OnCurrentIndexChanged: mw.lb_CurrentIndexChanged, //当选择的index变化的时候， 选择事件
						OnItemActivated:       mw.lb_ItemActivated,       //双击事件，如果不需要的话，可以设置nil
						MinSize:  Size{370, 200},
						MaxSize:  Size{370, 200},
					},
					VSplitter{
						Children: []Widget{
							PushButton{
								Text: "增加一个IP",
								OnClicked: func() {
									mw.model.ip = append(mw.model.ip, "0.0.0.0")
									mw.model.PublishItemsReset()
								},
							},
							PushButton{
								Text: "删除一个IP",
								OnClicked: func() {
									i := mw.lb.CurrentIndex()
									if i>=0{
										// 如果选择的列表index正确,那么删除这个位置的元素
										mw.model.ip=append(mw.model.ip[:i],mw.model.ip[i+1:]...)
										mw.model.PublishItemsReset()
									}else {
										walk.MsgBox(mw, "提示", "需要选择一个ip进行删除!", walk.MsgBoxIconInformation)
									}
								},
							},
							PushButton{
								Text: "保存到数据库",
								OnClicked: func() {
									fmt.Println("开始保存到数据库中！")
									saveDataToSql(mw.model.ip)
									fmt.Println("end！")
								},
							},
							PushButton{
								Text: "修改当前行的值",
								OnClicked: func() {
									i := mw.lb.CurrentIndex()
									if i>=0 {
										mw.model.ip[i] = mw.te.Text()
										mw.model.PublishItemsReset()
									}else {
										walk.MsgBox(mw, "提示", "需要选择一个ip进行修改!", walk.MsgBoxIconInformation)
									}
								},
							},
							TextEdit{AssignTo: &mw.te,},

						},
					},
				},
			},

		},
	}.Run()

}

// 把列表数据保存到数据库中
func saveDataToSql(list []string){
	engine := getDataBase()
	engine.Query("delete  from test_ip")		// 删除全部

	test_ip_list := make([]Test_ip,0)
	for _ ,v := range list{
		test_ip_list= append(test_ip_list, Test_ip{Ip:v})
	}

	fmt.Printf("%v",test_ip_list)
	fmt.Println("")
	_,err :=engine.Insert(test_ip_list) 		 //插入多条，因为test_iii是结构数组
	if err != nil {
		fmt.Println("保存数据出错！", err)
		log.Fatal(err)
	}

}
// 获取数据引擎
func getDataBase() *xorm.Engine{
	//************************************************************************************
	// 读取配置文件
	f, err := ini.Load("Setting.ini")
	if err != nil {
		fmt.Println("ini配置文件出错！", err)
		log.Fatal(err)
		return nil
	}
	ServerIP := f.Section("author").Key("ServerIP").Value()
	Database := f.Section("author").Key("Database").Value()
	uid := f.Section("author").Key("uid").Value()
	pwd := f.Section("author").Key("pwd").Value()

	engine, err := xorm.NewEngine("mssql", "server="+ServerIP+";user id="+uid+";password="+pwd+";Database="+Database)
	//engine, err := xorm.NewEngine("mysql", uid+":"+pwd+"@tcp("+ServerIP+")/"+Database+"?charset=utf8")
	err = engine.Ping()
	if err != nil {
		fmt.Println("数据库连接出错！", err)
		log.Fatal(err)
		return nil
	}
	engine.ShowSQL(true)
	engine.SetMapper(core.SameMapper{})
	return engine
}

// 从数据库中把数据读出来
func getSqlData() []Test_ip {

	var Engine *xorm.Engine
	Engine = getDataBase()

	has,_ := Engine.IsTableExist(new(Test_ip))

	if has{
		//**************************************************************************************
		// 从sql中获取数据

		err := Engine.Sync2(new(Test_ip))

		var testIpList []Test_ip
		err = Engine.Find(&testIpList)

		if err != nil {
			fmt.Println("数据库查询出错！", err)
			log.Fatal(err)
			return nil
		}
		return testIpList
	}else{
		fmt.Println("表不存在，name创建它")
		Engine.CreateTables(new(Test_ip))
		return nil
	}



}
