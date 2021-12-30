package logs

import (
	"time"
)

/*
返会日志打印时间
*/
func logtime() string{
	return time.Now().Format("2006-01-02 15:04:05")
}

func Debug(log Logs,err error){
	log.DeBug(err)
}

func Trace(log Logs,err error){
	log.Trace(err)
}

func Info(log Logs,err error){
	log.Info(err)
}

func Warring(log Logs,err error){
	log.Warring(err)
}

func Error(log Logs,err error){
	log.Error(err)
}

func Fatal(log Logs,err error){
	log.Fatal(err)
}