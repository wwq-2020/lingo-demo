package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/wwq-2020/lingo-demo/functions"

	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/eval"
)

func register(fn eval.Function, err error) {
	if err != nil {
		log.Fatalf("failed to create %s function %s:", fn.Symbol(), err.Error())
	}
	err = eval.HookFunction(fn)
	if err != nil {
		log.Fatalf("failed to hook bool function %s:", err.Error())
	}
}

func main() {
	register(functions.NewFunctionJoin())
	register(functions.NewFunctionRange())

	if len(os.Args) <= 1 {
		fmt.Println("No script provided")
		os.Exit(1)
	}

	result, err := eval.RunScriptPath(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(strings.ReplaceAll(result.String(), "\\n", "\n"))

	os.Exit(0)
}
