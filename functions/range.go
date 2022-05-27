package functions

import (
	"fmt"
	"strconv"

	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/eval"
	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/parser"
)

type FunctionRange struct{}

func (f *FunctionRange) Desc() (string, string) {
	return fmt.Sprintf("%s%s %s%s",
			string(parser.TokLeftPar),
			f.Symbol(),
			"n 'x",
			string(parser.TokRightPar)),
		"range from n to 0"
}

func (f *FunctionRange) Symbol() parser.TokLabel {
	return parser.TokLabel("range")
}

func (f *FunctionRange) Validate(env *eval.Environment, stack *eval.StackFrame) error {
	if stack.Empty() {
		return eval.TooFewArgs(f.Symbol(), 0, 1)
	}
	return nil
}

func (f *FunctionRange) Evaluate(env *eval.Environment, stack *eval.StackFrame) (eval.Result, error) {
	intVal := stack.GetArgument(0).(*eval.IntResult)
	result := eval.NewVecResult()

	for i := intVal.Val; i > 0; i-- {
		result.AppendResult(eval.NewStringResult(strconv.Itoa(i)))
	}

	stack.Reset()
	return result, nil
}

func NewFunctionRange() (eval.Function, error) {
	fun := &FunctionRange{}
	parser.HookToken(fun.Symbol())
	return fun, nil
}
