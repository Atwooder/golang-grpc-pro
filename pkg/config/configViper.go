package config

import "time"

type ServiceSettings struct {
	RunModel     string
	HttpPort     string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

type LoggerSettings struct {
	LogSavePath        string
	LogDefaultFileName string
	LogFileExt         string
	LogsRotationTime   int
	LogsRotationCount  int
}

type DatabaseSettings struct {
	DBtype       string
	UserName     string
	Password     string
	Host         string
	DbName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}
