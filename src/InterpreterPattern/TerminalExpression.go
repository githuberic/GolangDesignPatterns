package InterpreterPattern

import "strings"

type TerminalExpression struct {
	Data string
}

func (t *TerminalExpression) TerminalExpression(data string) {
	t.Data = data
}

func (t *TerminalExpression) Interpret(context string) bool {
	if strings.Contains(context, t.Data) {
		return true
	}
	return false
}
