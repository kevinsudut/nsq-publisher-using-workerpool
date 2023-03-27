package log

// Logger is interface that must be implemented by log
type Logger interface {
	Print(args ...interface{})
	Println(args ...interface{})
	Printf(format string, args ...interface{})

	Debug(args ...interface{})
	Debugln(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infoln(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnln(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorln(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalln(args ...interface{})
	Fatalf(format string, args ...interface{})
}

// Print log
func (l *Log) Print(args ...interface{}) {
	l.logger.Print(args...)
}

// Println log
func (l *Log) Println(args ...interface{}) {
	l.logger.Println(args...)
}

// Printf log
func (l *Log) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

// Debug log
func (l *Log) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

// Debugln log
func (l *Log) Debugln(args ...interface{}) {
	l.logger.Debugln(args...)
}

// Debugf log
func (l *Log) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

// Info log
func (l *Log) Info(args ...interface{}) {
	l.logger.Info(args...)
}

// Infoln log
func (l *Log) Infoln(args ...interface{}) {
	l.logger.Infoln(args...)
}

// Infof log
func (l *Log) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// Warn log
func (l *Log) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

// Warnln log
func (l *Log) Warnln(args ...interface{}) {
	l.logger.Warnln(args...)
}

// Warnf log
func (l *Log) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

// Error log
func (l *Log) Error(args ...interface{}) {
	l.logger.Error(args...)
}

// Errorln log
func (l *Log) Errorln(args ...interface{}) {
	l.logger.Errorln(args...)
}

// Errorf log
func (l *Log) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

// Fatal log
func (l *Log) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

// Fatalln log
func (l *Log) Fatalln(args ...interface{}) {
	l.logger.Fatalln(args...)
}

// Fatalf log
func (l *Log) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}
