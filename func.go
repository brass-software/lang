package lang

type Func struct {
	Inputs  []*Field
	Outputs []*Field
	Body    []*Statement
}
