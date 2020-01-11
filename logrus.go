package abstractlogger

import (
	"github.com/sirupsen/logrus"
)

// LogrusLogger implements the Logger frontend using the popular logrus library as a backend
// It makes use of the LevelCheck helper to increase performance
type LogrusLogger struct {
	l          *logrus.Logger
	levelCheck LevelCheck
}

func NewLogrusLogger(l *logrus.Logger,level Level) *LogrusLogger {
	return &LogrusLogger{
		l:          l,
		levelCheck: NewLevelCheck(level),
	}
}

func (l *LogrusLogger) DebugString(msg string, key, value string) {
	if !l.levelCheck.Check(DebugLevel){
		return
	}
	l.l.WithField(key,value).Debug(msg)
}

func (l *LogrusLogger) DebugInt(msg string, key string, value int) {
	if !l.levelCheck.Check(DebugLevel){
		return
	}
	l.l.WithField(key,value).Debug(msg)
}

func (l *LogrusLogger) DebugInt64(msg string, key string, value int64) {
	if !l.levelCheck.Check(DebugLevel){
		return
	}
	l.l.WithField(key,value).Debug(msg)
}

func (l *LogrusLogger) DebugFloat32(msg string, key string, value float32) {
	if !l.levelCheck.Check(DebugLevel){
		return
	}
	l.l.WithField(key,value).Debug(msg)
}

func (l *LogrusLogger) DebugFloat64(msg string, key string, value float64) {
	if !l.levelCheck.Check(DebugLevel){
		return
	}
	l.l.WithField(key,value).Debug(msg)
}

func (l *LogrusLogger) DebugByteString(msg string, key string, value []byte) {
	if !l.levelCheck.Check(DebugLevel){
		return
	}
	l.l.WithField(key,value).Debug(msg)
}

func (l *LogrusLogger) InfoString(msg string, key, value string) {
	if !l.levelCheck.Check(InfoLevel){
		return
	}
	l.l.WithField(key,value).Info(msg)
}

func (l *LogrusLogger) InfoInt(msg string, key string, value int) {
	if !l.levelCheck.Check(InfoLevel){
		return
	}
	l.l.WithField(key,value).Info(msg)
}

func (l *LogrusLogger) InfoInt64(msg string, key string, value int64) {
	if !l.levelCheck.Check(InfoLevel){
		return
	}
	l.l.WithField(key,value).Info(msg)
}

func (l *LogrusLogger) InfoFloat32(msg string, key string, value float32) {
	if !l.levelCheck.Check(InfoLevel){
		return
	}
	l.l.WithField(key,value).Info(msg)
}

func (l *LogrusLogger) InfoFloat64(msg string, key string, value float64) {
	if !l.levelCheck.Check(InfoLevel){
		return
	}
	l.l.WithField(key,value).Info(msg)
}

func (l *LogrusLogger) InfoByteString(msg string, key string, value []byte) {
	if !l.levelCheck.Check(InfoLevel){
		return
	}
	l.l.WithField(key,string(value)).Info(msg)
}

func (l *LogrusLogger) WarnString(msg string, key, value string) {
	if !l.levelCheck.Check(WarnLevel){
		return
	}
	l.l.WithField(key, value).Warn(msg)
}

func (l *LogrusLogger) WarnInt(msg string, key string, value int) {
	if !l.levelCheck.Check(WarnLevel){
		return
	}
	l.l.WithField(key,value).Warn(msg)
}

func (l *LogrusLogger) WarnInt64(msg string, key string, value int64) {
	if !l.levelCheck.Check(WarnLevel){
		return
	}
	l.l.WithField(key,value).Warn(msg)
}

func (l *LogrusLogger) WarnFloat32(msg string, key string, value float32) {
	if !l.levelCheck.Check(WarnLevel){
		return
	}
	l.l.WithField(key, value).Warn(msg)
}

func (l *LogrusLogger) WarnFloat64(msg string, key string, value float64) {
	if !l.levelCheck.Check(WarnLevel){
		return
	}
	l.l.WithField(key, value).Warn(msg)
}

