package ChainofResponsibilityPattern

import "fmt"

type FileLogger struct {
	abstractLogger AbstractLogger
}

func (f *FileLogger) write(message string) {
	fmt.Println("File::Logger: " + message)
}

func (f *FileLogger) LogMessage(level int, message string, logger Logger) {
	f.abstractLogger.LogMessage(level, message, f)
}
