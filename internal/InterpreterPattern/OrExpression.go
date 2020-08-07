package InterpreterPattern

type OrExpression struct {
	expr1 Expression
	expr2 Expression
}

func (o *OrExpression) OrExpression(expr1 Expression, expr2 Expression) {
	o.expr1 = expr1
	o.expr2 = expr2
}

func (o *OrExpression) Interpret(context string) bool {
	return o.expr1.Interpret(context) || o.expr2.Interpret(context);
}
