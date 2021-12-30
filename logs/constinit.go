package logs

const (
	DEBUG = iota
	TRACE
	INFO
	WARRING
	ERROR
	FATAL
)

func getLevel(level int) string{
	switch level{
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARRING:
		return "WARRING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "ERROR"
	}
}