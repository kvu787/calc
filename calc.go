package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var help string = `usage: calc [-h]

Calc evalutes arithmetic expressions in Polish notation.

Newline delimited expressions are read from standard input.`

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-h" {
		fmt.Println(help)
		os.Exit(0)
	}

	r := bufio.NewReader(os.Stdin)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			os.Exit(0)
		}
		s = strings.TrimSpace(s)
		ss := strings.Split(s, " ")
		ans, _, err := eval(ss)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(ans)
	}
	os.Exit(0)
}

func eval(expr []string) (int, []string, error) {
	if expr[0] == "+" || expr[0] == "-" || expr[0] == "*" || expr[0] == "/" {
		a, rest, err := eval(expr[1:])
		if err != nil {
			return 0, nil, err
		}
		b, rest, err := eval(rest)
		if err != nil {
			return 0, nil, err
		}

		switch expr[0] {
		case "+":
			return a + b, rest, nil
		case "-":
			return a - b, rest, nil
		case "*":
			return a * b, rest, nil
		case "/":
			return a / b, rest, nil
		default:
			return 0, nil, fmt.Errorf("unrecognized operator: %s", expr[0])
		}
	} else {
		i, err := strconv.Atoi(expr[0])
		return i, expr[1:], err
	}
}
