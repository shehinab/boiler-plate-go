package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"time"
)

type Log struct{}


// Create a new instance of the logger. You can have any number of instances.
var _log = logrus.New()

func (l *Log)WriteFile(msg string,data logrus.Fields)  {


	//You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("src/logs/Log_"+time.Now().Format("2006-01-02") + ".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		_log.Out = file
	} else {
		_log.Info("Failed to log to file, using default stderr")
	}

	_log.WithFields(data).Info(msg)
}

func (l *Log)WriteFileError(msg string,data logrus.Fields)  {


	//You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("src/logs/Log_"+time.Now().Format("2006-01-02") + ".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		_log.Out = file
	} else {
		_log.Info("Failed to log to file, using default stderr")
	}
	pc, fn, line, _ := runtime.Caller(1)

	ErrorInfo := fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
	_log :=_log.WithFields(logrus.Fields{"ErrorInfo":ErrorInfo})
	_log.WithFields(data).Error(msg)
}