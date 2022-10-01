package main

import (
	"dariodemix/go-design-patterns/patterns"
	"dariodemix/go-design-patterns/patterns/behavioral/chain_of_responsibility"
	"dariodemix/go-design-patterns/patterns/structural/decorator"
	"fmt"
	"os"
	"strings"
)

func main() {
	argsWithoutProg := os.Args[1:]

	for i, patternArg := range argsWithoutProg {
		buildPattern(patternArg).Demo()

		if i < len(argsWithoutProg)-1 {
			fmt.Println()
		}
	}
}

func buildPattern(pattern string) patterns.Pattern {
	switch strings.ToLower(pattern) {
	case "decorator":
		return &decorator.Decorator{}
	case "chain of responsibility", "cor":
		return &chain_of_responsibility.ChainOfResponsibility{}
	}
	panic(fmt.Sprintf("Pattern %s not found", pattern))
}
