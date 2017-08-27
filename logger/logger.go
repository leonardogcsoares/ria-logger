package logger

// Logger interface to RIA apps
type Logger interface {
	Info(app, msg string, fields map[string]interface{})
	Error(app, msg string, fields map[string]interface{})
	Warn(app, msg string, fields map[string]interface{})
	Fatal(app, msg string, fields map[string]interface{})
	SetEmailHook(app string, ec EmailConfig) error
}

// EmailConfig for setting up email hook
type EmailConfig struct {
	From     string
	To       string
	Username string
	Password string
}
