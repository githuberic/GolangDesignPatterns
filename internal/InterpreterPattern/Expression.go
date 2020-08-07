package InterpreterPattern

type Expression interface {
	Interpret(context string) bool
}
