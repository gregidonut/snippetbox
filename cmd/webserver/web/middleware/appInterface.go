package middleware

type App interface {
	Info(msg string, args ...any)
}
