package logs

import (
	log "github.com/cihub/seelog"
	"fmt"
)



// Warnln("ffffff%d--%d",44,44556)
func Warnln(arg0 string, args ...interface{})  {
	s:= fmt.Sprintf(arg0,args...)
	log.Warn(s)
	defer log.Flush()
}

func Warn(arg0 string, args ...interface{})  {
	s:= fmt.Sprintf(arg0,args...)
	log.Warn(s)
	defer log.Flush()
}

func Debug(arg0 string, args ...interface{})  {
	s:= fmt.Sprintf(arg0,args...)
	log.Debug(s)
	defer log.Flush()
}


func Infoln(arg0 interface{}, args ...interface{}) {
	s:= fmt.Sprintf( arg0.(string),args...)
	log.Info(s)
	defer log.Flush()
}
func Info(arg0 string, args ...interface{}) {
	s:= fmt.Sprintf(arg0,args...)
	log.Info(s)
	defer log.Flush()
}

func Error(arg0 string, args ...interface{}) {
	s:= fmt.Sprintf(arg0,args...)
	log.Error(s)
	defer log.Flush()
}
func Critical(arg0 string, args ...interface{}) {
	s:= fmt.Sprintf(arg0,args...)
	log.Critical(s)
	defer log.Flush()
}
func Panicln(arg0 string, args ...interface{}) {
	s:= fmt.Sprintf(arg0,args...)
	log.Critical(s)
	defer log.Flush()
}




func InitLogs(){
	//加载配置文件
	logger, err := log.LoggerFromConfigAsFile("./core/logs/seelog.xml")
	if err!=nil{
		fmt.Println("parse config.xml error")
	}
	log.ReplaceLogger(logger)
}

//func main()  {
//	initLogs()
//	Warnln("ffffff%d--%d",44,44556)
//
//}