package db

import (
	"github.com/tellor-io/telliot/pkg/format"
)

const ComponentName = "db"

type Config struct {
	LogLevel string
	Path     string
}
