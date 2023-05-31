package controller

import (
	"hechuangfil/conf"
	log "hechuangfil/logrus"
)

type Log struct {
	fileName string
}

func NewLog(logFile *conf.LogFile) *Log{
	lg := new(Log)
	lg.fileName = logFile.FileName
	log.NewLogger(lg.fileName)
	return lg
}