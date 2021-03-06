package rpcx

import (
	"sync"

	"github.com/yakaa/cuter/lib/logx"

	"google.golang.org/grpc/grpclog"
)

var once sync.Once

type Logger struct{}

func InitLogger() {
	once.Do(func() {
		grpclog.SetLoggerV2(new(Logger))
	})
}

func (l *Logger) Error(args ...interface{}) {
	logx.Error(args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	logx.Errorf(format, args...)
}

func (l *Logger) Errorln(args ...interface{}) {
	logx.Error(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	logx.Error(args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	logx.Errorf(format, args...)
}

func (l *Logger) Fatalln(args ...interface{}) {
	logx.Error(args...)
}

func (l *Logger) Info(args ...interface{}) {
	// ignore builtin grpc info
}

func (l *Logger) Infoln(args ...interface{}) {
	// ignore builtin grpc info
}

func (l *Logger) Infof(format string, args ...interface{}) {
	// ignore builtin grpc info
}

func (l *Logger) V(v int) bool {
	return false
}

func (l *Logger) Warning(args ...interface{}) {
	logx.Error(args...)
}

func (l *Logger) Warningln(args ...interface{}) {
	logx.Error(args...)
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	logx.Errorf(format, args...)
}
