package InterpreterPattern

func GetMaleExpression() Expression {
	robert := new(TerminalExpression)
	robert.TerminalExpression("Robert")

	john := new(TerminalExpression)
	john.TerminalExpression("John")

	orExpression := new(OrExpression)
	orExpression.OrExpression(robert, john)

	return orExpression
}

func GetMarriedWomanExpression() Expression {
	julie := new(TerminalExpression)
	julie.TerminalExpression("Julie")

	married := new(TerminalExpression)
	married.TerminalExpression("Married")

	andExpression := new(AndExpression)
	andExpression.AndExpression(julie, married)

	return andExpression
}
