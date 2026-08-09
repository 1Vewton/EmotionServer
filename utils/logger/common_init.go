package logger

import (
	"log/slog"
)

var handler *slog.TextHandler

// SysLogger manages the whole logging survice
var SysLogger *slog.Logger
