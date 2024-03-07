package lang

type IfStatement struct {
	Condition *Expression
	Body      []*Statement
}
