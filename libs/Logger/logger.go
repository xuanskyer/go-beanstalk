package Logger

import (
	"reflect"
	"github.com/op/go-logging"
	"os"
	"log"
)

var (
	fileName string = "/tmp/zzz.log"
)

type Priority int

var LogTypeStd interface{} = "std"
var LogTypeFile interface{} = "file"

const (
	// From /usr/include/sys/syslog.h.
	// These are the same on Linux, BSD, and OS X.
	LOG_EMERG   Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

func Debug(params ... interface{}) {

	Dump(LOG_DEBUG, params)
}

func Info(params ... interface{}) {

	Dump(LOG_INFO, params)
}

func Notice(params ... interface{}) {

	Dump(LOG_NOTICE, params)
}

//
func Warning(params ... interface{}) {

	Dump(LOG_WARNING, params)
}

func Error(params ... interface{}) {
	Dump(LOG_ERR, params)
}

func Critical(params ...interface{}) {
	Dump(LOG_CRIT, params)
}

func Dump(log_priority Priority, params interface{}) {

	params_reflected := reflect.ValueOf(params)
	params_count := params_reflected.Len()

	log_type := LogTypeStd
	if params_count <= 0 {
		log.Print("dump content empty !")
		return
	} else if params_count > 1 {
		log_type = params_reflected.Index(1).Interface()
	}
	content := params_reflected.Index(0)

	var logObj = logging.MustGetLogger("example")
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.8s} %{id:09x} %{message}%{color:reset}`,
	)
	var backend = logging.NewLogBackend(os.Stderr, "", 0)
	if LogTypeFile == log_type {

		logFile, err := os.Create(fileName)
		defer logFile.Close()
		if err != nil {
			log.Print("open file error !")
			return
		}
		backend = logging.NewLogBackend(logFile, "", 0)
	}
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)

	switch log_priority {
	case LOG_DEBUG:
		logObj.Debug(content)
	case LOG_NOTICE:
		logObj.Notice(content)
	case LOG_WARNING:
		logObj.Warning(content)
	case LOG_ERR:
		logObj.Error(content)
	case LOG_CRIT:
		logObj.Critical(content)
	default:
		logObj.Info(content)
	}

}
