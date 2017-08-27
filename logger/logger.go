package logger

// Logger interface to RIA apps
type Logger interface {
	Info(msg string)
	Error(msg string)
	Warn(msg string)
	Fatal(msg string)
	InfoWithFields(msg string, fields map[string]interface{})
	ErrorWithFields(msg string, fields map[string]interface{})
	WarnWithFields(msg string, fields map[string]interface{})
	FatalWithFields(msg string, fields map[string]interface{})
	SetEmailHook(ec EmailConfig) error
}

// EmailConfig for setting up email hook
type EmailConfig struct {
	From     string
	To       string
	Username string
	Password string
}
