package logs

import (
	"bufio"
	"fmt"
	"os"
)

/*
日志接口
*/
type Logs interface {
	DeBug(logMsg error)
	Trace(logMsg error)
	Info(logMsg error)
	Warring(logmsg error)
	Error(logMsg error)
	Fatal(logMsg error)	
}

/*
日志文件
*/
type FileLog struct {
	Logpath string `json:logpath`
}

/*
终端日志
*/
type ConsoleLog struct {
}

/*
定义终端日志和文件日志的常量
*/
const (
	CONSOLE = iota
	FILE
)

/*
返回一个接口
*/
func NewLogger(style int, path string) Logs {
	switch style {
	case CONSOLE:
		var logs Logs
		consolelog := newConsoleLog()
		logs = consolelog
		return logs
	case FILE:
		var logs Logs
		if path == "" {
			path = "/var/log/TennetlDB.log"
		}
		fileLog := newFileLog(path)
		logs = fileLog
		return logs
	default:
		return nil
	}
}

/*
创建一个文件日志
*/
func newFileLog(path string) *FileLog {
	return &FileLog{
		Logpath: path,
	}
}

/*
创建一个终端日志
*/
func newConsoleLog() *ConsoleLog {
	return &ConsoleLog{}
}

func (flog *FileLog)writeTo(level int,logMsg error){
	timeStr := logtime()
	file, err := os.OpenFile(flog.Logpath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	writer := bufio.NewWriter(file)
	levelstr := getLevel(level)
	str := fmt.Sprintf("%s: [%s] -%v\n", timeStr, levelstr, logMsg)
	_,err = writer.Write([]byte(str))
	if err !=nil {
		return
	}
	err = writer.Flush()
	if err != nil {
		return
	}
	defer file.Close()
}

/*
filelog 
*/

func (flog *FileLog)DeBug(errMsg error) {
	flog.writeTo(DEBUG,errMsg)
}

func (flog *FileLog)Trace(errMsg error){
	flog.writeTo(TRACE,errMsg)
}

func (flog *FileLog)Info(errMgs error){
	flog.writeTo(INFO,errMgs)
}

func (flog *FileLog)Warring(errMsg error){
	flog.writeTo(WARRING,errMsg)
}

func (flog *FileLog)Error(errMsg error){
	flog.writeTo(ERROR,errMsg)
}

func (flog *FileLog)Fatal(errMsg error) {
	flog.writeTo(FATAL,errMsg)		
}

/*consloelog*/
func (clog *ConsoleLog)DeBug(errMsg error){
	clog.writeTo(DEBUG,errMsg)
}

func (clog *ConsoleLog)Trace(errMsg error){
	clog.writeTo(TRACE,errMsg)
}

func (clog *ConsoleLog)Info(errMsg error){
	clog.writeTo(TRACE,errMsg)
}

func (clog *ConsoleLog)Warring(errMsg error){
	clog.writeTo(WARRING,errMsg)
}

func (clog *ConsoleLog)Error(errMsg error){
	clog.writeTo(ERROR,errMsg)
}

func (clog *ConsoleLog)Fatal(errMgs error){
	clog.writeTo(FATAL,errMgs)
}

/*
终端日志打印
*/
func (clog *ConsoleLog)writeTo(level int,errMsg error){
	timeStr := logtime()
	levelstr := getLevel(level)
	fmt.Fprintf(os.Stdout, "%s: [%s] -%v\n", timeStr, levelstr, errMsg)
}
