package env

import (
	"context"
	"fmt"
	appengineLog "google.golang.org/appengine/log"
	"log"
)

type Logger interface {
	Debugf(ctx context.Context, format string, args ...interface{})

	Infof(ctx context.Context, format string, args ...interface{})

	Warningf(ctx context.Context, format string, args ...interface{})

	Errorf(ctx context.Context, format string, args ...interface{})

	Criticalf(ctx context.Context, format string, args ...interface{})
}

type RemoteLogger struct{}

func (l *RemoteLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	appengineLog.Debugf(ctx, format, args)
}

func (l *RemoteLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	appengineLog.Infof(ctx, format, args)
}

func (l *RemoteLogger) Warningf(ctx context.Context, format string, args ...interface{}) {
	appengineLog.Warningf(ctx, format, args)
}

func (l *RemoteLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	appengineLog.Errorf(ctx, format, args)
}

func (l *RemoteLogger) Criticalf(ctx context.Context, format string, args ...interface{}) {
	appengineLog.Criticalf(ctx, format, args)
}

type LocalLogger struct{}

func (l *LocalLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	fmt.Printf(format+"\n", args)
}

func (l *LocalLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	fmt.Printf(format+"\n", args)
}

func (l *LocalLogger) Warningf(ctx context.Context, format string, args ...interface{}) {
	fmt.Printf(format+"\n", args)
}

func (l *LocalLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	log.Printf(format+"\n", args)
}

func (l *LocalLogger) Criticalf(ctx context.Context, format string, args ...interface{}) {
	log.Printf(format+"\n", args)
}
