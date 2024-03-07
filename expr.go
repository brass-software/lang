package lang

type Expr struct {
	IsLiteral bool
	IsName    bool
	Value     string

	IsCall bool
	Fn     string
	Inputs []*Expr
}
