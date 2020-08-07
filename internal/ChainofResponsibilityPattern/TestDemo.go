package ChainofResponsibilityPattern

func TestChainofResponsibilityPattern() Logger {
	errorLogger := new(ErrorLogger)
	errorLogger.abstractLogger.level = ERROR

	fileLogger := new(FileLogger)
	fileLogger.abstractLogger.level = DEBUG

	consoleLogger := new(ConsoleLogger)
	consoleLogger.abstractLogger.level = INFO

	errorLogger.abstractLogger.SetNextLogger(fileLogger)
	fileLogger.abstractLogger.SetNextLogger(consoleLogger)

	return errorLogger
}
