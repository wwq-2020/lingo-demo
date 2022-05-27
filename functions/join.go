package functions

import (
	"fmt"
	"strings"

	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/eval"
	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/parser"
)

type FunctionJoin struct{}

func (f *FunctionJoin) Desc() (string, string) {
	return fmt.Sprintf("%s%s %s%s",
			string(parser.TokLeftPar),
			f.Symbol(),
			"s0 ... sN",
			string(parser.TokRightPar)),
		"joins strings"
}

func (f *FunctionJoin) Symbol() parser.TokLabel {
	return parser.TokLabel("join")
}

func (f *FunctionJoin) Validate(env *eval.Environment, stack *eval.StackFrame) error {
	if stack.Empty() {
		return eval.TooFewArgs(f.Symbol(), 0, 1)
	}
	return nil
}

func (f *FunctionJoin) Evaluate(env *eval.Environment, stack *eval.StackFrame) (eval.Result, error) {
	vec := stack.GetArgument(0).(*eval.VecResult)
	var items []string
	for _, each := range vec.Data {
		items = append(items, each.(*eval.StringResult).Val)

	}
	sep := stack.GetArgument(1).(*eval.StringResult).Val

	return eval.NewStringResult(strings.Join(items, sep)), nil
}

func NewFunctionJoin() (eval.Function, error) {
	fun := &FunctionJoin{}
	parser.HookToken(fun.Symbol())
	return fun, nil
}
