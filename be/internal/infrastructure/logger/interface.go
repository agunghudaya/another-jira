package logger

import "github.com/sirupsen/logrus"

type Logger interface {
	Error(args ...any)
	Errorf(format string, args ...any)
	Errorln(args ...any)
	Fatalf(format string, args ...any)
	Info(args ...any)
	Infof(format string, args ...any)
	Infoln(args ...any)
	Printf(format string, args ...any)
	Println(args ...any)
	WithFields(fields map[string]any) Logger
}

type LogrusAdapter struct {
	logger *logrus.Logger
}

type LogrusEntryAdapter struct {
	entry *logrus.Entry
}

func (l *LogrusAdapter) Info(args ...any) {
	l.logger.Info(args...)
}

func (l *LogrusAdapter) Infoln(args ...any) {
	l.logger.Infoln(args...)
}

func (l *LogrusAdapter) Error(args ...any) {
	l.logger.Error(args...)
}

func (l *LogrusAdapter) Errorln(args ...any) {
	l.logger.Errorln(args...)
}

func (l *LogrusAdapter) Fatalf(format string, args ...any) {
	l.logger.Fatalf(format, args...)
}

func (l *LogrusAdapter) Println(args ...any) {
	l.logger.Println(args...)
}

func (l *LogrusAdapter) Printf(format string, args ...any) {
	l.logger.Printf(format, args...)
}

func (l *LogrusAdapter) Errorf(format string, args ...any) {
	l.logger.Errorf(format, args...)
}

func (l *LogrusAdapter) Infof(format string, args ...any) {
	l.logger.Infof(format, args...)
}

func (l *LogrusAdapter) WithFields(fields map[string]any) Logger {
	return &LogrusEntryAdapter{
		entry: l.logger.WithFields(logrus.Fields(fields)),
	}
}

func (e *LogrusEntryAdapter) Error(args ...any) {
	e.entry.Error(args...)
}

func (e *LogrusEntryAdapter) Errorf(format string, args ...any) {
	e.entry.Errorf(format, args...)
}

func (e *LogrusEntryAdapter) Errorln(args ...any) {
	e.entry.Errorln(args...)
}

func (e *LogrusEntryAdapter) Fatalf(format string, args ...any) {
	e.entry.Fatalf(format, args...)
}

func (e *LogrusEntryAdapter) Info(args ...any) {
	e.entry.Info(args...)
}

func (e *LogrusEntryAdapter) Infof(format string, args ...any) {
	e.entry.Infof(format, args...)
}

func (e *LogrusEntryAdapter) Infoln(args ...any) {
	e.entry.Infoln(args...)
}

func (e *LogrusEntryAdapter) Printf(format string, args ...any) {
	e.entry.Printf(format, args...)
}

func (e *LogrusEntryAdapter) Println(args ...any) {
	e.entry.Println(args...)
}

func (e *LogrusEntryAdapter) WithFields(fields map[string]any) Logger {
	return &LogrusEntryAdapter{
		entry: e.entry.WithFields(logrus.Fields(fields)),
	}
}
