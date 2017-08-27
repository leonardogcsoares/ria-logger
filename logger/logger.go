package logger

// Logger interface to RIA apps
type Logger interface {
	Info(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
	Warn(msg string, fields map[string]interface{})
	Fatal(msg string, fields map[string]interface{})
	SetEmailHook(ec EmailConfig) error
}

// EmailConfig for setting up email hook
type EmailConfig struct {
	From     string
	To       string
	Username string
	Password string
}
