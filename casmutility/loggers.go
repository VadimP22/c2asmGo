package casmutility


import "fmt"


type Logger interface {
	Println(msg string)
}


//CONSOLE LOGGER
type ConsoleLogger struct {
	prefix string
}


func (logger ConsoleLogger) Println(msg string) {
	fmt.Println(logger.prefix ,msg)
}


func NewConsoleLogger(prefix string) Logger {
	var logger ConsoleLogger
	logger.prefix = "[" + prefix + "]"
	return logger
}


//VOID LOGGER
type VoidLogger struct {
	prefix string
}


func (logger VoidLogger) Println(msg string) {

}


func NewVoidLogger() Logger {
	var logger VoidLogger
	return logger
}