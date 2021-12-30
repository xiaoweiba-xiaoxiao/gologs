package logs

import (
	"errors"
	"testing"
)

func TestFatal(t *testing.T) {
	logger1 := NewLogger(FILE, "test.log")
	logger2 := newConsoleLog()
	errMsg := errors.New("this is a test")
	Debug(logger1,errMsg)
	Debug(logger2,errMsg)
	Trace(logger1,errMsg)
	Trace(logger2,errMsg)
	Warring(logger1,errMsg)	
	Warring(logger2,errMsg)
}
