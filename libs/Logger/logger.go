package Logger

import (
	"os"
	"github.com/op/go-logging"
	"log"
)

var (
	fileName string = "/tmp/zzz.log"
)

type Priority int

const (
	LogTypeStd  = "std"
	LogTypeFile = "file"
)
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

func Debug(params ... string) {

	Dump(LOG_DEBUG, params)
}

func Info(params ... string) {

	Dump(LOG_INFO, params)
}

func Notice(params ... string) {

	Dump(LOG_NOTICE, params)
}

func Warning(params ... string) {

	Dump(LOG_WARNING, params)
}

func Error(params ... string) {
	Dump(LOG_ERR, params)
}

func Critical(params ...string) {
	Dump(LOG_CRIT, params)
}

func Dump(log_priority Priority, params []string) {

	total := 0
	for range params {
		total += 1
	}
	log_type := LogTypeStd
	if total <= 0 {
		log.Print("dump content empty !")
		return
	} else if total > 1 && params[1] != "" {
		log_type = params[1]
	}
	content := params[0]

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
