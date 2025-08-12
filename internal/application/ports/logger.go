package ports

type LoggerPort interface {
	Info(msg string, fields ...any)
	Error(msg string, fields ...any)
}
