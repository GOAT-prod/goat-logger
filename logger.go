package goatlogger

type Logger struct {
	application string `json:"app"`
	tag         string `json:"tag"`
}

type logInfo struct {
	application string   `json:"app"`
	level       LogLevel `json:"level"`
	time        string   `json:"time"`
	tag         string   `json:"tag"`
	msg         string   `json:"message"`
}

func New(app string) logger {
	return logger{
		application: app,
	}
}

func (l *logger) SetTag(tag string) {
	l.tag = tag
}

func (l *logger) Info(message string) {
	printLog(Info, l.application, l.tag, message)
}

func (l *logger) Debug(message string) {
	printLog(Debug, l.application, l.tag, message)
}

func (l *logger) Error(message string) {
	printLog(Error, l.application, l.tag, message)
}

func (l *logger) Panic(message string) {
	printLog(Panic, l.application, l.tag, message)
}
