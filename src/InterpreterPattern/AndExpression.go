package InterpreterPattern

type AndExpression struct {
	expr1 Expression
	expr2 Expression
}

func (a *AndExpression) AndExpression(expr1 Expression, expr2 Expression) {
	a.expr1 = expr1
	a.expr2 = expr2
}

func (a *AndExpression) Interpret(context string) bool {
	return a.expr1.Interpret(context) && a.expr2.Interpret(context);
}
