package ChainofResponsibilityPattern

import "fmt"

type ConsoleLogger struct {
	abstractLogger AbstractLogger
}

func (c *ConsoleLogger) write(message string) {
	fmt.Println("Standard Console::Logger: " + message)
}

func (c *ConsoleLogger) LogMessage(level int, message string, logger Logger) {
	c.abstractLogger.LogMessage(level, message, c)
}
