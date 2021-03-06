package mlog

type StdLogger interface {
	Print(v ...interface{})
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

type LevelLogger interface {
	StdLogger

	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Notice(v ...interface{})
	Noticef(format string, v ...interface{})

	Warning(v ...interface{})
	Warningf(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	Critical(v ...interface{})
	Criticalf(format string, v ...interface{})

	Alert(v ...interface{})
	Alertf(format string, v ...interface{})

	Emergency(v ...interface{})
	Emergencyf(format string, v ...interface{})

	Reload() error
	Close()

	AddHandler(Handler)
}

type Option struct {
	Type         uint8
	File         string
	Levels       []string
	Handlers     []*HandlerOption
	PreventSmart bool
	Json         bool
}

type Logger struct {
	handlers []Handler
}

func NewNullLogger() (l *Logger) {
	l = &Logger{}

	return
}

func NewLogger(opt *Option) (l *Logger) {
	if nil == opt {
		opt = &Option{}
	}

	l = NewNullLogger()

	switch opt.Type {
	case TFile:
		l.AddHandler(NewLevelHandler(opt.File, opt.Levels...))
	case TBare:
		l.AddHandler(NewBareHandler(opt.File))
	case TMultiHandler:
		for _, v := range opt.Handlers {
			h := NewLevelHandler(v.File, v.Levels...)
			l.AddHandler(h)
		}
	}

	if opt.Json {
		for k, v := range l.handlers {
			l.handlers[k] = NewJsonHandler(v)
		}
	}

	if !opt.PreventSmart {
		for k, v := range l.handlers {
			l.handlers[k] = NewSmartHandler(v)
		}
	}

	return
}

func (l *Logger) AddHandler(h Handler) {
	l.handlers = append(l.handlers, h)
}

func (l *Logger) Print(v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(LevelAll, "", v))
	}
}

func (l *Logger) Println(v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(LevelAll, "", v))
	}
}

func (l *Logger) Printf(format string, v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(LevelAll, format, v))
	}
}

func (l *Logger) Debug(v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(Debug, "", v))
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(Debug, format, v))
	}
}

func (l *Logger) Info(v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(Info, "", v))
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(Info, format, v))
	}
}

func (l *Logger) Notice(v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(Notice, "", v))
	}
}

func (l *Logger) Noticef(format string, v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(Notice, format, v))
	}
}

func (l *Logger) Warning(v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(Warning, "", v))
	}
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(Warning, format, v))
	}
}

func (l *Logger) Error(v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(Error, "", v))
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(Error, format, v))
	}
}

func (l *Logger) Alert(v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(Alert, "", v))
	}
}

func (l *Logger) Alertf(format string, v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(Alert, format, v))
	}
}

func (l *Logger) Critical(v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(Critical, "", v))
	}
}

func (l *Logger) Criticalf(format string, v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(Critical, format, v))
	}
}

func (l *Logger) Emergency(v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(Emergency, "", v))
	}
}

func (l *Logger) Emergencyf(format string, v ...interface{}) {
	for _, h := range l.handlers {
		h.Log(NewRecord(Emergency, format, v))
	}
}

func (l *Logger) Reload() (err error) {
	for _, h := range l.handlers {
		if err = h.Reload(); nil != err {
			return
		}
	}

	return
}

func (l *Logger) Close() {
	for _, h := range l.handlers {
		h.Close()
	}
}
