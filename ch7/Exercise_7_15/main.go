//Exercise 7.15 Write a program on Go that reads a single expression from the standard input,
//prompts the user to provide values for any variables, then evaluates the expression in the resulting environment.
//Handle all errors gracefully.

package main

import (
	"bufio"
	"fmt"
	"gopl.io/ch7/eval"
	"os"
	"strconv"
	"strings"
)

func main() {
	exitCode := 0
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Printf("Expression: ")
	stdin.Scan()
	exprStr := stdin.Text()
	fmt.Printf("Variables (<var>=<val>, eg: x=3): ")
	stdin.Scan()
	envStr := stdin.Text()
	if stdin.Err() != nil {
		fmt.Fprintln(os.Stderr, stdin.Err())
		os.Exit(1)
	}

	env := eval.Env{}
	assignments := strings.Fields(envStr)
	for _, a := range assignments {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "bad assignment: %s\n", a)
			exitCode = 2
		}
		ident, valStr := fields[0], fields[1]
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "bad value for %s, using zero: %s\n", ident, err)
			exitCode = 2
		}
		env[eval.Var(ident)] = val
	}

	expr, err := eval.Parse(exprStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bad expression: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(expr.Eval(env))
	os.Exit(exitCode)
}