func (l *LogrusLogger) WarnByteString(msg string, key string, value []byte) {
	if !l.levelCheck.Check(WarnLevel){
		return
	}
	l.l.WithField(key, string(value)).Warn(msg)
}

func (l *LogrusLogger) ErrorString(msg string, key, value string) {
	if !l.levelCheck.Check(ErrorLevel){
		return
	}
	l.l.WithField(key, value).Error(msg)
}

func (l *LogrusLogger) ErrorInt(msg string, key string, value int) {
	if !l.levelCheck.Check(ErrorLevel){
		return
	}
	l.l.WithField(key, value).Error(msg)
}

func (l *LogrusLogger) ErrorInt64(msg string, key string, value int64) {
	if !l.levelCheck.Check(ErrorLevel){
		return
	}
	l.l.WithField(key, value).Error(msg)
}

func (l *LogrusLogger) ErrorFloat32(msg string, key string, value float32) {
	if !l.levelCheck.Check(ErrorLevel){
		return
	}
	l.l.WithField(key, value).Error(msg)
}

func (l *LogrusLogger) ErrorFloat64(msg string, key string, value float64) {
	if !l.levelCheck.Check(ErrorLevel){
		return
	}
	l.l.WithField(key, value).Error(msg)
}

func (l *LogrusLogger) ErrorByteString(msg string, key string, value []byte) {
	if !l.levelCheck.Check(ErrorLevel){
		return
	}
	l.l.WithField(key, string(value)).Error(msg)
}

func (l *LogrusLogger) FatalString(msg string, key, value string) {
	if !l.levelCheck.Check(FatalLevel){
		return
	}
	l.l.WithField(key, value).Fatal(msg)
}

func (l *LogrusLogger) FatalInt(msg string, key string, value int) {
	if !l.levelCheck.Check(FatalLevel){
		return
	}
	l.l.WithField(key, value).Fatal(msg)
}

func (l *LogrusLogger) FatalInt64(msg string, key string, value int64) {
	if !l.levelCheck.Check(FatalLevel){
		return
	}
	l.l.WithField(key, value).Fatal(msg)
}

func (l *LogrusLogger) FatalFloat32(msg string, key string, value float32) {
	if !l.levelCheck.Check(FatalLevel){
		return
	}
	l.l.WithField(key, value).Fatal(msg)
}

func (l *LogrusLogger) FatalFloat64(msg string, key string, value float64) {
	if !l.levelCheck.Check(FatalLevel){
		return
	}
	l.l.WithField(key, value).Fatal(msg)
}

func (l *LogrusLogger) FatalByteString(msg string, key string, value []byte) {
	if !l.levelCheck.Check(FatalLevel){
		return
	}
	l.l.WithField(key, string(value)).Fatal(msg)
}

func (l *LogrusLogger) PanicString(msg string, key, value string) {
	if !l.levelCheck.Check(PanicLevel){
		return
	}
	l.l.WithField(key, value).Panic(msg)
}

func (l *LogrusLogger) PanicInt(msg string, key string, value int) {
	if !l.levelCheck.Check(PanicLevel){
		return
	}
	l.l.WithField(key, value).Panic(msg)
}

func (l *LogrusLogger) PanicInt64(msg string, key string, value int64) {
	if !l.levelCheck.Check(PanicLevel){
		return
	}
	l.l.WithField(key, value).Panic(msg)
}

func (l *LogrusLogger) PanicFloat32(msg string, key string, value float32) {
	if !l.levelCheck.Check(PanicLevel){
		return
	}
	l.l.WithField(key, value).Panic(msg)
}

func (l *LogrusLogger) PanicFloat64(msg string, key string, value float64) {
	if !l.levelCheck.Check(PanicLevel){
		return
	}
	l.l.WithField(key, value).Panic(msg)
}

func (l *LogrusLogger) PanicByteString(msg string, key string, value []byte) {
	if !l.levelCheck.Check(PanicLevel){
		return
	}
	l.l.WithField(key, string(value)).Panic(msg)
}
