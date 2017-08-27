package logrus

import (
	"fmt"
	"os"

	"github.com/leonardogcsoares/ria-logger/logger"
	"github.com/sirupsen/logrus"
	"github.com/zbindenren/logrus_mail"
)

// Logrus is the wrapper to the actual Logrus implementation
type Logrus struct {
	l *logrus.Logger
}

// New returns an instance to the wrapper to the Logrus logger
func New() *Logrus {

	var log = logrus.New()

	// Log as JSON instead of the default ASCII formatter.
	log.Formatter = &logrus.JSONFormatter{}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.Out = os.Stdout

	// Only log the warning severity or above.
	log.Level = logrus.WarnLevel

	return &Logrus{l: log}
}

// Info level logging
func (l *Logrus) Info(app, msg string, fields map[string]interface{}) {
	fields["app"] = app
	l.l.WithFields(fields).Info(msg)
}

// Warn level logging
func (l *Logrus) Warn(app, msg string, fields map[string]interface{}) {
	fields["app"] = app
	l.l.WithFields(fields).Warn(msg)
}

// Error level logging
func (l *Logrus) Error(app, msg string, fields map[string]interface{}) {
	fields["app"] = app
	l.l.WithFields(fields).Error(msg)
}

// Fatal level logging
func (l *Logrus) Fatal(app, msg string, fields map[string]interface{}) {
	fields["app"] = app
	l.l.WithFields(fields).Fatal(msg)
}

// SetEmailHook sets an email hook to Logrus for sending error messages
func (l *Logrus) SetEmailHook(app string, ec logger.EmailConfig) error {
	hook, err := logrus_mail.NewMailAuthHook(
		app,
		"smtp.gmail.com",
		587,
		ec.From,
		ec.To,
		ec.Username,
		ec.Password,
	)
	if err != nil {
		return fmt.Errorf("unable to add email hook: %v", err)
	}

	l.l.Hooks.Add(hook)
	return nil
}
