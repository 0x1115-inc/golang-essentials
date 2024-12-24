// Copyright 2024 0x1115 Inc

package logger

import "log"

// SimpleLogger is the simplest struct that defines the basic methods for a logger

type SimpleLogger struct {
	Level int
}

func (l *SimpleLogger) Debugln(args ...interface{}) {
	if l.Level <= LevelDebug {
		log.Println(args...)
	}
}

func (l *SimpleLogger) Infoln(args ...interface{}) {
	if l.Level <= LevelInfo {
		log.Println(args...)
	}
}

func (l *SimpleLogger) Println(args ...interface{}) {
	log.Println(args...)
}

func (l *SimpleLogger) Warnln(args ...interface{}) {
	if l.Level <= LevelWarn {
		log.Println(args...)
	}
}

func (l *SimpleLogger) Warningln(args ...interface{}) {
	if l.Level <= LevelWarn {
		log.Println(args...)
	}
}

func (l *SimpleLogger) Errorln(args ...interface{}) {
	if l.Level <= LevelError {
		log.Println(args...)
	}
}

func (l *SimpleLogger) Fatalln(args ...interface{}) {
	if l.Level <= LevelFatal {
		log.Fatalln(args...)
	}
}

func (l *SimpleLogger) Debugf(format string, args ...interface{}) {
	if l.Level <= LevelDebug {
		log.Printf(format, args...)
	}
}

func (l *SimpleLogger) Infof(format string, args ...interface{}) {
	if l.Level <= LevelInfo {
		log.Printf(format, args...)
	}
}

func (l *SimpleLogger) Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *SimpleLogger) Warnf(format string, args ...interface{}) {
	if l.Level <= LevelWarn {
		log.Printf(format, args...)
	}
}

func (l *SimpleLogger) Warningf(format string, args ...interface{}) {
	if l.Level <= LevelWarn {
		log.Printf(format, args...)
	}
}

func (l *SimpleLogger) Errorf(format string, args ...interface{}) {
	if l.Level <= LevelError {
		log.Printf(format, args...)
	}
}

func (l *SimpleLogger) Fatalf(format string, args ...interface{}) {
	if l.Level <= LevelFatal {
		log.Fatalf(format, args...)
	}
}

func (l *SimpleLogger) Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func NewSimpleLogger(level int) *SimpleLogger {
	return &SimpleLogger{
		Level: level,
	}
}