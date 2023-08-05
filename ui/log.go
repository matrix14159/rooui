package ui

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Tie *zap.Logger
var Log *zap.SugaredLogger

func initLog() {
	core := zapcore.NewTee(mkConsoleCore())
	Tie = zap.New(core, zap.AddCaller())
	Log = Tie.Sugar()
}

func mkConsoleCore() zapcore.Core {
	lw := &LogWriter{}
	writer := zapcore.Lock(lw)

	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeTime = nil //zapcore.ISO8601TimeEncoder
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.ConsoleSeparator = " "
	cfg.EncodeCaller = zapcore.ShortCallerEncoder
	encoder := zapcore.NewConsoleEncoder(cfg)

	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)
	return core
}

type LogWriter struct{}

var l_debug = []byte("DEBUG")
var l_info = []byte("INFO")
var l_warn = []byte("WARN")
var l_error = []byte("ERROR")

func (*LogWriter) Write(p []byte) (n int, err error) {
	if Console == nil {
		err = fmt.Errorf("console is nil, can't write log. data:%v", string(p))
		return
	}

	if bytes.Contains(p, l_debug) {
		Console.Debug(string(p))
	}
	if bytes.Contains(p, l_info) {
		Console.Log(string(p))
		return
	}
	if bytes.Contains(p, l_warn) {
		Console.Warn(string(p))
		return
	}
	if bytes.Contains(p, l_error) {
		Console.Error(string(p))
		return
	}
	Console.Log(string(p))
	return
}

func (p *LogWriter) Sync() (err error) {
	return
}

// MarkError format an error and log to console
func MarkError(msg string, a ...any) (err error) {
	err = fmt.Errorf(msg, a...)
	_, file, line, ok := runtime.Caller(1)
	if ok {
		Console.Error(fmt.Sprintf("ERROR %v/%v:%v %v",
			filepath.Base(filepath.Dir(file)), filepath.Base(file), line, err))
		return
	}
	Console.Error(fmt.Sprintf("ERROR %v", err))
	return
}
