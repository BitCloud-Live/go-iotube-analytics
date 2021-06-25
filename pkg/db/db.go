package db

import "github.com/IoTube-analytics/go-iotube-analytics/pkg/format"

const ComponentName = "db"

type Config struct {
	LogLevel      string
	Path          string
	RemoteTimeout format.Duration
}
