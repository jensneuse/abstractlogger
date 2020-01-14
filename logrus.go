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

func NewLogrusLogger(l *logrus.Logger, level Level) *LogrusLogger {
	return &LogrusLogger{
		l:          l,
		levelCheck: NewLevelCheck(level),
	}
}

func (l *LogrusLogger) fields(fields []Field) logrus.Fields {
	out := make(logrus.Fields, len(fields))
	for i := range fields {
		switch fields[i].Kind {
		case StringField:
			out[fields[i].Key] = fields[i].StringValue
		case ByteStringField:
			out[fields[i].Key] = string(fields[i].ByteValue)
		}
	}
	return out
}

func (l *LogrusLogger) Debug(msg string) {
	if !l.levelCheck.Check(DebugLevel) {
		return
	}
	l.l.Debug(msg)
}

func (l *LogrusLogger) DebugField(msg string, field Field) {
	if !l.levelCheck.Check(DebugLevel) {
		return
	}
	l.l.WithField(field.Key, field.InterfaceValue()).
		Debug(msg)
}

func (l *LogrusLogger) DebugField2(msg string, field1 Field, field2 Field) {
	if !l.levelCheck.Check(DebugLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		Debug(msg)
}

func (l *LogrusLogger) DebugField3(msg string, field1 Field, field2 Field, field3 Field) {
	if !l.levelCheck.Check(DebugLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		Debug(msg)
}

func (l *LogrusLogger) DebugField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field) {
	if !l.levelCheck.Check(DebugLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		WithField(field4.Key, field4.InterfaceValue()).
		Debug(msg)
}

func (l *LogrusLogger) DebugField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field) {
	if !l.levelCheck.Check(DebugLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		WithField(field4.Key, field4.InterfaceValue()).
		WithField(field5.Key, field5.InterfaceValue()).
		Debug(msg)
}

func (l *LogrusLogger) DebugField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field) {
	if !l.levelCheck.Check(DebugLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		WithField(field4.Key, field4.InterfaceValue()).
		WithField(field5.Key, field5.InterfaceValue()).
		WithField(field6.Key, field6.InterfaceValue()).
		Debug(msg)
}

func (l *LogrusLogger) Info(msg string) {
	if !l.levelCheck.Check(InfoLevel) {
		return
	}
}

func (l *LogrusLogger) InfoField(msg string, field Field) {
	if !l.levelCheck.Check(InfoLevel) {
		return
	}
	l.l.WithField(field.Key, field.InterfaceValue()).
		Info(msg)
}

func (l *LogrusLogger) InfoField2(msg string, field1 Field, field2 Field) {
	if !l.levelCheck.Check(InfoLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		Info(msg)
}

func (l *LogrusLogger) InfoField3(msg string, field1 Field, field2 Field, field3 Field) {
	if !l.levelCheck.Check(InfoLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		Info(msg)
}

func (l *LogrusLogger) InfoField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field) {
	if !l.levelCheck.Check(InfoLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		WithField(field4.Key, field4.InterfaceValue()).
		Info(msg)
}

func (l *LogrusLogger) InfoField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field) {
	if !l.levelCheck.Check(InfoLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		WithField(field4.Key, field4.InterfaceValue()).
		WithField(field5.Key, field5.InterfaceValue()).
		Info(msg)
}

func (l *LogrusLogger) InfoField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field) {
	if !l.levelCheck.Check(InfoLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		WithField(field4.Key, field4.InterfaceValue()).
		WithField(field5.Key, field5.InterfaceValue()).
		WithField(field6.Key, field6.InterfaceValue()).
		Info(msg)
}

func (l *LogrusLogger) Warn(msg string) {
	if !l.levelCheck.Check(WarnLevel) {
		return
	}
	l.l.Warn(msg)
}

func (l *LogrusLogger) WarnField(msg string, field Field) {
	if !l.levelCheck.Check(WarnLevel) {
		return
	}
	l.l.WithField(field.Key, field.InterfaceValue()).
		Warn(msg)
}

func (l *LogrusLogger) WarnField2(msg string, field1 Field, field2 Field) {
	if !l.levelCheck.Check(WarnLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		Warn(msg)
}

func (l *LogrusLogger) WarnField3(msg string, field1 Field, field2 Field, field3 Field) {
	if !l.levelCheck.Check(WarnLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		Warn(msg)
}

func (l *LogrusLogger) WarnField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field) {
	if !l.levelCheck.Check(WarnLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		WithField(field4.Key, field4.InterfaceValue()).
		Warn(msg)
}

func (l *LogrusLogger) WarnField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field) {
	if !l.levelCheck.Check(WarnLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		WithField(field4.Key, field4.InterfaceValue()).
		WithField(field5.Key, field5.InterfaceValue()).
		Warn(msg)
}

func (l *LogrusLogger) WarnField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field) {
	if !l.levelCheck.Check(WarnLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		WithField(field4.Key, field4.InterfaceValue()).
		WithField(field5.Key, field5.InterfaceValue()).
		WithField(field6.Key, field6.InterfaceValue()).
		Warn(msg)
}

func (l *LogrusLogger) Error(msg string) {
	if !l.levelCheck.Check(ErrorLevel) {
		return
	}
	l.l.Error(msg)
}

func (l *LogrusLogger) ErrorField(msg string, field Field) {
	if !l.levelCheck.Check(ErrorLevel) {
		return
	}
	l.l.WithField(field.Key, field.InterfaceValue()).
		Error(msg)
}

func (l *LogrusLogger) ErrorField2(msg string, field1 Field, field2 Field) {
	if !l.levelCheck.Check(ErrorLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		Error(msg)
}

func (l *LogrusLogger) ErrorField3(msg string, field1 Field, field2 Field, field3 Field) {
	if !l.levelCheck.Check(ErrorLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		Error(msg)
}

func (l *LogrusLogger) ErrorField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field) {
	if !l.levelCheck.Check(ErrorLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		WithField(field4.Key, field4.InterfaceValue()).
		Error(msg)
}

func (l *LogrusLogger) ErrorField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field) {
	if !l.levelCheck.Check(ErrorLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		WithField(field4.Key, field4.InterfaceValue()).
		WithField(field5.Key, field5.InterfaceValue()).
		Error(msg)
}

func (l *LogrusLogger) ErrorField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field) {
	if !l.levelCheck.Check(ErrorLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		WithField(field4.Key, field4.InterfaceValue()).
		WithField(field5.Key, field5.InterfaceValue()).
		WithField(field6.Key, field6.InterfaceValue()).
		Error(msg)
}

func (l *LogrusLogger) Fatal(msg string) {
	if !l.levelCheck.Check(FatalLevel) {
		return
	}
	l.l.Fatal(msg)
}

func (l *LogrusLogger) FatalField(msg string, field Field) {
	if !l.levelCheck.Check(FatalLevel) {
		return
	}
	l.l.WithField(field.Key, field.InterfaceValue()).
		Fatal(msg)
}

func (l *LogrusLogger) FatalField2(msg string, field1 Field, field2 Field) {
	if !l.levelCheck.Check(FatalLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		Fatal(msg)
}

func (l *LogrusLogger) FatalField3(msg string, field1 Field, field2 Field, field3 Field) {
	if !l.levelCheck.Check(FatalLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		Fatal(msg)
}

func (l *LogrusLogger) FatalField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field) {
	if !l.levelCheck.Check(FatalLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		WithField(field4.Key, field4.InterfaceValue()).
		Fatal(msg)
}

func (l *LogrusLogger) FatalField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field) {
	if !l.levelCheck.Check(FatalLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		WithField(field4.Key, field4.InterfaceValue()).
		WithField(field5.Key, field5.InterfaceValue()).
		Fatal(msg)
}

func (l *LogrusLogger) FatalField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field) {
	if !l.levelCheck.Check(FatalLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		WithField(field4.Key, field4.InterfaceValue()).
		WithField(field5.Key, field5.InterfaceValue()).
		WithField(field6.Key, field6.InterfaceValue()).
		Fatal(msg)
}

func (l *LogrusLogger) Panic(msg string) {
	if !l.levelCheck.Check(PanicLevel) {
		return
	}
	l.l.Panic(msg)
}

func (l *LogrusLogger) PanicField(msg string, field Field) {
	if !l.levelCheck.Check(PanicLevel) {
		return
	}
	l.l.WithField(field.Key, field.InterfaceValue()).
		Panic(msg)
}

func (l *LogrusLogger) PanicField2(msg string, field1 Field, field2 Field) {
	if !l.levelCheck.Check(PanicLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		Panic(msg)
}

func (l *LogrusLogger) PanicField3(msg string, field1 Field, field2 Field, field3 Field) {
	if !l.levelCheck.Check(PanicLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		Panic(msg)
}

func (l *LogrusLogger) PanicField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field) {
	if !l.levelCheck.Check(PanicLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		WithField(field4.Key, field4.InterfaceValue()).
		Panic(msg)
}

func (l *LogrusLogger) PanicField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field) {
	if !l.levelCheck.Check(PanicLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		WithField(field4.Key, field4.InterfaceValue()).
		WithField(field5.Key, field5.InterfaceValue()).
		Panic(msg)
}

func (l *LogrusLogger) PanicField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field) {
	if !l.levelCheck.Check(PanicLevel) {
		return
	}
	l.l.WithField(field1.Key, field1.InterfaceValue()).
		WithField(field2.Key, field2.InterfaceValue()).
		WithField(field3.Key, field3.InterfaceValue()).
		WithField(field4.Key, field4.InterfaceValue()).
		WithField(field5.Key, field5.InterfaceValue()).
		WithField(field6.Key, field6.InterfaceValue()).
		Panic(msg)
}
