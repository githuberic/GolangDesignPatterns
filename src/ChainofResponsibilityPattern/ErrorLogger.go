package ChainofResponsibilityPattern

import "fmt"

type ErrorLogger struct {
	abstractLogger AbstractLogger
}

func (e *ErrorLogger) write(message string) {
	fmt.Println("Error Console::Logger: " + message)
}

func (e *ErrorLogger) LogMessage(level int, message string, logger Logger) {
	e.abstractLogger.LogMessage(level, message, e)
}
